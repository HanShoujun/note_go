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
a := 1
b := 1
```

> 变量声明了必须使用，不使用编译报错

> 变量名推荐使用驼峰式命名

# 常量

### 类似枚举常量

```go
const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)
```
