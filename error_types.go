package main

import (
   "fmt"
   "reflect"
   "strconv"
)

func main() {
   var s string = "10x"

   i, error := strconv.ParseInt(s,10,8) 
   fmt.Println(i)
   fmt.Println(error)
   fmt.Println(reflect.TypeOf(error))
}
