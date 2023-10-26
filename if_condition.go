package main

import (
  "fmt"
)

func falseAndFalse() (bool, bool) {
  return true, true 
}

func main() {
  if falsification1, falsification2 := falseAndFalse(); falsification2 {
    fmt.Println("Falsified")
  }
  
  if falsification1 {
    fmt.Println("falsification1", falsification1)
  }
  if falsification2 {
    fmt.Println("falsification2", falsification2)
  }
}
