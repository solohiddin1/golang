package main
import "fmt"

type Person struct{
	name string
	age int
} 


type Student struct{
	
	age, grade int
}


var (
	st2 = Student{age:1} // this way you can define objects of struct
)


func main(){
	// p := Person{"person", 14} // this is fully inputted
	
	// with this way we dont have to input both values to struct, 
	//its gonna make as default not inputted ones
	p := Person{name :"person"}

	st1 := Student{age :14}
	fmt.Println(st1)
	p.name = "updated_person"
	fmt.Println(p.name,p.age, "this is Student which was inputed from var()",st2)

}