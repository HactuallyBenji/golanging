package main

import(
   "fmt"
   "strings"
)

// create a function that main will call
// function checks that string is upper case
func isupper(x *string) bool {
   if strings.ToUpper(*x) == *x{
      return true
   }
   return false
}

func main() {
   var message string = "HELLO WORLD"
   messageptr := &message

   // return true if string is all upper case
   fmt.Println(isupper(messageptr))
}
