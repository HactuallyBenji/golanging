package main

import "fmt"

type account struct {
   number string
   balance float64
}

// method defined with value receiver
func(a account) withdraw(value float64) bool {
   if a.balance >= value{
      a.balance = a.balance - value
      return true
   }
   return false
}

func main() {
   a := account{}
   a.number="C21345345345355"
   a.balance=159

   ptra :=&a // create a pointer

   // this is ok because the method can accept both value and pointer receiver.
   ptra.withdraw(100)
   fmt.Println(ptra)
}

// despite no errors, the balance will still not be decreased
