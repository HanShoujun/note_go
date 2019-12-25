# 基本数据类型

### 整型


| 类型 | 描述 |
|-|-|
| uint8 | 无符号 8位整型（ 0 到 255 ）|
| uint16 | 无符号 8位整型（ 0 到 65535 ）|
| uint32 | 无符号 8位整型（ 0 到 4294967295 ）|
| uint64 | 无符号 8位整型（0 到 18446744073709551615）|
| int8 | 有符号 8位整型（-128 到 127）|
| int16 | 有符号 8位整型（-32768 到 32767）|
| int32 | 有符号 8位整型（-2147483648 到 2147483647）|
| int64 | 有符号 8位整型（-9223372036854775808 到 9223372036854775807）|

### 特殊整型

| 类型 | 描述 |
|-|-|
| uint | 32位操作系统上就是uint32，64位操作系统上就是uint64|
| int | 32位操作系统上就是int32，64位操作系统上就是int64|
| uintptr | 无符号整型，用于存放一个指针 |

> 在使用`int`和 `uint`类型时，不能假定它是32位或64位的整型，而是考虑`int`和`uint`可能在不同平台上的差异。

> 获取对象的长度的内建`len()`函数返回的长度可以根据不同平台的字节长度进行变化。实际使用中，切片或 map 的元素数量等都可以用`int`来表示。在涉及到二进制传输、读写文件的结构描述时，为了保持文件的结构不会受到不同编译目标平台字节长度的影响，不要使用`int`和 `uint`。

### 数字字面量语法

```go
v := 0b00101101 // 二进制的 101101, 相当于十进制的 45

v := 0o377 // 八进制的 377，相当于十进制的 255

v := 0x1p-2 // 十六进制的 1 除以 2²，也就是 0.25

v := 123_456 // 用 _ 来分隔数字，等于 123456
```

```go
package main
 
import "fmt"
 
func main(){
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a)  // 10
	fmt.Printf("%b \n", a)  // 1010  占位符%b表示二进制
 
	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b)  // 77
 
	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c)  // ff
	fmt.Printf("%X \n", c)  // FF
}
```

### 浮点型

两种浮点型数：`float32`和`float64`

```go
package main
import (
        "fmt"
        "math"
)
func main() {
        fmt.Printf("%f\n", math.Pi)
        fmt.Printf("%.2f\n", math.Pi)
}
```

### 复数

complex64和complex128
```go
var c1 complex64
c1 = 1 + 2i
var c2 complex128
c2 = 2 + 3i
fmt.Println(c1)
fmt.Println(c2)
```
复数有实部和虚部，complex64的实部和虚部为32位，complex128的实部和虚部为64位。

### 布尔值

布尔型数据只有 true和false 两个值

> 1. 布尔类型变量的默认值为false。
> 2. Go 语言中不允许将整型强制转换为布尔型.
> 3. 布尔型无法参与数值运算，也无法与其他类型进行转换。

### fmt

```go
package main

import "fmt"

// fmt占位符
func main() {
	var n = 100
	// 查看类型
	fmt.Printf("%T\n", n)
	fmt.Printf("%v\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	var s = "Hello 沙河！"
	fmt.Printf("字符串：%s\n", s)
	fmt.Printf("字符串：%v\n", s)
	fmt.Printf("字符串：%#v\n", s)
}
```

### 字符串

字符串以原生数据类型出现,内部实现使用UTF-8编码。字符串的值为双引号(")中的内容
```go
s1 := "hello"
s2 := "你好"
```

### 多行字符串

多行字符串时，就必须使用`反引号`字符
```go
s1 := `第一行
第二行
第三行
`
fmt.Println(s1)
```

### 字符串的常用操作

| 方法 | 介绍 |
|-|-|
| len(str) | 求长度 |
| +或fmt.Sprintf | 拼接字符串 |
| strings.Split | 分割 |
| strings.contains | 判断是否包含 |
| strings.HasPrefix,strings.HasSuffix | 前缀/后缀判断 |
| strings.Index(),strings.LastIndex() | 子串出现的位置 |
| strings.Join(a[]string, sep string) | join操作 |

### byte和rune类型

每个字符串的元素叫做“字符”，可以通过遍历或者单个获取字符串元素获得字符。 字符用单引号（’）包裹起来

```go
var a := '中'
var b := 'x'
```
字符有以下两种
1. uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
2. rune类型，代表一个 UTF-8字符。

当需要处理中文、日文或者其他复合字符时，则需要用到`rune`类型。`rune`类型实际是一个`int32`。

Go 使用了特殊的 rune 类型来处理 Unicode，让基于 Unicode 的文本处理更为方便，也可以使用 byte 型进行默认字符串处理，性能和扩展性都有照顾。

```go
// 遍历字符串
func traversalString() {
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}
```
输出：
```
104(h) 101(e) 108(l) 108(l) 111(o) 230(æ) 178(²) 153() 230(æ) 178(²) 179(³) 
104(h) 101(e) 108(l) 108(l) 111(o) 27801(沙) 27827(河)
```
字符串底层是一个byte数组，所以可以和[]byte类型相互转换。字符串是不能修改的 字符串是由byte字节组成，所以字符串的长度是byte字节的长度。 rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。

### 修改字符串

要修改字符串，需要先将其转换成`[]rune`或`[]byte`，完成后再转换为`string`。无论哪种转换，都会重新分配内存，并复制字节数组。

```go
func changeString() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}
```

### 类型转换

只有强制类型转换，没有隐式类型转换。该语法只能在两个类型之间支持相互转换的时候使用。

```
T(表达式)
```
其中，T表示要转换的类型。表达式包括变量、复杂算子和函数返回值等.

比如计算直角三角形的斜边长时使用math包的Sqrt()函数，该函数接收的是float64类型的参数，而变量a和b都是int类型的，这个时候就需要将a和b强制类型转换为float64类型。
```go
func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
```
