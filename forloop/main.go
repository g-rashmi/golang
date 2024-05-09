package main

import "fmt"

func main() {
	fmt.Println("for loop in GO")
	for i:=1 ; i<2;i++ {
		fmt.Println(i)
}
//infinite loop 
c:=0 
for{
	fmt.Println("hii")
	if(c==1){
	break}
	c++;
}
//range simply in golang 
num :=[]int{1,2,3,3} ; 
for index,values :=range num{
	fmt.Println(index,values)
}
data :="jii ooo" ; 
for index,char :=range data{
	fmt.Printf("index is %d and char is %c\n" ,index,char)
}
}