package main

import "fmt"
func add(a,b int) int{
	return a+b;
}
func main() {
	fmt.Println("learn defer in go") 
	defer fmt.Println("hii how are u")
	fmt.Println("hii defer ") 
	fmt.Println(" hii go") 
	data:= add(3,4);
	defer fmt.Println("hii",data)
} 