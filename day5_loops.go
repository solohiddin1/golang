package main 
import ("fmt"; "time")

func main(){
//fmt.Printf("this is go\n")
//var n int16
//fmt.Scanf("%d", &n)
start := time.Now()
//for i:=0;  i<1000_000_0; i++ {
//fmt.Printf("i is -> %d\n",i)
//}
res := time.Since(start)
//fmt.Printf("n is -> %d\n",n)
fmt.Println("time spent ->",res)

}
