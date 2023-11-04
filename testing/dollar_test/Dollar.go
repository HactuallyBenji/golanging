package main

import "fmt"

func FormatAmount(a float64) string {
  return "USD " + fmt.Sprintf("%.2f", a)
}

func main() {
  fmt.Println("main function")
  fmt.Println(FormatAmount(2.00))
}
