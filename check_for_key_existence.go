package main

import (
   "fmt"
)

func main() {
   m := make(map[string]int)
   m["Hi"] = 20
   m["How"] = 245
   fmt.Println("Map value:", m)

   // this retrieves i and ok.
   // i returns the value stored in this key and zero if the key doesn't exist.
   // ok is boolean and returns true if the key exists and false if it doesn't
   a, b := m["Hi"]
   fmt.Println("Value of a:", a)
   fmt.Println("Value of b:", b)

   if b == true{
      fmt.Println("The key exists")
   } else{
      fmt.Println("The key doesn't exist")
   }
}
