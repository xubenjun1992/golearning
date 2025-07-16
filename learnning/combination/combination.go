package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Employee struct {
	Person
	EmployeeID string `json:"employeeId"`
}

func PrintInfo(e Employee) {
	fmt.Printf("姓名:%v,年龄:%d", e.Name, e.Age)
}

func (e *Employee) setProp(name string, age int, employeeId string) {
	e.Age = age
	e.Name = name
	e.EmployeeID = employeeId
}

func main() {
	p := Person{
		Name: "张三",
		Age:  20,
	}
	e := Employee{
		Person:     p,
		EmployeeID: "AB123",
	}

	PrintInfo(e)

	j, erro := json.Marshal(p)
	if erro != nil {
		fmt.Println("序列化失败", erro)
		return
	}
	fmt.Println(string(j))

	e.setProp("李四", 30, "ttcc")
	f, erro := json.Marshal(e)
	if erro != nil {
		fmt.Println("序列化失败", erro)
		return
	}
	fmt.Println(string(f))
}
