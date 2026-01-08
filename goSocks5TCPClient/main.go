package main

import (
	// flag：用于把命令行参数解析成 Go 变量（本程序所有行为都由 flags 驱动）
	"flag"
	// fmt：用于格式化输出/拼接字符串（例如拼一个最简 HTTP 请求）
	"fmt"
	// io：提供 io.Copy / io.WriteString 等工具，负责数据“搬运/透传”
	"io"
	// log：用于错误打印并直接退出（log.Fatalf）
	"log"
	// net：网络基础库（Dialer / Conn 等接口都在这里）
	"net"
	// os：访问 stdin/stdout/stderr、退出码等
	"os"
	// strings：做一些字符串清理/小写转换
	"strings"
	// time：超时、deadline 相关
	"time"

	// x/net/proxy：实现了 SOCKS5 客户端 dialer。
	// 注意：这里不是自己手写 SOCKS5 协议，而是调用成熟库完成：
	// - 与 SOCKS5 代理的握手（Methods 协商、可选用户名密码认证）
	// - CONNECT 请求（把目标地址交给代理去连）
	// - 返回一个 net.Conn，之后的读写就是“对目标的透明通道”
	"golang.org/x/net/proxy"
)

func main() {
	var (
		// SOCKS5 代理地址（一般是本机/局域网/远端的代理服务监听地址）
		// 形如：127.0.0.1:1080
		socks5Addr = flag.String("proxy", "127.0.0.1:1080", "SOCKS5 代理地址，例如 127.0.0.1:1080")
		// 最终要访问的目标地址：主机:端口
		// 注意：这是“让代理去连接的目标”，不是本机直连。
		targetAddr = flag.String("target", "", "目标地址，例如 example.com:80（必填）")
		// 如果 SOCKS5 代理启用了用户名密码认证，则填 user/pass；
		// 不需要认证时留空即可。
		user = flag.String("user", "", "SOCKS5 用户名（可选）")
		pass = flag.String("pass", "", "SOCKS5 密码（可选）")
		// mode 决定连接建立后怎么使用这条 TCP 通道：
		// - pipe：把 stdin 透传到 conn，同时把 conn 输出透传到 stdout（适合调试任意 TCP 协议）
		// - http：向目标发一个简单的 HTTP GET 并把响应打印出来（用于快速验证代理可用性）
		mode = flag.String("mode", "pipe", "模式：pipe(透传 stdin/stdout) 或 http(发一个简单GET并打印响应)")
		// timeout：拨号超时（包含连接 SOCKS5 代理本身的拨号超时；目标连接是代理去做的）
		// 如果你的代理在远端或网络不稳定，适当调大。
		timeout = flag.Duration("timeout", 5*time.Second, "拨号超时")
		// rwTimeout：读写超时（基于 deadline 实现）
		//
		// 重要说明：
		// - 这里的实现是“每次 Read/Write 前都 SetRead/SetWriteDeadline(now+rwTimeout)”
		//   因此表现更像“每次 I/O 的超时”，而不是整个连接的总超时。
		// - 设置过短可能会在大响应/慢网络时频繁超时。
		rwTimeout = flag.Duration("rw-timeout", 0, "读写超时（0 表示不设置；非 0 时为每次读/写设置 deadline）")
		// http 模式下用于构造 Host 头（不一定等于 target 的 host：比如 target 是 IP 但想要 Host 为域名）
		httpHost = flag.String("http-host", "example.com", "http 模式下 Host 头")
		// http 模式下的请求路径，例如 "/" 或 "/index.html"
		httpPath = flag.String("http-path", "/", "http 模式下请求路径")
	)
	// 解析命令行参数。解析后 *targetAddr 等才是最终值。
	flag.Parse()

	// target 是必填项：如果没给就打印用法并用“参数错误”的退出码退出。
	if strings.TrimSpace(*targetAddr) == "" {
		fmt.Fprintln(os.Stderr, "错误：-target 必填，例如 -target example.com:80")
		flag.Usage()
		os.Exit(2)
	}

	// SOCKS5 认证信息：只要 user 或 pass 任意一个不为空，就尝试走用户名密码认证。
	//（是否真正需要认证由代理服务端决定；如果代理不支持/不要求也可能失败或被忽略）
	var auth *proxy.Auth
	if *user != "" || *pass != "" {
		auth = &proxy.Auth{User: *user, Password: *pass}
	}

	// “底层拨号器”：用于先连到 SOCKS5 代理地址（socks5Addr）。
	// 这里设置了拨号超时，避免卡死在 TCP 握手阶段。
	base := &net.Dialer{Timeout: *timeout}
	// 构造 SOCKS5 dialer：
	// - network: "tcp"（SOCKS5 走 TCP）
	// - addr: socks5Addr（代理地址）
	// - auth: 可选认证
	// - forward: base（底层如何拨号到代理）
	dialer, err := proxy.SOCKS5("tcp", *socks5Addr, auth, base)
	if err != nil {
		log.Fatalf("创建SOCKS5 dialer失败: %v", err)
	}

	// 通过 SOCKS5 dialer 去“连接目标地址”：
	// 注意：这里不是本机直接连 targetAddr，而是：
	// 1) 本机先连 socks5Addr
	// 2) 完成 SOCKS5 握手/认证
	// 3) 向代理发送 CONNECT targetAddr
	// 4) 代理去连 targetAddr 成功后，返回一条可读写的“隧道 conn”
	conn, err := dialer.Dial("tcp", *targetAddr)
	if err != nil {
		log.Fatalf("通过SOCKS5连接目标失败: %v", err)
	}
	defer conn.Close()

	// 统一把 mode 做清理：trim + lower，避免用户传入 " HTTP " 之类导致不匹配。
	switch strings.ToLower(strings.TrimSpace(*mode)) {
	case "http":
		// 示例：发一个最简单的HTTP请求（证明代理+目标连通）
		//
		// 说明：
		// - 这是最“朴素”的 HTTP/1.1 明文请求（不含 TLS）。
		// - 如果你 target 是 :443 或需要 HTTPS，这里不会自动做 TLS 握手；
		//   你会看到服务端返回错误/或无法解析的内容。
		// - Connection: close 让服务端尽快关闭连接，便于我们 io.Copy 读到 EOF 结束。
		req := fmt.Sprintf("GET %s HTTP/1.1\r\nHost: %s\r\nConnection: close\r\n\r\n", *httpPath, *httpHost)
		// 写请求前可选地设置写超时，避免 Write 卡住。
		if *rwTimeout > 0 {
			_ = conn.SetWriteDeadline(time.Now().Add(*rwTimeout))
		}
		if _, err := io.WriteString(conn, req); err != nil {
			log.Fatalf("写入HTTP请求失败: %v", err)
		}
		// 读响应前可选地设置读超时，避免 Read 卡住。
		if *rwTimeout > 0 {
			_ = conn.SetReadDeadline(time.Now().Add(*rwTimeout))
		}
		// 把响应原样拷贝到 stdout。打印的内容包含响应头+响应体。
		if _, err := io.Copy(os.Stdout, conn); err != nil {
			log.Fatalf("读取HTTP响应失败: %v", err)
		}
	case "pipe":
		// 透传：stdin -> conn、conn -> stdout
		//
		// 这是一个典型的“全双工 pipe”模式：
		// - goroutine A：从 os.Stdin 读，写入 conn（把你输入的字节发给对端）
		// - goroutine B：从 conn 读，写到 os.Stdout（把对端发回来的字节打印出来）
		//
		// 常见用途：
		// - 调试 TCP 服务：比如 telnet 的简化版（但不处理行编辑/回显）
		// - 结合重定向/管道：echo / type / cat 把数据喂给目标
		//
		// 注意：
		// - 这是“字节级别”的透传，不做任何协议解析。
		// - Windows/终端对 stdin 的缓冲/行模式会影响交互体验，这是正常现象。
		errCh := make(chan error, 2)

		go func() {
			// 写方向：stdin -> conn
			//
			// 如果配置了 rwTimeout，则用包装器在每次 Write 前刷新写 deadline，
			// 防止在慢网络/阻塞时无限卡住。
			var w io.Writer = conn
			if *rwTimeout > 0 {
				w = &writeDeadlineWriter{Conn: conn, Timeout: *rwTimeout}
			}
			// io.Copy 会持续读取 stdin，直到 EOF 或出错，然后把内容写到 conn。
			_, err := io.Copy(w, os.Stdin)
			// 如果底层连接支持半关闭（CloseWrite），我们在 stdin 结束时对“写方向”做半关闭：
			// - 对一些协议（例如 HTTP/1.0、某些自定义协议）可能更友好：告诉对端“我写完了”
			// - 但仍允许继续读对端返回的数据
			//
			// 提示：net.TCPConn 支持 CloseWrite；但 conn 可能是包装类型，
			// 所以这里用 interface 断言探测。
			if cw, ok := conn.(interface{ CloseWrite() error }); ok {
				_ = cw.CloseWrite()
			}
			// 把结果送到 errCh。主协程会等待任意一边先结束。
			errCh <- err
		}()

		go func() {
			// 读方向：conn -> stdout
			//
			// 同理，如果配置了 rwTimeout，则每次 Read 前刷新读 deadline。
			var r io.Reader = conn
			if *rwTimeout > 0 {
				r = &readDeadlineReader{Conn: conn, Timeout: *rwTimeout}
			}
			// io.Copy 会持续读 conn 并写 stdout，直到对端关闭/EOF/或出错。
			_, err := io.Copy(os.Stdout, r)
			errCh <- err
		}()

		// 任一方向结束就返回（通常 stdin 结束或对端关闭）
		//
		// 为什么“任一方向结束就退出”：
		// - 交互式使用时，用户 Ctrl+Z/Ctrl+D 或输入结束，stdin 会 EOF，此时就该退出。
		// - 对端主动断开时，读方向会先 EOF，也该退出。
		//
		// 代价：
		// - 另一条 goroutine 可能还在阻塞 I/O；但 main 返回后 defer conn.Close()
		//   会关闭底层连接，从而促使另一边也结束（通常足够）。
		if err := <-errCh; err != nil && err != io.EOF {
			log.Fatalf("透传失败: %v", err)
		}
	default:
		log.Fatalf("未知 mode: %q（支持 pipe/http）", *mode)
	}
}

