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
fmt.Println(myslice)
myslice[0] = "updated"
myslice = delete(myslice, 1)
fmt.Println(myslice)
}
