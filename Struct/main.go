package main

import "fmt"
type person struct {
firstname string
lastname string 
age int
}
type contact struct{
	phone int
	address string
}
type employee struct{
	contact_details contact 
	person_details person
}
func main() {
// 	fmt.Println("learn about struct in go ") 
// var person1 person ;
// person1.firstname="rashmi"
// person1.lastname="gupta"
// person1.age=17;
// fmt.Println(person1);
rashmi:= person {
	firstname:"hii",
	age:12,
}
fmt.Println(rashmi)
// var person2= new(person) ; 
// person2.firstname="gupta" ;
// person2.lastname="rashmi";
// person2.age=33 
// fmt.Println(person2)

var ramesh employee ; 
ramesh.person_details = person{
	firstname:"ramesh",
	lastname:"gupta",
	age:23,
}
ramesh.contact_details =contact{
	phone:699999999,
address: "kot bazaar",
}
fmt.Println(ramesh,ramesh.person_details,ramesh.contact_details)
}