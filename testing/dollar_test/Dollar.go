package main

import "fmt"

func FormatAmount(a float64) string {
  if a < 0 {
    return "Impossible operation"
  }
  return "USD " + fmt.Sprintf("%.2f", a)
}

func SubtractFormatAmount(a, b float64) string {
  if a < 0 || b < 0 {
    return "Impossible operation"
  }
  if a >= b {
    return "USD " + fmt.Sprintf("%.2f", (a - b))
  }
  return "Impossible operation"
}

func main() {
  fmt.Println("main function")
  fmt.Println(FormatAmount(2.00))
  fmt.Println(FormatAmount(-10.00))

  fmt.Println(SubtractFormatAmount(2.00, 1.14))
  fmt.Println(SubtractFormatAmount(0.03 , 0.42))
  fmt.Println(SubtractFormatAmount(-0.03 , -0.42))
}
