# C语言结构体与文件操作

## 1. 结构体基础

### 1.1 什么是结构体

**专业解释**：结构体（struct）是一种用户自定义的数据类型，可以将多个不同类型的数据组合在一起，形成一个逻辑整体。

**通俗解释**：结构体就像"一个盒子，里面可以装不同类型的东西"：
- 普通变量：只能装一种类型的数据（比如只能装整数）
- 结构体：可以装多种类型的数据（比如可以同时装姓名、年龄、分数等）

**生活中的类比**：
- 学生信息卡：包含姓名（字符串）、年龄（整数）、分数（浮点数）
- 这些信息组合在一起，就形成了一个"学生"的概念
- 结构体就是用来表示这种"组合信息"的工具

### 1.2 结构体的定义

```c
#include <stdio.h>
#include <string.h>

// 定义结构体类型
struct Student {
    char name[50];
    int age;
    float score;
};

int main() {
    // 声明结构体变量
    struct Student student1;
    
    // 给结构体成员赋值
    strcpy(student1.name, "张三");
    student1.age = 20;
    student1.score = 85.5;
    
    // 访问结构体成员
    printf("姓名: %s\n", student1.name);
    printf("年龄: %d\n", student1.age);
    printf("分数: %.2f\n", student1.score);
    
    return 0;
}
```

**专业解释**：
- `struct Student`：定义了一个名为`Student`的结构体类型
- 结构体包含三个成员：`name`（字符数组）、`age`（整数）、`score`（浮点数）
- 使用`.`运算符访问结构体成员

**通俗解释**：
- `struct Student { ... }`：定义一个"模板"，说明"学生"这个盒子应该装什么
- `struct Student student1`：按照模板创建一个"学生盒子"
- `student1.name`：访问这个盒子里"姓名"这个部分
- `.`就像"的"，`student1.name`就是"student1的name"

### 1.3 结构体的初始化

```c
#include <stdio.h>

struct Student {
    char name[50];
    int age;
    float score;
};

int main() {
    // 方式1：声明后逐个赋值
    struct Student s1;
    strcpy(s1.name, "李四");
    s1.age = 21;
    s1.score = 90.0;
    
    // 方式2：声明时初始化
    struct Student s2 = {"王五", 19, 88.5};
    
    // 方式3：指定成员初始化（C99标准）
    struct Student s3 = {
        .name = "赵六",
        .age = 22,
        .score = 92.0
    };
    
    return 0;
}
```

**通俗解释**：
- 方式1：先创建盒子，然后一个一个往里装东西
- 方式2：创建盒子的同时就把东西装好（按顺序）
- 方式3：创建盒子的同时装东西，但可以指定"哪个东西装到哪个位置"（更清晰）

### 1.4 typedef简化结构体

```c
#include <stdio.h>
#include <string.h>

// 使用typedef定义结构体，可以省略struct关键字
typedef struct {
    char name[50];
    int age;
    float score;
} Student;

int main() {
    Student s1;  // 不需要写struct Student
    strcpy(s1.name, "张三");
    s1.age = 20;
    s1.score = 85.5;
    
    return 0;
}
```

**专业解释**：`typedef`为类型创建别名，`typedef struct { ... } Student;`定义了一个名为`Student`的类型，使用时不需要`struct`关键字。

**通俗解释**：`typedef`就像"起外号"：
- 原来要写`struct Student s1`（全名）
- 用了`typedef`后，可以写`Student s1`（外号）
- 更简洁，但意思完全一样

## 2. 结构体数组

### 2.1 结构体数组的定义

```c
#include <stdio.h>
#include <string.h>

typedef struct {
    char name[50];
    int age;
    float score;
} Student;

int main() {
    // 定义结构体数组
    Student students[3];
    
    // 给每个学生赋值
    strcpy(students[0].name, "张三");
    students[0].age = 20;
    students[0].score = 85.5;
    
    strcpy(students[1].name, "李四");
    students[1].age = 21;
    students[1].score = 90.0;
    
    strcpy(students[2].name, "王五");
    students[2].age = 19;
    students[2].score = 88.5;
    
    // 遍历结构体数组
    int i;
    for (i = 0; i < 3; i++) {
        printf("学生%d: %s, %d岁, 分数%.2f\n", 
               i+1, students[i].name, students[i].age, students[i].score);
    }
    
    return 0;
}
```

**通俗解释**：结构体数组就像"一排学生信息卡"：
- `Student students[3]`：创建3张学生信息卡
- `students[0]`：第1张卡
- `students[0].name`：第1张卡上的姓名

### 2.2 结构体数组的初始化

