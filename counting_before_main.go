package main

import (
  "fmt"
  "time"
)

func countUp(iterations int) {
  for i := 0; i < iterations + 1; i++ {
    time.Sleep(10 * time.Millisecond)
    fmt.Println("Counting up:", i)
  }
} 

func countDown(iterations int) {
  for i := iterations; i >= 0; i-- {
    time.Sleep(10 * time.Millisecond)
    fmt.Println("Counting down:", i)
  }
} 

func main() {
  go countUp(100)
  go countDown(100)

  time.Sleep(3 * time.Second)
  fmt.Println("Done main")
}
