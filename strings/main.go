package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("learn string in go")  
	str:="apple,banana,orange,mango,princes" ; 
	parts:=strings.Split(str,",");
	fmt.Println(parts)
	stri:="one one one two"
	count:=strings.Count(stri,"one")
	fmt.Println(count) 
	str1:="hii"; 
	str2:="rashmi"  
	result:=strings.Join([]string{str1,str2},"")
	fmt.Println(result)
} 