```c
Student students[3] = {
    {"张三", 20, 85.5},
    {"李四", 21, 90.0},
    {"王五", 19, 88.5}
};
```

**通俗解释**：创建数组的同时，把每个学生的信息都填好。

## 3. 结构体指针

### 3.1 指向结构体的指针

```c
#include <stdio.h>
#include <string.h>

typedef struct {
    char name[50];
    int age;
    float score;
} Student;

int main() {
    Student s1 = {"张三", 20, 85.5};
    Student *ptr = &s1;  // 指向结构体的指针
    
    // 通过指针访问成员：使用->运算符
    printf("姓名: %s\n", ptr->name);
    printf("年龄: %d\n", ptr->age);
    printf("分数: %.2f\n", ptr->score);
    
    // 也可以通过解引用和.运算符访问
    printf("姓名: %s\n", (*ptr).name);  // 等价于ptr->name
    
    return 0;
}
```

**专业解释**：
- `Student *ptr`：指向`Student`类型的指针
- `ptr->name`：通过指针访问成员（箭头运算符）
- `(*ptr).name`：等价写法（先解引用，再用点运算符）

**通俗解释**：
- `ptr`：指向"学生盒子"的地址条
- `ptr->name`：根据地址条找到盒子，然后访问"姓名"部分
- `->`就像"指向的"，`ptr->name`就是"ptr指向的name"

**重要**：`ptr->name`和`(*ptr).name`完全等价，但`ptr->name`更常用、更简洁。

### 3.2 结构体指针作为函数参数

```c
#include <stdio.h>
#include <string.h>

typedef struct {
    char name[50];
    int age;
    float score;
} Student;

// 值传递：传递结构体的副本（效率低）
void printStudent1(Student s) {
    printf("姓名: %s, 年龄: %d, 分数: %.2f\n", s.name, s.age, s.score);
}

// 指针传递：传递结构体的地址（效率高，推荐）
void printStudent2(Student *s) {
    printf("姓名: %s, 年龄: %d, 分数: %.2f\n", s->name, s->age, s->score);
}

// 通过指针修改结构体
void updateScore(Student *s, float newScore) {
    s->score = newScore;
}

int main() {
    Student s1 = {"张三", 20, 85.5};
    
    printStudent1(s1);      // 值传递
    printStudent2(&s1);     // 指针传递
    
    updateScore(&s1, 95.0); // 修改分数
    printStudent2(&s1);
    
    return 0;
}
```

**专业解释**：
- 值传递结构体：会复制整个结构体，如果结构体很大，效率低
- 指针传递：只传递地址（4或8字节），效率高
- 通过指针可以修改原结构体

**通俗解释**：
- 值传递：把整个"学生盒子"复制一份给函数（如果盒子很大，复制很慢）
- 指针传递：只给函数"盒子的地址条"（很快），函数通过地址条找到盒子，可以直接修改

**建议**：结构体作为函数参数时，优先使用指针传递。

## 4. 嵌套结构体

### 4.1 结构体中包含结构体

```c
#include <stdio.h>
#include <string.h>

// 定义日期结构体
typedef struct {
    int year;
    int month;
    int day;
} Date;

// 学生结构体包含日期结构体
typedef struct {
    char name[50];
    Date birthday;  // 嵌套结构体
    float score;
} Student;

int main() {
    Student s1;
    strcpy(s1.name, "张三");
    s1.birthday.year = 2000;
    s1.birthday.month = 5;
    s1.birthday.day = 15;
    s1.score = 85.5;
    
    printf("姓名: %s\n", s1.name);
    printf("生日: %d年%d月%d日\n", 
           s1.birthday.year, s1.birthday.month, s1.birthday.day);
    printf("分数: %.2f\n", s1.score);
    
    return 0;
}
```

**通俗解释**：嵌套结构体就像"大盒子里装小盒子"：
- `Student`是大盒子，里面有个`birthday`小盒子
- `s1.birthday.year`：大盒子里的小盒子的"年份"部分

## 5. 联合体（union）

### 5.1 什么是联合体

**专业解释**：联合体（union）是一种特殊的数据类型，所有成员共享同一块内存空间。联合体的大小等于最大成员的大小。

**通俗解释**：联合体就像"一个房间，不同时间可以住不同的人，但一次只能住一个人"：
- 结构体：每个成员有自己的房间（内存）
- 联合体：所有成员共用一个房间（内存），同时只能使用一个成员

### 5.2 联合体的使用

