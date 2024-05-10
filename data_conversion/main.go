package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("learn pointers in go")
	var num int =33 ; 
	var floatnum float64 = float64(num)  ; 
	fmt.Println(floatnum)
	fmt.Printf("float is %T\n",floatnum) 

	number :=43  
	str:= strconv.Itoa(number) ; 
	fmt.Printf("string is %s and type is %T\n",str,str) 
strnum:="222"
	numb,err:=strconv.Atoi(strnum)
	fmt.Println(numb,err) 
} 