package main

import "fmt"

type account struct {
   number string
   balance float64
}

// function defined with value receiver
func withdraw(a account,value float64) account {
   if a.balance >= value{
      a.balance = a.balance - value
   }
   return a
}

func main() {
   a := account{}
   a.number="C21345345345355"
   a.balance=159

   ptra :=&a // create a pointer

   // this not ok because the function accepts only value arguments.
   // This statement won't execute and will throw an error.
   withdraw(ptra,150)
   fmt.Println(ptra)
}
