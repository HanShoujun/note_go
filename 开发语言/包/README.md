# 包

每个Go程序都是由包构成的。

程序入口必须是main包: package main

入口方法必须是main方法: func main()

入口的文件名无要求，不一定是main.go

```go
package main //包，表明代码所在的模块(包)

import ( //引⼊入代码依赖
	"fmt"
	"math/rand"
)

//功能实现
func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
```

### 定义包
```go
package 包名
```

- 一个文件夹下面直接包含的文件只能归属一个package，同样一个package的文件不能在多个文件夹下。
- 包名可以不和文件夹的名字一样，包名不能包含 - 符号。
- 包名为main的包为应用程序的入口包，这种包编译后会得到一个可执行文件，而编译不包含main包的源代码则不会得到可执行文件。

### 导入

关键字 `import` 导入包

```
import "fmt"
```
或者组合导入
```
import (
	"fmt"
	"math"
)
```
### 导出

如果名字以大写字母开头，那么它是导出的，外部包可导入引用。
如果名字以写字母开头，那么它是不可导出的，外部不能引用。

```go
// 小写字母不可导出
func random(n int) int {
	return rand.Intn(n)
}

// 大写字母开头，可导出
func Add(x int, y int) int {
	return x+y
}
```

### init()初始化函数

在Go语言程序执行时导入包语句会自动触发包内部`init()`函数的调用。需要注意的是： `init()`函数没有参数也没有返回值。 `init()`函数在程序运行时自动被调用执行，不能在代码中主动调用它。

包初始化执行的顺序如下图所示：

![20191227133325.png](https://raw.githubusercontent.com/HanShoujun/picBed/master/imgs/20191227133325.png)

### init()函数执行顺序

Go语言包会从`main`包开始检查其导入的所有包，每个包中又可能导入了其他的包。Go编译器由此构建出一个树状的包引用关系，再根据引用顺序决定编译顺序，依次编译这些包的代码。

在运行时，被最后导入的包会最先初始化并调用其`init()`函数， 如下图示：

![20191227134329.png](https://raw.githubusercontent.com/HanShoujun/picBed/master/imgs/20191227134329.png)


