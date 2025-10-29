package main
import (
	"fmt"
	"errors"
)

func error_handler(a int , b int) (float32, error){
	if b == 0{
		return 0 , errors.New("cannot divide by 0")
	}
	return float32(float32(a)/float32(b)) , nil
}

func main(){
	fmt.Println("Start")
	res , err := error_handler(2,2345)
	if err != nil{
		fmt.Println("error",err)
	}
	fmt.Println(res)
}


// package kata

// func MakeUpperCase(str string) string {
// 	res := make([]byte, len(str))
// 	for i := 0; i < len(str); i++ {
// 		c := str[i]
// 		if c >= 97 && c <= 122 {
// 			res[i] = c - 32
// 		} else {
// 			res[i] = c
// 		}
// 	}
// 	return string(res)
// }