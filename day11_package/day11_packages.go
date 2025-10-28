package main

import ("fmt"
	"day11_package/function"
)

func main(){
	fmt.Println("start")
	num1 := 1
	num2 := 2
	fmt.Println(function.Add(num1,num2))
	p1 := function.Person{Name:"name1",Age:14}
	fmt.Println(p1)

}
