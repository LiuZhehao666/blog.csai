# C语言数组与字符串

## 1. 数组基础

### 1.1 什么是数组

**专业解释**：数组是一种数据结构，用于存储相同类型的多个元素。数组在内存中是连续存储的，通过索引（下标）访问元素。

**通俗解释**：数组就像一排"带编号的盒子"：
- 所有盒子类型相同（都装整数，或都装字符等）
- 盒子排成一排，编号从0开始（第1个盒子是0号，第2个是1号...）
- 你可以通过编号快速找到任何一个盒子
- 盒子在内存中是"挨着放的"，所以访问很快

### 1.2 一维数组

**声明和初始化**：

```c
#include <stdio.h>

int main() {
    // 方式1：声明时指定大小，然后逐个赋值
    int arr1[5];
    arr1[0] = 10;
    arr1[1] = 20;
    arr1[2] = 30;
    arr1[3] = 40;
    arr1[4] = 50;
    
    // 方式2：声明时初始化
    int arr2[5] = {10, 20, 30, 40, 50};
    
    // 方式3：不指定大小，由初始化列表决定
    int arr3[] = {10, 20, 30};  // 自动确定大小为3
    
    // 方式4：部分初始化，未初始化的元素自动为0
    int arr4[5] = {10, 20};  // arr4[0]=10, arr4[1]=20, 其余为0
    
    return 0;
}
```

**专业解释**：
- 数组索引从0开始，最大索引是`数组大小-1`
- 数组名代表数组首元素的地址
- 未初始化的数组元素值是未定义的（可能是随机值）

**通俗解释**：
- 数组索引就像"门牌号"，从0开始。如果数组有5个元素，索引是0、1、2、3、4。
- `int arr[5]`：创建一个能装5个整数的数组。
- `arr[0]`：访问第1个元素（索引0）。
- `arr[4]`：访问第5个元素（索引4）。
- 注意：`arr[5]`是**错误的**，因为索引最大只能是4（越界访问很危险！）。

### 1.3 访问数组元素

```c
#include <stdio.h>

int main() {
    int scores[5] = {85, 90, 78, 92, 88};
    
    // 读取元素
    printf("第一个分数: %d\n", scores[0]);
    printf("第三个分数: %d\n", scores[2]);
    
    // 修改元素
    scores[1] = 95;
    printf("修改后的第二个分数: %d\n", scores[1]);
    
    // 使用循环遍历数组
    int i;
    for (i = 0; i < 5; i++) {
        printf("scores[%d] = %d\n", i, scores[i]);
    }
    
    return 0;
}
```

**通俗解释**：访问数组就像"按门牌号找房子"。`scores[0]`是0号"房子"里的值，`scores[2]`是2号"房子"里的值。用循环可以依次访问所有元素。

### 1.4 数组越界

**重要警告**：访问数组时，索引不能超出范围！

```c
int arr[5] = {1, 2, 3, 4, 5};
printf("%d\n", arr[5]);  // 错误！索引最大是4
```

**专业解释**：数组越界访问是未定义行为，可能导致：
- 读取到其他内存区域的数据（得到错误的值）
- 修改了不应该修改的内存（导致程序崩溃）
- 安全漏洞（缓冲区溢出）

**通俗解释**：数组越界就像"走错门"：你只有5个房间（索引0-4），却要进第6个房间（索引5）。这个房间可能：
- 不存在（程序崩溃）
- 是别人的房间（读取到错误数据）
- 是危险区域（导致安全问题）

**安全做法**：始终检查索引是否在有效范围内：
```c
int index = 5;
if (index >= 0 && index < 5) {
    printf("%d\n", arr[index]);
} else {
    printf("索引越界！\n");
}
```

### 1.5 数组的大小

```c
#include <stdio.h>

int main() {
    int arr[] = {10, 20, 30, 40, 50};
    int size = sizeof(arr) / sizeof(arr[0]);
    printf("数组大小: %d\n", size);  // 输出：5
    
    return 0;
}
```

**专业解释**：
- `sizeof(arr)`返回整个数组占用的字节数
- `sizeof(arr[0])`返回一个元素占用的字节数
- 两者相除得到元素个数

**通俗解释**：`sizeof(arr)`是"整个数组占多少字节"，`sizeof(arr[0])`是"一个元素占多少字节"。用总数除以单个，就得到"有多少个元素"。

**注意**：这种方法只在数组定义的作用域内有效。如果数组作为参数传递给函数，`sizeof`会失效（因为数组会退化为指针，后面会讲）。

## 2. 多维数组

### 2.1 二维数组

