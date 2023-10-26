package main

import "fmt"

// a function that takes a channel as input
func message(ch chan string) {
   // we use ch followed by <- to write data to the channel
   ch <- "Hello World"
}

func main() {
   // create a channel that transports string
   ch := make(chan string)

   // execute the goroutine with input as the channel
   go message(ch)

   // we read from channel into a variable b
   b := <- ch
   fmt.Println(b)
   
   fmt.Println("This will execute last")
}
