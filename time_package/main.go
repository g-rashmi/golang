package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("learn time packages in go") 
	currentTime:= time.Now(); 
	fmt.Println(currentTime)
	format:=currentTime.Format("02-01-2006 ,  15:04:05 , Monday,3:04PM") 
	fmt.Println(format)
	datestr:="2002-07-09" 
	layoutstr:= "2006-02-01"
	format_time,_:=time.Parse(layoutstr,datestr)
	fmt.Println(format_time) 
	new_day:=currentTime.Add(24*time.Hour)
	fmt.Println(new_day)
	formatt:=new_day.Format("2006-02-01")
	fmt.Println(formatt)
} 