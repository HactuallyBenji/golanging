package main

import (
   "fmt"
   "time"
)

func goroutine() {
   fmt.Println("This is my first goroutine.")
}

func main() {
   go goroutine()

   time.Sleep(4* time.Second)
   fmt.Println("main goroutine")
}