```c
#include <stdio.h>

typedef union {
    int intValue;
    float floatValue;
    char charValue;
} Data;

int main() {
    Data d;
    
    d.intValue = 100;
    printf("整数: %d\n", d.intValue);
    
    d.floatValue = 3.14;
    printf("浮点数: %.2f\n", d.floatValue);
    printf("整数（已覆盖）: %d\n", d.intValue);  // 值已被覆盖
    
    return 0;
}
```

**专业解释**：联合体的所有成员共享同一块内存，修改一个成员会影响其他成员的值。

**通俗解释**：联合体就像"一个房间，不同时间放不同的东西"：
- 先放整数100，房间里有100
- 再放浮点数3.14，房间里的100被覆盖了，现在只有3.14
- 同时只能存一种类型的数据

**应用场景**：联合体常用于需要节省内存的场景，或者需要以不同方式解释同一块内存的场景。

## 6. 枚举（enum）

### 6.1 什么是枚举

**专业解释**：枚举（enum）是一种用户自定义类型，用于定义一组命名的整数常量。

**通俗解释**：枚举就像"给数字起名字"：
- 不用枚举：用0表示红色，1表示绿色，2表示蓝色（容易忘记）
- 用枚举：用`RED`表示红色，`GREEN`表示绿色，`BLUE`表示蓝色（更清晰）

### 6.2 枚举的使用

```c
#include <stdio.h>

// 定义枚举类型
enum Color {
    RED,      // 0
    GREEN,    // 1
    BLUE      // 2
};

// 也可以指定值
enum Status {
    PENDING = 1,
    RUNNING = 2,
    COMPLETED = 3,
    FAILED = 4
};

int main() {
    enum Color c = RED;
    
    if (c == RED) {
        printf("颜色是红色\n");
    }
    
    enum Status s = RUNNING;
    printf("状态值: %d\n", s);  // 输出：2
    
    return 0;
}
```

**专业解释**：
- 枚举值默认从0开始，依次递增
- 可以手动指定枚举值
- 枚举本质上是整数，但使用名字更清晰

**通俗解释**：枚举让代码更易读：
- `if (c == RED)`比`if (c == 0)`更清楚
- 编译器会把`RED`替换成0，但代码更容易理解

## 7. 文件操作基础

### 7.1 文件的概念

**专业解释**：文件是存储在外部存储设备（如硬盘）上的数据集合。C语言通过文件指针（FILE*）来操作文件。

**通俗解释**：文件就像"笔记本"：
- 程序运行时，数据在内存中（就像"草稿纸"，程序结束就没了）
- 文件在硬盘上（就像"笔记本"，可以永久保存）
- 文件操作：把数据从内存写到文件（保存），或从文件读到内存（读取）

### 7.2 文件打开和关闭

```c
#include <stdio.h>

int main() {
    FILE *fp;  // 文件指针
    
    // 打开文件
    fp = fopen("test.txt", "w");  // "w"表示写模式
    
    if (fp == NULL) {
        printf("文件打开失败！\n");
        return 1;
    }
    
    // 文件操作...
    
    // 关闭文件
    fclose(fp);
    
    return 0;
}
```

**专业解释**：
- `FILE*`：文件指针类型，指向文件结构体
- `fopen(filename, mode)`：打开文件，返回文件指针
- `fclose(fp)`：关闭文件，释放资源

**文件打开模式**：
- `"r"`：只读模式（文件必须存在）
- `"w"`：只写模式（如果文件存在则清空，不存在则创建）
- `"a"`：追加模式（在文件末尾添加）
- `"r+"`：读写模式（文件必须存在）
- `"w+"`：读写模式（如果文件存在则清空，不存在则创建）
- `"a+"`：读写模式（追加）

**通俗解释**：
- `fopen`：打开"笔记本"，告诉系统你要"读"还是"写"
- `fclose`：合上"笔记本"，告诉系统你用完了
- 打开文件后一定要关闭，否则可能丢失数据或占用资源

## 8. 文件的读写操作

### 8.1 字符读写

```c
#include <stdio.h>

int main() {
    FILE *fp;
    char ch;
    
    // 写入文件
    fp = fopen("test.txt", "w");
    if (fp != NULL) {
        fputc('H', fp);  // 写入一个字符
        fputc('e', fp);
        fputc('l', fp);
        fputc('l', fp);
        fputc('o', fp);
        fclose(fp);
    }
    
    // 读取文件
    fp = fopen("test.txt", "r");
    if (fp != NULL) {
        while ((ch = fgetc(fp)) != EOF) {  // EOF是文件结束标志
            printf("%c", ch);
        }
        printf("\n");
        fclose(fp);
    }
    
    return 0;
}
```

