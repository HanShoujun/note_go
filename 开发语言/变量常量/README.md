# 变量

变量声明方式：
```go
// 单个声明
//var a int = 1
//var b int = 1
// 批量声明
//var (
//	a = 1
//	b = 1
//)
// 短变量声明
a := 1
b := 1
```

> 变量声明了必须使用，不使用编译报错

> 变量名推荐使用驼峰式命名

### 匿名变量

匿名变量用一个下划线 `_` 表示
```go
func foo() (int, string) {
	return 10, "Q1mi"
}
func main() {
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)
}
```
匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明

> 1. 函数外的每个语句都必须以关键字开始（var、const、func等）
> 2. :=不能使用在函数外。
> 3. _多用于占位，表示忽略值。



# 常量

常量的声明使用 `const`, 常量在定义的时候必须赋值,且运行期间不能改变（赋值）
```go 
const pi = 3.1415
const e = 2.7182
// 或
const (
    pi = 3.1415
    e = 2.7182
)
// 常量n1、n2、n3的值都是100。
const (
    n1 = 100
    n2
    n3
)
```

### iota

`iota` 是go语言的常量计数器，只能在常量的表达式中使用。

`iota` 在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次

```go
const (
    n1 = iota //0
    n2 //1
    n3 //2
    n4 //3
)
```

其他实例
```go
const (
    a1 = iota //0
    a2 //1
    _
    a3 //3
)

const (
    b1 = iota //0
    b2 = 100 //100
    b3 = iota //2
    b4 // 3
)
```

```go
const (
	Readable = 1 << iota // 1
	Writable // 2
	Executable // 4
)
```
