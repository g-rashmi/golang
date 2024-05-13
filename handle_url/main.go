package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("learn url in go ")
	//dikh rhi url bt exact string h ,represnt string ke format main h 
	myurl:="https://jsonplaceholder.typicode.com/todos/1" 
parseurl,err:=url.Parse(myurl);

if err!=nil {
	fmt.Println("error of this url",err)
	return ;
}
fmt.Printf("type of url %T %T\n",myurl,parseurl)
fmt.Println("scheme",parseurl.Scheme)
fmt.Println("host",parseurl.Host)
fmt.Println("rawquery",parseurl.RawQuery)
fmt.Println("path",parseurl.Path) 
//modify
parseurl.Path="newme/hii"
parseurl.Host="example.com/"

//createurl
fmt.Println(parseurl.String())
}