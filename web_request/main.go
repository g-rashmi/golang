package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("learn about the web requests in go ") 
res,err:=http.Get("https://jsonplaceholder.typicode.com/todos/1") 
if err!=nil{
	fmt.Println(err)
	 return ; 
}

defer res.Body.Close() 

data,er:=ioutil.ReadAll(res.Body) 
if er!=nil{
	fmt.Println(er)
	return ;
}
fmt.Println("response",string(data))
}  