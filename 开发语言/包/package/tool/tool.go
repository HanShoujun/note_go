package tool

import "fmt"

func init() {
	fmt.Println("tool package init")
}

type Network struct {
	Name string
	int
}

func (n *Network) FindByName(name string) {
	fmt.Println(n.Name)
}