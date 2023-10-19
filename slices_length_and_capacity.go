package main

import (
   "fmt"
)

func main() {
  // define an array with 7 elements
  numbers := [7]int{0,1,2,5,798,43,78}
  fmt.Println(numbers)

  // define a slice s based on numbers the array
  s := numbers[1:5]
  fmt.Println(s)

  fmt.Println("Length of slice:", len(s))

  fmt.Println("Slice capacity:", cap(s))
}
