package main

import (
   "fmt"
   "regexp"
)

func main() {
   r, _ := regexp.Compile("[0-9]")
   fmt.Println("Search term:", r)

   // check if the string contains digits
   fmt.Println("S54366456SDfhdgstf7986:", r.MatchString("S54366456SDfhdgstf7986"))
   fmt.Println("It's five o'clock now:", r.MatchString("It's five o'clock now:"))

   // return the first match
   fmt.Println("The phone number is 555-9980", r.FindString("The phone number is 555-9980"))
   fmt.Println("Alexander Hamilton:", r.FindString("Alexander Hamilton"))
}
