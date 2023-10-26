package main

import (
   "fmt"
   "time"
)

func goroutine() {
  time.Sleep(2* time.Second)
   fmt.Println("This is my first goroutine.")
}

func anothergoroutine() {
   fmt.Println("This is my second goroutine.")
}

func main() {
   go goroutine()
   go anothergoroutine()

   time.Sleep(4* time.Second)
   fmt.Println("main Goroutine")
}
