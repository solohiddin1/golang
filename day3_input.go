package main
import "fmt"

func main(){
fmt.Println("go")
var input int
var str string
var bool_value bool
fmt.Printf("enter int->")
fmt.Scan(&input)
fmt.Printf("enter string->")
fmt.Scan(&str)
fmt.Printf("enter boolean->")
fmt.Scan(&bool_value)
fmt.Println("string is = ",str)
fmt.Println("int = ",input)
fmt.Println("boolean =",bool_value)

}
