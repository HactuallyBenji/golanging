package main

import "fmt"

func main() {
   var a = make(chan int)
   
   fmt.Println(a)
   fmt.Printf("Channel Type is %T", a)
}
