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
   go numberGenerator(ch, 20)
   for b := range ch {
      fmt.Println("Number:", b)
   }
}
