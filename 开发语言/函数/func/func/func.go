package main

import "fmt"

func main() {

	var c calculation
	c = add
	fmt.Println(c(3, 42))
	fmt.Printf("%T \n", c)

	f := add
	fmt.Printf("%T \n", f)
	fmt.Println(f(23, 232))

	ret2 := calc(10, 20, add)
	fmt.Println(ret2) //30

	// 匿名函数
	// 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20) // 通过变量调用匿名函数
	sub := func(x, y int) int {
		return x - y
	}
	result := sub(23, 6)
	fmt.Println(result)

	{
		// 变量adder是一个函数并且它引用了其外部作用域中的x变量，
		// 此时adder就是一个闭包。 在adder的生命周期内，变量x也一直有效
		var adder = adder()
		fmt.Println(adder(9))
		fmt.Println(adder(19))

		f1, f2 := calculator(10)
		fmt.Println(f1(1), f2(2)) //11 9
		fmt.Println(f1(3), f2(4)) //12 8
		fmt.Println(f1(5), f2(6)) //13 7
	}

	// 延迟处理的语句按defer定义的逆序进行执行
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")

	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())

	x := 1
	y := 2
	// defer 函数会添加到延迟列表中，但函数中的参数如果是一个函数表达式，那么这个函数表达式则立即执行
	// 即 calcForDefer("A", x, y) 会立即执行
	defer calcForDefer("AA", x, calcForDefer("A", x, y))
	x = 10
	defer calcForDefer("BB", x, calcForDefer("B", x, y))
	y = 20

	funcA()
	funcB()
	funcC()
}

// 无参数
func foo() int {
	return 100
}
// 多个参数
func add(x int, y int) int {
	return x+y
}
// 无返回
func printSome(n int) {
	fmt.Println(n)
}
// 多值返回
func swap(x, y string) (string, string) {
	return y, x
}
// 命名返回值:应当仅用在短函数中。在长的函数中它们会影响代码的可读性。
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
// 可变参数
func intSum2(x ...int) int {
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

//使用type关键字来定义一个函数类型
type calculation func(int, int) int

// 高阶函数
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

// 闭包=函数+引用环境
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

// 闭包
func calculator(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func calcForDefer(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func funcA() {
	fmt.Println("func A")
}

func funcB() {
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}