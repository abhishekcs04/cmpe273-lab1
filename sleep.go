package main

import	"fmt"
import "time"

func main(){
	
	Sleep(20)
	fmt.Println("Hello World")
}

func Sleep(x int) {
    <-time.After(time.Second * time.Duration(x))
}