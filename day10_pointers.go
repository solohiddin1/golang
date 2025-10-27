package main
import "fmt"
type person struct {
	name string
	age int
}

// func happy_birthday(person1 *person){
// 	(person1).age = (person1).age + 1
// }
func happy_birthday(person1 *person){
	(*person1).age = (*person1).age + 1
}


func main(){

var num1 *int // this is nil, if we dont assign memory of pointer
fmt.Println("num is -> ",num1) // this is nil

p := 10
fmt.Println("p is -> ",p) // this is p = 10
num2 := &p // now num2 stores address of p , address of operator

// Dereference operator
*num2 = 15 // this will change value of data in that address

fmt.Println("p after update -> ",p) // this is p = 10
fmt.Println("num2 is -> ",*num2) // pointer of p, also = 10
fmt.Println("num2 address is -> ",num2) // address of p

p1 := person{"monkey", 12}
fmt.Println("person = ",p1) // person class
happy_birthday(&p1)
fmt.Println("after birthday = ",p1) 
}