**专业解释**：
- `fputc(ch, fp)`：向文件写入一个字符
- `fgetc(fp)`：从文件读取一个字符
- `EOF`：End Of File，文件结束标志（通常是-1）

**通俗解释**：
- `fputc`：往"笔记本"里写一个字符
- `fgetc`：从"笔记本"里读一个字符
- `EOF`：读到文件末尾时返回的特殊值，表示"读完了"

### 8.2 字符串读写

```c
#include <stdio.h>

int main() {
    FILE *fp;
    char str[100];
    
    // 写入文件
    fp = fopen("test.txt", "w");
    if (fp != NULL) {
        fputs("Hello World\n", fp);  // 写入字符串
        fputs("C语言文件操作\n", fp);
        fclose(fp);
    }
    
    // 读取文件
    fp = fopen("test.txt", "r");
    if (fp != NULL) {
        while (fgets(str, sizeof(str), fp) != NULL) {  // 读取一行
            printf("%s", str);
        }
        fclose(fp);
    }
    
    return 0;
}
```

**专业解释**：
- `fputs(str, fp)`：向文件写入字符串（不自动添加换行符）
- `fgets(str, size, fp)`：从文件读取一行（最多size-1个字符，自动添加'\0'）

**通俗解释**：
- `fputs`：往"笔记本"里写一行文字
- `fgets`：从"笔记本"里读一行文字

### 8.3 格式化读写

```c
#include <stdio.h>

int main() {
    FILE *fp;
    int num = 100;
    float score = 85.5;
    char name[50] = "张三";
    
    // 格式化写入（类似printf）
    fp = fopen("data.txt", "w");
    if (fp != NULL) {
        fprintf(fp, "姓名: %s\n", name);
        fprintf(fp, "编号: %d\n", num);
        fprintf(fp, "分数: %.2f\n", score);
        fclose(fp);
    }
    
    // 格式化读取（类似scanf）
    fp = fopen("data.txt", "r");
    if (fp != NULL) {
        char readName[50];
        int readNum;
        float readScore;
        
        fscanf(fp, "姓名: %s\n", readName);
        fscanf(fp, "编号: %d\n", &readNum);
        fscanf(fp, "分数: %f\n", &readScore);
        
        printf("读取的数据: %s, %d, %.2f\n", readName, readNum, readScore);
        fclose(fp);
    }
    
    return 0;
}
```

**专业解释**：
- `fprintf(fp, format, ...)`：格式化写入文件（类似`printf`）
- `fscanf(fp, format, ...)`：格式化读取文件（类似`scanf`）

**通俗解释**：
- `fprintf`：往"笔记本"里写格式化的文字（可以包含变量值）
- `fscanf`：从"笔记本"里按格式读取数据

### 8.4 二进制文件读写

```c
#include <stdio.h>

typedef struct {
    char name[50];
    int age;
    float score;
} Student;

int main() {
    FILE *fp;
    Student s1 = {"张三", 20, 85.5};
    Student s2;
    
    // 二进制写入
    fp = fopen("student.dat", "wb");  // "wb"表示二进制写模式
    if (fp != NULL) {
        fwrite(&s1, sizeof(Student), 1, fp);  // 写入一个Student结构体
        fclose(fp);
    }
    
    // 二进制读取
    fp = fopen("student.dat", "rb");  // "rb"表示二进制读模式
    if (fp != NULL) {
        fread(&s2, sizeof(Student), 1, fp);  // 读取一个Student结构体
        printf("读取: %s, %d, %.2f\n", s2.name, s2.age, s2.score);
        fclose(fp);
    }
    
    return 0;
}
```

**专业解释**：
- `fwrite(ptr, size, count, fp)`：向文件写入二进制数据
  - `ptr`：数据指针
  - `size`：每个元素的大小
  - `count`：元素个数
- `fread(ptr, size, count, fp)`：从文件读取二进制数据

**通俗解释**：
- 文本文件：像"手写的笔记本"，人类可以读
- 二进制文件：像"机器码"，计算机直接读，人类看不懂但效率高
- `fwrite`/`fread`：直接读写内存中的数据，适合结构体等复杂数据

## 9. 文件位置操作

### 9.1 文件位置指针

```c
#include <stdio.h>

int main() {
    FILE *fp = fopen("test.txt", "r");
    if (fp != NULL) {
        // 移动到文件开头
        fseek(fp, 0, SEEK_SET);
        
        // 移动到文件末尾
        fseek(fp, 0, SEEK_END);
        
        // 从当前位置向后移动10字节
        fseek(fp, 10, SEEK_CUR);
        
        // 获取当前位置
        long pos = ftell(fp);
        printf("当前位置: %ld\n", pos);
        
        // 回到文件开头
        rewind(fp);
        
        fclose(fp);
    }
    
    return 0;
}
```

