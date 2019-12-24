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

#### 导入

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
#### 导出

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
