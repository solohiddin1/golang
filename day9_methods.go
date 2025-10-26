package main
import "fmt"

type Person struct{
	age int 
	name string
}


func (p *Person) greet(name string){
	fmt.Printf("hello from %s\n",name)
}

func (*Person) add(num1 int, num2 int) int{
	return num1+num2
}

func main(){
	// fmt.Println("start")
	p1 := Person{age:15,name:"name1"}
	p2 := Person{age:12,name:"name2"}
	
	p1.greet(p1.name)
	fmt.Println(p1.add(p1.age, p2.age))
	fmt.Println(p1)
	fmt.Println(p2)
}