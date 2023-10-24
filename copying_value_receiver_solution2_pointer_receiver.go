package main

import "fmt"

type account struct {
   number string
   balance float64
}

// the method uses a pointer to account
func(a *account) withdraw(value float64) bool{
   if a.balance >= value{
      a.balance = a.balance-value
      return true
   }
   return false
}

func main() {
   a := account{}
   a.number="C21345345345355"
   a.balance=159

   fmt.Println(a.withdraw(150))
   fmt.Println(a)
}
