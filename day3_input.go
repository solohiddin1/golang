package main

import "fmt"

func main() {

	fmt.Println("go")
	var text string
	var age int16
	fmt.Printf("enter your name and age ->")
	fmt.Scanf("%s  %d ", &text, &age)
	fmt.Printf("your name is %v and you are %v years old\n", text, age)
	//var input int
	//var str string
	//var bool_value bool
	//fmt.Printf("enter int->")
	//fmt.Scan(&input)
	//fmt.Printf("enter string->")
	//fmt.Scan(&str)
	//fmt.Printf("enter boolean->")
	//fmt.Scan(&bool_value)
	//fmt.Println("string is = ",str)
	//fmt.Println("int = ",input)
	//fmt.Println("boolean =",bool_value)

}
