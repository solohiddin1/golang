package main
import "fmt"

type Person struct{
	age int 
	name string
}


type Animal struct{
	age int
	name string
	owner_name string
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
		
	a1 := Animal{age:1,name:"animal1",owner_name:p1.name}
	a2 := Animal{age:2,name:"animal2",owner_name:p2.name}

	p1.greet(p1.name)
	fmt.Println(p1.add(p1.age, p2.age))
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println("owner name is -> ",a1.owner_name)
	fmt.Println(a2)
}