**专业解释**：二维数组可以看作"数组的数组"，常用于表示表格、矩阵等二维结构。

**通俗解释**：二维数组就像"教室里的座位"：
- 有行和列（比如5行3列）
- `arr[行号][列号]`访问特定位置
- 在内存中仍然是连续存储的（按行存储）

**声明和初始化**：

```c
#include <stdio.h>

int main() {
    // 方式1：声明时初始化
    int matrix[3][4] = {
        {1, 2, 3, 4},
        {5, 6, 7, 8},
        {9, 10, 11, 12}
    };
    
    // 方式2：按顺序初始化
    int matrix2[3][4] = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12};
    
    // 方式3：部分初始化
    int matrix3[3][4] = {
        {1, 2},
        {5},
        {9, 10, 11}
    };  // 未初始化的元素为0
    
    return 0;
}
```

### 2.2 访问二维数组

```c
#include <stdio.h>

int main() {
    int matrix[3][4] = {
        {1, 2, 3, 4},
        {5, 6, 7, 8},
        {9, 10, 11, 12}
    };
    
    // 访问单个元素
    printf("第2行第3列: %d\n", matrix[1][2]);  // 输出：7
    
    // 遍历二维数组
    int i, j;
    for (i = 0; i < 3; i++) {
        for (j = 0; j < 4; j++) {
            printf("%d\t", matrix[i][j]);
        }
        printf("\n");
    }
    
    return 0;
}
```

**输出**：
```
第2行第3列: 7
1       2       3       4
5       6       7       8
9       10      11      12
```

**通俗解释**：
- `matrix[1][2]`：第2行（索引1）第3列（索引2）的元素
- 注意：行和列的索引都从0开始
- 外层循环控制行，内层循环控制列

### 2.3 三维及多维数组

```c
int cube[2][3][4];  // 三维数组：2层，每层3行4列
```

**通俗解释**：三维数组就像"多层楼，每层有座位"：
- 第一维：层数
- 第二维：行数
- 第三维：列数

实际开发中，三维以上数组较少使用。

## 3. 数组作为函数参数

### 3.1 传递数组给函数

```c
#include <stdio.h>

// 方式1：指定大小
void printArray1(int arr[5]) {
    int i;
    for (i = 0; i < 5; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");
}

// 方式2：不指定大小（推荐）
void printArray2(int arr[], int size) {
    int i;
    for (i = 0; i < size; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");
}

// 方式3：使用指针（后面会详细讲）
void printArray3(int *arr, int size) {
    int i;
    for (i = 0; i < size; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");
}

int main() {
    int numbers[5] = {1, 2, 3, 4, 5};
    
    printArray1(numbers);
    printArray2(numbers, 5);
    printArray3(numbers, 5);
    
    return 0;
}
```

**专业解释**：数组作为函数参数时，实际上传递的是数组首元素的地址（指针），而不是整个数组的副本。因此函数内对数组的修改会影响原数组。

**通俗解释**：传递数组就像"给地址"而不是"给房子"：
- 你给函数的是"房子的地址"
- 函数通过地址找到房子，可以直接修改
- 所以函数内修改数组，原数组也会改变
- 这就是为什么不需要`return`就能改变数组的原因

**重要**：函数内无法用`sizeof`获取数组大小，必须额外传递大小参数。

### 3.2 修改数组的函数

```c
#include <stdio.h>

void doubleArray(int arr[], int size) {
    int i;
    for (i = 0; i < size; i++) {
        arr[i] *= 2;  // 每个元素乘以2
    }
}

int main() {
    int numbers[5] = {1, 2, 3, 4, 5};
    
    printf("修改前: ");
    int i;
    for (i = 0; i < 5; i++) {
        printf("%d ", numbers[i]);
    }
    printf("\n");
    
    doubleArray(numbers, 5);
    
    printf("修改后: ");
    for (i = 0; i < 5; i++) {
        printf("%d ", numbers[i]);
    }
    printf("\n");
    
    return 0;
}
```

**输出**：
```
修改前: 1 2 3 4 5
修改后: 2 4 6 8 10
```

**通俗解释**：因为传递的是地址，函数直接修改了原数组，所以`main`函数中的数组也改变了。

## 4. 字符串基础

### 4.1 字符数组表示字符串

**专业解释**：C语言没有专门的字符串类型，字符串用字符数组表示，以空字符`'\0'`结尾。

**通俗解释**：字符串就像"装字符的数组"，但有个特殊标记`'\0'`表示"字符串到这里结束"。比如字符串"Hello"实际上在内存中是：`'H' 'e' 'l' 'l' 'o' '\0'`（6个字符，包括结束符）。

