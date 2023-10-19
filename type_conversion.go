package main

import (
  "fmt"
  "strconv"
  )

func main() {
  var firstNumber string
  var secondNumber string

  var firstInt int
  var secondInt int

  var firstFloat float64
  var secondFloat float64

  fmt.Print("Enter the first integer: ") // user prompt
  fmt.Scan(&firstNumber)  // store input

  fmt.Print("Enter the second integer: ")
  fmt.Scan(&secondNumber)

  firstInt, _ = strconv.Atoi(firstNumber) // convert Ascii to int
  secondInt, _ = strconv.Atoi(secondNumber)

  firstFloat, _ = strconv.ParseFloat(firstNumber, 64) // convert Ascii to float
  secondFloat, _ = strconv.ParseFloat(secondNumber, 64)

  fmt.Println(firstNumber + secondNumber) // addition of two strings

  fmt.Println(firstInt + secondInt) // addition of two ints

  fmt.Println(firstFloat + secondFloat + 0.1) // addition of two floats
}
