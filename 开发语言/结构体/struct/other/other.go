package other

import "fmt"

type CustomStruct struct {
	Name string
	Des string
}

func (o *CustomStruct) ChangeName(name string) {
	o.Name = name
	fmt.Println(o)
}
