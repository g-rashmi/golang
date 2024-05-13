package main

import (
	"encoding/json"
	"fmt"
)
type Person struct {
	Name string	`json:"name"`
	Age int       `json:"age"`
	IsAdult bool   `json:"isadult"`
}
//priority yeh json key hota then (so isadult prefer than IsAdult)
func main() {

	fmt.Println("learn json in go")
	person:=Person{
		Name:"rashmi",
		Age :44,
		IsAdult: false,
	}

	fmt.Println(person)
	//conver perosn into json(marshal)
data,err:=	json.Marshal(person)
if err!=nil{
	fmt.Println(err)
	return ;
} 
fmt.Println(string(data))
//decoding 

var persondata Person
error:=json.Unmarshal(data,&persondata)
if error!=nil{
	fmt.Println(error)
	return ;
}  
fmt.Println(persondata)
}