package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	 var name string
	// fmt.Scan(&name)
	// fmt.Println("hello",  name) 
	reader :=bufio.NewReader(os.Stdin) 
	name,_=reader.ReadString('\n') 
	fmt.Println("hii" ,name)
}