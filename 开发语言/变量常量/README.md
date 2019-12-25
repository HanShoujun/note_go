# 变量

变量声明方式：
```go
//var a int = 1
//var b int = 1
//var (
//	a = 1
//	b = 1
//)
a := 1
b := 1
```

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
