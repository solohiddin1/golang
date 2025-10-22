package main
import ("fmt" 
"math")

func  main(){
var num int
fmt.Printf("input number to check=")
fmt.Scanf("%d",&num)
pi := math.Pi
fmt.Printf("pi is- %0.2f \n",pi)
if float64(num) > pi {
	fmt.Println("greater than pi")
	fmt.Printf("num is %T pi is %T\n", float64(num), pi)
	} else{
	fmt.Println("smaller than pi")
	}

if num %2 == 0{
	fmt.Println("Number is even")
	}else {
	fmt.Println("Number is odd")
    }

}
