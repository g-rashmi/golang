package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"net/http"

	"golang.org/x/tools/go/analysis/passes/nilfunc"
)
type Todo struct {
	Userid int `json:"userid"`
	Id int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

func performupdaterequest(){
	todo:=Todo{
		Userid: 2300,
		Title:"rashmi gupta learning golang ", 
		Completed :true,
	} 
	jsondata,err:=json.Marshal(todo)
	if err!=nil{
fmt.Println("eroor",err)
return ;
	}
	const myurl="https://jsonplaceholder.typicode.com/todos/1"
	jsonstring:=string(jsondata) ;
	jsonreader:= strings.NewReader(jsonstring)
	//create put request 
req,err:=http.NewRequest(http.MethodPut,myurl,jsonreader)
if err!=nil{
	fmt.Println(err)
	return ;
} 
req.Header.Set("Content-type","application/json")
client:=http.Client{}
res,er:=client.Do(req)

if er!=nil{
	fmt.Println(er) 
	return;
}
fmt.Println(res)
}
func performgetrequest(){
	res,err:=http.Get("https://jsonplaceholder.typicode.com/todos/1") 
	if err!=nil{
		fmt.Println(err)
		 return ; 
	}
	if(res.StatusCode!=http.StatusOK){
		fmt.Println("error getting in response")
	}
	// defer res.Body.Close() 
	
	// data,er:=ioutil.ReadAll(res.Body) 
	// if er!=nil{
	// 	fmt.Println(er)
	// 	return ;
	// }
	// fmt.Println("response",string(data))
	var todo Todo 
	error:=json.NewDecoder(res.Body).Decode(&todo)
	if(error!=nil){
		fmt.Println(error)
		return ; 
	}
	fmt.Println("todo",todo) 
}
func performpostrequest(){
	todo:=Todo{
		Userid: 23,
		Title:"rashmi gupta", 
		Completed :true,
	} 
	jsondata,err:=json.Marshal(todo)
	if err!=nil{
fmt.Println("eroor",err)
return ;
	}
	jsonstring:=string(jsondata) ;
	fmt.Println(jsonstring)
	//convert string to reader;
	jsonreader:= strings.NewReader(jsonstring)

myurl:="https://jsonplaceholder.typicode.com/todos"
	response,err:= http.Post(myurl,"application/json",jsonreader);
	if(err!=nil){
		fmt.Println(err)
	}
	// defer response.Body.Close();
	// data,_:= ioutil.ReadAll(response.Body)
	// fmt.Println(string(data)) 
	fmt.Println(response.Status)
}
func main() {
	fmt.Println("learn crud in go ") 
	
	//performgetrequest()
performpostrequest()
}