### 4.2 字符串的声明和初始化

```c
#include <stdio.h>

int main() {
    // 方式1：字符数组初始化（自动添加'\0'）
    char str1[] = "Hello";
    
    // 方式2：指定大小（必须留一个位置给'\0'）
    char str2[10] = "Hello";
    
    // 方式3：逐个字符初始化（必须手动添加'\0'）
    char str3[] = {'H', 'e', 'l', 'l', 'o', '\0'};
    
    // 方式4：不初始化（危险！）
    char str4[10];  // 内容是未定义的
    
    printf("str1: %s\n", str1);
    printf("str2: %s\n", str2);
    printf("str3: %s\n", str3);
    
    return 0;
}
```

**重要提示**：
- 字符串字面量`"Hello"`会自动在末尾添加`'\0'`
- 字符数组必须足够大，能容纳所有字符+`'\0'`
- `'\0'`的ASCII值是0，不是字符'0'

### 4.3 字符串的输入输出

```c
#include <stdio.h>

int main() {
    char name[50];
    
    // 方式1：使用scanf（遇到空格会停止）
    printf("请输入姓名（不能有空格）: ");
    scanf("%s", name);  // 注意：name前不需要&
    printf("您的姓名: %s\n", name);
    
    // 方式2：使用gets（不安全，不推荐）
    // gets(name);  // 可能溢出，已废弃
    
    // 方式3：使用fgets（安全，推荐）
    printf("请输入姓名（可以有空格）: ");
    fgets(name, sizeof(name), stdin);  // 最多读sizeof(name)-1个字符
    printf("您的姓名: %s\n", name);
    
    return 0;
}
```

**专业解释**：
- `scanf("%s", name)`：读取字符串，遇到空格、制表符、换行符停止
- `fgets(name, size, stdin)`：安全读取，指定最大长度，会保留换行符
- 字符串数组名本身就是地址，所以`scanf`中不需要`&`

**通俗解释**：
- `scanf("%s")`：只能读"一个词"，遇到空格就停止
- `fgets`：可以读"一行"，包括空格，更安全

### 4.4 字符串长度

```c
#include <stdio.h>
#include <string.h>

int main() {
    char str[] = "Hello";
    
    // 使用strlen函数（需要包含string.h）
    int len = strlen(str);
    printf("字符串长度: %d\n", len);  // 输出：5（不包括'\0'）
    
    // 手动计算
    int i = 0;
    while (str[i] != '\0') {
        i++;
    }
    printf("手动计算长度: %d\n", i);  // 输出：5
    
    return 0;
}
```

**专业解释**：`strlen`返回字符串中字符的个数，不包括结尾的`'\0'`。

**通俗解释**：`strlen`就像"数字符"，从开头数到`'\0'`之前，不包括`'\0'`本身。

## 5. 字符串操作函数

C语言提供了丰富的字符串处理函数（需要包含`<string.h>`）：

### 5.1 strcpy - 字符串复制

```c
#include <stdio.h>
#include <string.h>

int main() {
    char src[] = "Hello";
    char dest[20];
    
    strcpy(dest, src);  // 把src复制到dest
    printf("dest: %s\n", dest);  // 输出：Hello
    
    return 0;
}
```

**专业解释**：`strcpy(dest, src)`将`src`指向的字符串（包括`'\0'`）复制到`dest`。

**通俗解释**：`strcpy`就像"抄写"：把源字符串的内容完整地复制到目标字符串。

**安全版本**：`strncpy`可以指定最大复制长度：
```c
char dest[10];
strncpy(dest, src, sizeof(dest) - 1);
dest[sizeof(dest) - 1] = '\0';  // 确保以'\0'结尾
```

### 5.2 strcat - 字符串连接

```c
#include <stdio.h>
#include <string.h>

int main() {
    char str1[20] = "Hello";
    char str2[] = " World";
    
    strcat(str1, str2);  // 把str2连接到str1后面
    printf("连接后: %s\n", str1);  // 输出：Hello World
    
    return 0;
}
```

**专业解释**：`strcat(dest, src)`将`src`追加到`dest`的末尾，覆盖`dest`原来的`'\0'`，并在新字符串末尾添加`'\0'`。

**通俗解释**：`strcat`就像"拼接"：把第二个字符串接到第一个字符串的后面。

**注意**：`dest`必须有足够的空间容纳连接后的字符串。

### 5.3 strcmp - 字符串比较

