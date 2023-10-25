package main

import (
  "fmt"
  "math/rand"
  "time"
  "strconv"
)

func main() {
  ns := rand.NewSource(time.Now().UnixNano())
  generator := rand.New(ns)

  var arrayLength string

  fmt.Println("How many random values would you like?")
  fmt.Scan(&arrayLength)

  length, _ := strconv.Atoi(arrayLength)

  var randomArray []int

  for i := 0; i < length; i++ {
    randomArray = append(randomArray, generator.Intn(200) - 100)
  }

  fmt.Println(randomArray)
 }
