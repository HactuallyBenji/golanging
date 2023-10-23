package main

import (
   "fmt"
)

func main() {
  var m map[string]int

  fmt.Println(m) // this will display the nil map as empty: Reading
  m["Hi"] = 1 // this will throw an error: Writing.
}