```c
#include <stdio.h>
#include <string.h>

int main() {
    char str1[] = "apple";
    char str2[] = "banana";
    char str3[] = "apple";
    
    int result1 = strcmp(str1, str2);  // 比较str1和str2
    int result2 = strcmp(str1, str3);  // 比较str1和str3
    
    printf("str1 vs str2: %d\n", result1);  // 负数（str1 < str2）
    printf("str1 vs str3: %d\n", result2);  // 0（相等）
    
    if (strcmp(str1, str2) < 0) {
        printf("str1 在字典序中排在 str2 前面\n");
    }
    
    return 0;
}
```

**专业解释**：
- `strcmp(str1, str2)`返回：
  - 负数：如果`str1 < str2`（按字典序）
  - 0：如果`str1 == str2`
  - 正数：如果`str1 > str2`

**通俗解释**：`strcmp`就像"比大小"：
- 返回负数：第一个字符串"小"（字典序靠前）
- 返回0：两个字符串相同
- 返回正数：第一个字符串"大"（字典序靠后）

### 5.4 strchr - 查找字符

```c
#include <stdio.h>
#include <string.h>

int main() {
    char str[] = "Hello World";
    char *p = strchr(str, 'o');  // 查找字符'o'
    
    if (p != NULL) {
        printf("找到字符'o'，位置: %ld\n", p - str);  // 输出：4
        printf("从该位置开始的字符串: %s\n", p);  // 输出：o World
    } else {
        printf("未找到\n");
    }
    
    return 0;
}
```

**专业解释**：`strchr(str, ch)`在字符串`str`中查找字符`ch`，返回指向第一次出现位置的指针，如果未找到返回`NULL`。

**通俗解释**：`strchr`就像"找字符"：在字符串里找指定的字符，找到了就告诉你位置。

### 5.5 其他常用字符串函数

```c
#include <stdio.h>
#include <string.h>

int main() {
    char str[] = "Hello World";
    
    // strstr - 查找子字符串
    char *p = strstr(str, "World");
    if (p != NULL) {
        printf("找到子串，位置: %ld\n", p - str);
    }
    
    // strtok - 分割字符串
    char str2[] = "apple,banana,orange";
    char *token = strtok(str2, ",");
    while (token != NULL) {
        printf("分割结果: %s\n", token);
        token = strtok(NULL, ",");
    }
    
    return 0;
}
```

## 6. 字符数组 vs 字符串字面量

### 6.1 字符数组（可修改）

```c
char str[] = "Hello";
str[0] = 'h';  // 可以修改
printf("%s\n", str);  // 输出：hello
```

### 6.2 字符串字面量（不可修改）

```c
char *str = "Hello";
// str[0] = 'h';  // 错误！可能导致程序崩溃
```

**专业解释**：字符串字面量存储在只读内存区域，试图修改会导致未定义行为。

**通俗解释**：
- `char str[] = "Hello"`：创建一个数组，把"Hello"复制进去，可以修改
- `char *str = "Hello"`：`str`指向一个只读的字符串，不能修改

**建议**：如果需要修改字符串，使用字符数组；如果只是读取，可以使用指针。

## 7. 常见错误和注意事项

### 7.1 数组越界

```c
char str[5] = "Hello";  // 错误！"Hello"需要6个字符（包括'\0'）
```

**正确做法**：
```c
char str[6] = "Hello";  // 或者
char str[] = "Hello";   // 让编译器自动确定大小
```

### 7.2 忘记'\0'

```c
char str[5] = {'H', 'e', 'l', 'l', 'o'};  // 缺少'\0'
printf("%s\n", str);  // 可能输出乱码或崩溃
```

**正确做法**：
```c
char str[6] = {'H', 'e', 'l', 'l', 'o', '\0'};
```

### 7.3 缓冲区溢出

```c
char name[10];
scanf("%s", name);  // 如果输入超过9个字符，会溢出！
```

**安全做法**：
```c
char name[10];
scanf("%9s", name);  // 最多读9个字符，留一个给'\0'
// 或者使用fgets
fgets(name, sizeof(name), stdin);
```

## 8. 总结

本章介绍了C语言的数组和字符串：
- **数组**：存储多个相同类型的数据，通过索引访问
- **多维数组**：表示表格、矩阵等结构
- **字符串**：用字符数组表示，以`'\0'`结尾
- **字符串函数**：复制、连接、比较、查找等操作

**关键要点**：
- 数组索引从0开始，注意不要越界
- 字符串必须以`'\0'`结尾
- 数组作为函数参数时传递的是地址
- 使用字符串函数要注意缓冲区大小

**下一步学习**：指针是C语言的核心概念，理解指针对于掌握C语言至关重要。下一章将深入讲解指针。
