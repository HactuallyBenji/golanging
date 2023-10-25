package main

import (
   "fmt"
   "strconv"
)

func main() {
   var s string = "10x"

   i,error := strconv.ParseInt(s,10,8) 

   if error == nil{
      fmt.Println(i)
   } else {
      fmt.Println("You cannot convert text into a number.")
      fmt.Println(error)
   }
}
