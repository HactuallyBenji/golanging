package main

import (
   "fmt"
)

func main() {
  // create an empty slice
  var s []int

  fmt.Println(s)
  fmt.Printf("Address of the slice: %p\n", &s)
  fmt.Println(len(s))
  fmt.Println(cap(s))

  s = append(s, 10)
  fmt.Println(s)
  fmt.Printf("Address of the slice: %p\n", &s)
  fmt.Println(len(s))
  fmt.Println(cap(s))

  s = append(s, 11)
  fmt.Println(s)
  fmt.Printf("Address of the slice: %p\n", &s)
  fmt.Println(len(s))
  fmt.Println(cap(s))

  words := [...] int {12, 13, 14, 15, 16, 17, 18, 19, 20}

  s = append(s, words)
  fmt.Println(s)
  fmt.Printf("Address of the slice: %p\n", &s)
  fmt.Println(len(s))
  fmt.Println(cap(s))

}
