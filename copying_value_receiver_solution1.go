package main

import "fmt"

type account struct {
   number string
   balance float64
}

func(a account) withdraw(value float64) account {
   if a.balance >= value {
      a.balance = a.balance-value
   }
   return a
}

func main() {
   a := account{}
   a.number = "C21345345345355"
   a.balance = 159

   // assign the result of withdraw to a
   a = a.withdraw(150)

   // now the changes are properly recorded in the caller a
   fmt.Println(a)
}
