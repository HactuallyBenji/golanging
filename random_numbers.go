package main

import (
  "fmt"
  "errors"
  "math/rand"
  "time"
  "strconv"
)

func findMax(array []int) (int, error) {

  if len(array) == 0 {
    return 0, errors.New("Empty array received. No max value exists") 
  }

  currentMax := array[0]

  for i := 1; i < len(array); i++ {
    if array[i] > currentMax {
      currentMax = array[i]
    }
  }

  return currentMax, nil
}

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

  maxValue, _ := findMax(randomArray)
  fmt.Println(maxValue)
 }
