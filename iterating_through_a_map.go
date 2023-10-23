package main

import (
   "fmt"
)

func main() {
   m := make(map[string]int)
   m["Hi"] = 20
   m["How"] = 245

   for key, value := range m {
      fmt.Println("Key:", key, "Value:", value)
   }
}
