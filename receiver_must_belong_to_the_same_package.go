package main

import "fmt"

// method attemps to use the string variable as the receiver
func(text string) IsEmpty() bool{
   if len(text) > 0 {
      return false
   }
   return true
}

func main() {
   text := "Hi"

   fmt.Println(text)
   fmt.Println(text.IsEmpty())
}
