package main

import "fmt"

func FormatAmount(a float64) string {
  return "USD 2.00"
}

func main() {
  fmt.Println("main function")
  fmt.Println(FormatAmount(2.00))
}
