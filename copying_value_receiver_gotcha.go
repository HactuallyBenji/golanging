package main

import "fmt"

type account struct {
   number string
   balance float64
}

// method with value receiver
func(a account) withdraw(value float64) bool{
   if a.balance >= value{
      a.balance = a.balance - value
      return true
   }
   return false
}

func main() {
   a := account{}
   a.number = "C21345345345355"
   a.balance = 159

   fmt.Println(a.withdraw(150)) // call the method defined on account.
   fmt.Println(a)
}

// a's balance does not decrease by 150 as expected
