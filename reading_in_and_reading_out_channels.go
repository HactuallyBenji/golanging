package main

import (
   "fmt"
)

func numberGenerator(ch chan int,limit int) {
   for i := 0; i < limit; i++ {
      ch <- i
   }
   close(ch)
}

func main() {
   ch := make(chan int)
   go numberGenerator(ch,20)
   
   // read the first number
   b := <- ch
   fmt.Println("Number:", b)
   
   // read the second number
   b = <-ch
   fmt.Println("Number:", b)
}
