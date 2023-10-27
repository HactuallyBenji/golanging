package main

import (
   "fmt"
)

func computeMax(ch chan int, numbers [4]int) {
   max := numbers[0]
   for i := 0; i < len(numbers); i++ {
      if numbers[i] > max {
         max = numbers[i]
      }
   }
   ch <- max
   close(ch)
}

func computeMin(ch chan int, numbers [4]int) {
   min := numbers[0]
   for i := 0; i < len(numbers); i++ {
      if numbers[i] < min {
         min = numbers[i]
      }
   }
   ch <- min
   close(ch)
}

func main() {
   numbers := [4]int{25, 64, 75, 5}
   fmt.Println(numbers)

   ch1 := make(chan int)
   go computeMax(ch1, numbers)
   b, ok := <- ch1
   fmt.Printf("Channel is closed: %v\n", ok)
   fmt.Printf("Max is: %v\n", b)

   ch2 := make(chan int)
   go computeMin(ch2, numbers)
   b, ok = <- ch2
   fmt.Printf("Channel is closed: %v\n", ok)
   fmt.Printf("Min is: %v\n", b)
}
