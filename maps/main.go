package main

import "fmt"

func main() {
	fmt.Println("learn mapS in GO") ;
	//create a map 
studentgrad:=make( map[string]int) 
studentgrad["raj"]=99 ; 
studentgrad["rashmi"]=78 ; 
studentgrad["bob"]=13 ; 
studentgrad["alice"]=22 
fmt.Println( studentgrad["raj"])
//delete the lkey 
delete(studentgrad,"raj")
fmt.Println( studentgrad["raj"])
//to check whether key present or not 
grades, exist :=studentgrad["david"]
fmt.Println(grades,exist) 
// to traverse map key value ,it is unordered way
for index ,value :=range studentgrad{
	fmt.Printf("index is %s and value is %d\n",index,value)
}
//create and assign vaue to map 
person :=map[string]int {
"boby" :99 ,
"alicey" :998 ,
}
for index,value :=range person{
	fmt.Printf("key is %s and value is %d",index,value)
}

}