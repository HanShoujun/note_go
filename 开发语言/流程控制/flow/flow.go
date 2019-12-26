package main

import "fmt"

func main() {
	s := "heel 是大佛"
	for _, c := range s {
		fmt.Printf("%c",c)
	}
}
