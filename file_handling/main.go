package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("learn filehandling in go") 
	//create a file 
	// file,err:= os.Create("example.txt")
	// if(err!=nil){
	// fmt.Println(err) 
	// return;
	//  } 
	// 	defer file.Close();
	// Content:="heelooo" 
	// _, er:=io.WriteString(file,Content+"\n")  ; 
	// if(er!=nil){
	// 	fmt.Println("written in file in not done ,",err)
	// 	return 
	// }
	// fmt.Println("succesfully written done ")
// file,err := os.Open("example.txt")
// if(err!=nil){
// 	fmt.Println("error ",err)
// 	return 
// }
// defer file.Close() 
// //buffer:temporary storage ,hold data thodi der 
// buffer:= make([]byte,1024) ; 

// for{

// 	n,errr:=file.Read(buffer)
// 	if errr==io.EOF {
// 		break;
// 	} 
// 	if errr!=nil{
// 		fmt.Println(errr) ;
// 	} 


// fmt.Println(string(buffer[:n]));
// 	} 


content,err:=ioutil.ReadFile("example.txt") 
if err!=nil {
	fmt.Println(err)
	return
}
fmt.Println(string(content))
} 
