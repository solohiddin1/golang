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

var num *int // this is nil, if we dont assign memory of pointer
p := 10

num2 := &p

p1 := person{"monkey", 12}

fmt.Println("person = ",p1) // person class

happy_birthday(&p1)

fmt.Println("num is -> ",num) //this is nil
fmt.Println("p is -> ",p) // this is p = 10
fmt.Println("num2 is -> ",*num2) // pointer of also = 10
fmt.Println("num2 address is -> ",num2) // address of p
fmt.Println("after birthday = ",p1) 
}
