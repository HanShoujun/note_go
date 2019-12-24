package main

import (
	"fmt"
	"package/other"
)

func main() {
	var x = 299
	var y = 344
	sum := other.Add(x, y)
	fmt.Println("sum: ", sum)

	fmt.Println(other.WrapRandom(100))
}
