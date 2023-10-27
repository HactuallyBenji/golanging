package main

import (
   "fmt"
   "time"
)

func main() {

   t1, e := time.Parse(
      time.RFC3339,
      "2025-05-21T12:50:41+00:00")

   fmt.Println(t1)
   fmt.Println(t1.Day())
   fmt.Println(t1.Month())
   
   // error if there is an error during parsing; 
   // nil if there is no error
   fmt.Println(e)
}
