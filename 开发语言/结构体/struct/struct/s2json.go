package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	c := &Class{
		Title:    "200",
		Students: make([]*Student, 0, 200),
	}

	for i := 0; i < 10; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%02d", i),
			ID:     i,
			Gender: "男",
		}
		c.Students = append(c.Students, stu)
	}

	fmt.Printf("%#v \n", c)
	data, err := json.Marshal(c)

	if err != nil {
		fmt.Println("json marshal failed")
		return
	}

	fmt.Printf("json:%s\n",data)

	str := `{"Title":"200","Students":[{"Name":"stu00","ID":0,"Gender":"男"},{"Name":"stu01","ID":1,"Gender":"男"},{"Name":"stu02","ID":2,"Gender":"男"},{"Name":"stu03","ID":3,"Gender":"男"},{"Name":"stu04","ID":4,"Gender":"男"},{"Name":"stu05","ID":5,"Gender":"男"},{"Name":"stu06","ID":6,"Gender":"男"},{"Name":"stu07","ID":7,"Gender":"男"},{"Name":"stu08","ID":8,"Gender":"男"},{"Name":"stu09","ID":9,"Gender":"男"}]}`
	c1 := &Class{}

	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed", err)
		return
	}
	fmt.Printf("%#v \n", c1)
}

type Student struct {
	Name string
	ID int
	Gender string
}

type Class struct {
	Title    string
	Students []*Student
}

func (c *Class) getStudentNumber() int {
	return len(c.Students)
}



