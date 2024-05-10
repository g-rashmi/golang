package main

import "fmt"
func modify(num *int){
*num=20 ;
}
func main() {
	fmt.Println("learn pointers  in go")
	var num int =2
var ptr *int =&num
	fmt.Println(num)
	fmt.Println(ptr)
	fmt.Println(*ptr)
	n:=3 
	modify(&n)
	fmt.Println(n)
} 