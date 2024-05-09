package main

import "fmt"

func main() {
	fmt.Println("learn switch in go")
	// day := 3
	// switch day {
	// case 1:
	// 	fmt.Println("monday")
	// case 2:
	// 	fmt.Println("TuesdAY")
	// default:
	// 	fmt.Println("weddd")
	// } 
	// month:="jan" 
	// switch month {
	// case "jan","feb","march":
	// 	fmt.Println("winter") ; 
	// case "april","may","june":
	// 	fmt.Println("summer")
	// default: 
	// fmt.Println("other season") 
	// }
	tem:=-99 ; 
	switch {
	case tem < 0:
		fmt.Println("freezing")
	case tem>0&&tem<45:
		fmt.Println("good")
	default:
		fmt.Println("hii")
	}
	
}