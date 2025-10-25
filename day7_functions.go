package main
import ("fmt"; "time")

func add(num1 int, num2 int) (int, string){
	return num1+num2, "done"

}

func loop_func(str1 string) []string{

	slice_res  := []string{}

	for i:=0; i<len(str1);i++{
		slice_res = append(slice_res, string(str1[i]))
	}
	return slice_res
	// return len(str1)

}

func main(){
	start := time.Now()
	fmt.Println(add(1,2))
	fmt.Println(loop_func("string"))
	finish := time.Since(start)
	fmt.Println("time= ",finish)
}