// writeDeadlineWriter 是一个 io.Writer 适配器：
// - 对外表现为 Writer
// - 内部把写操作代理给 net.Conn
// - 每次写之前先设置 write deadline（now + Timeout）
//
// 适用场景：你想让 io.Copy 这类持续写入的逻辑具备“每次写的超时保护”。
type writeDeadlineWriter struct {
	Conn    net.Conn
	Timeout time.Duration
}

func (w *writeDeadlineWriter) Write(p []byte) (int, error) {
	// 这里忽略 SetWriteDeadline 的错误（用 _ 接住），原因：
	// - deadline 设置失败通常不可恢复（例如连接已关闭）
	// - 真正的错误会在随后的 Write 中体现出来
	_ = w.Conn.SetWriteDeadline(time.Now().Add(w.Timeout))
	return w.Conn.Write(p)
}

// readDeadlineReader 是一个 io.Reader 适配器：
// - 对外表现为 Reader
// - 内部把读操作代理给 net.Conn
// - 每次读之前先设置 read deadline（now + Timeout）
//
// 这能避免 io.Copy 在网络卡死时永久阻塞。
type readDeadlineReader struct {
	Conn    net.Conn
	Timeout time.Duration
}

func (r *readDeadlineReader) Read(p []byte) (int, error) {
	// 同 writeDeadlineWriter：忽略 SetReadDeadline 的错误，让 Read 返回实际错误。
	_ = r.Conn.SetReadDeadline(time.Now().Add(r.Timeout))
	return r.Conn.Read(p)
}
