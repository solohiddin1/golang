package main

import (
	"fmt"
	"math"
)

func hello(name string){
	fmt.Println("this is hello from func!", name)
}

func main()  {
	fmt.Println("Hello,世界")
	fmt.Printf("number is %g, problem. \n", math.Sqrt(5))
	fmt.Println("Go" + "lang")

	// fmt.Println(rand.Intn(10))
hello("s")
}