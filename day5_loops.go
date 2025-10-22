package main 
import ("fmt"; "time")

func main(){
//fmt.Printf("this is go\n")
//var n int16
//fmt.Scanf("%d", &n)
start := time.Now()
for i:=0;  i<10; i++ {
fmt.Printf("i is -> %d\n",i)
}
i := 0
for  i<10{
fmt.Printf("still in the loop %d\n",i)
i++
}

for {
fmt.Println("this is infinite!")
}
res := time.Since(start)
//fmt.Printf("n is -> %d\n",n)
fmt.Println("time spent ->",res)

}
