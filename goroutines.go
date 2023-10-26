package main

import (
   "fmt"
)

func goroutine() {
   fmt.Println("This is my first goroutine.")
}

func main() {
   go goroutine()

   fmt.Println("main goroutine")
}
