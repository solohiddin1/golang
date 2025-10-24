package main
import "fmt"

func main(){
fmt.Println("start")
var  myslice = []string{"monday","tuesday"}
fmt.Println(myslice)
for i:=0; i<2;i++ {
fmt.Println(myslice[i])
}
slice2 := []string{"slice2_1","2_2"}
myslice = append(myslice,slice2...)
index := 1
fmt.Println(myslice)
// myslice[0] = "updated"
myslice = append(myslice[:index],myslice[index+1:]...)
fmt.Println(myslice)

numbers := []int{}
// fmt.Println(numbers)
for i:=0; i < len(myslice); i++{
	numbers = append(numbers,len(myslice[i]))
}
fmt.Println("lenght of every item = ",numbers)

updated := []string{}

for i:=0; i<len(myslice); i++{
	myslice[i] = myslice[i]+"1"
	updated = append(updated , myslice[i]) 
}
fmt.Println("updated",updated)
}