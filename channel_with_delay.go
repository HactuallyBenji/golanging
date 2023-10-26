package main

import (
   "fmt"
   "time"
)

func message(ch chan string) {
   // add a sleep delay to the channel
   time.Sleep(6 * time.Second)
   ch <- "Hello World"
}

func main() {
   ch := make(chan string)
   go message(ch)
   b := <- ch 
   fmt.Println(b)
   fmt.Println("This will execute last")
}
