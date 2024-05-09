package main

import "fmt"
func div(a , b float64) (float64,string){
	if(b==0){
		return 0 , "denominator is not zeeo"
	}
	return a/b ,"nil"; 
}
func main() {
	fmt.Println("error haandling in go")
	ans,err:=div(10,0)
fmt.Println(ans,err)
}