**专业解释**：
- `fseek(fp, offset, whence)`：移动文件位置指针
  - `SEEK_SET`：从文件开头
  - `SEEK_CUR`：从当前位置
  - `SEEK_END`：从文件末尾
- `ftell(fp)`：获取当前位置
- `rewind(fp)`：回到文件开头（等价于`fseek(fp, 0, SEEK_SET)`）

**通俗解释**：
- 文件位置指针就像"书签"，标记你读到哪了
- `fseek`：移动"书签"到指定位置
- `ftell`：看看"书签"在哪一页
- `rewind`：把"书签"放回第一页

## 10. 文件操作示例

### 10.1 学生信息管理系统（简化版）

```c
#include <stdio.h>
#include <string.h>

typedef struct {
    char name[50];
    int age;
    float score;
} Student;

// 添加学生
void addStudent(const char *filename) {
    FILE *fp = fopen(filename, "ab");  // 追加二进制模式
    if (fp == NULL) {
        printf("文件打开失败！\n");
        return;
    }
    
    Student s;
    printf("请输入姓名: ");
    scanf("%s", s.name);
    printf("请输入年龄: ");
    scanf("%d", &s.age);
    printf("请输入分数: ");
    scanf("%f", &s.score);
    
    fwrite(&s, sizeof(Student), 1, fp);
    fclose(fp);
    printf("添加成功！\n");
}

// 显示所有学生
void showStudents(const char *filename) {
    FILE *fp = fopen(filename, "rb");
    if (fp == NULL) {
        printf("文件不存在或无法打开！\n");
        return;
    }
    
    Student s;
    int count = 0;
    printf("\n学生列表:\n");
    printf("姓名\t\t年龄\t分数\n");
    printf("--------------------------------\n");
    
    while (fread(&s, sizeof(Student), 1, fp) == 1) {
        printf("%s\t\t%d\t%.2f\n", s.name, s.age, s.score);
        count++;
    }
    
    printf("共%d个学生\n\n", count);
    fclose(fp);
}

int main() {
    const char *filename = "students.dat";
    int choice;
    
    while (1) {
        printf("1. 添加学生\n");
        printf("2. 显示所有学生\n");
        printf("3. 退出\n");
        printf("请选择: ");
        scanf("%d", &choice);
        
        switch (choice) {
            case 1:
                addStudent(filename);
                break;
            case 2:
                showStudents(filename);
                break;
            case 3:
                return 0;
            default:
                printf("无效选择！\n");
        }
    }
    
    return 0;
}
```

**通俗解释**：这个程序实现了简单的学生信息管理：
- 可以添加学生信息到文件
- 可以从文件读取并显示所有学生信息
- 数据保存在文件中，程序关闭后数据不会丢失

## 11. 常见错误和注意事项

### 11.1 文件打开失败

```c
FILE *fp = fopen("file.txt", "r");
if (fp == NULL) {
    printf("文件打开失败！\n");
    return 1;  // 必须检查
}
```

**重要**：打开文件后必须检查是否成功，否则后续操作可能崩溃。

### 11.2 忘记关闭文件

```c
FILE *fp = fopen("file.txt", "w");
// ... 操作文件
fclose(fp);  // 必须关闭！
```

**重要**：打开的文件必须关闭，否则：
- 可能丢失数据（缓冲区未刷新）
- 占用系统资源
- 可能导致文件被锁定

### 11.3 文本模式和二进制模式

- 文本模式（`"r"`, `"w"`）：处理换行符转换（Windows下`\n`转换为`\r\n`）
- 二进制模式（`"rb"`, `"wb"`）：不进行转换，直接读写字节

**建议**：处理文本文件用文本模式，处理二进制数据（如图片、结构体）用二进制模式。

## 12. 总结

本章介绍了C语言的结构体和文件操作：
- **结构体**：组合不同类型的数据，形成逻辑整体
- **结构体数组**：存储多个结构体
- **结构体指针**：高效传递和修改结构体
- **联合体和枚举**：特殊的数据类型
- **文件操作**：实现数据的持久化存储

**关键要点**：
- 结构体用`.`访问成员，指针用`->`访问成员
- 结构体作为参数时，优先使用指针传递
- 文件操作必须检查打开是否成功
- 文件操作后必须关闭文件
- 文本文件用文本模式，二进制数据用二进制模式

**下一步学习**：内存管理、预处理器、位运算等高级特性，这些是深入理解C语言的关键。
