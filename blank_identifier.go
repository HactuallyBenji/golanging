package main

import (
   "fmt"
)

type account struct {
   number string
   balance float64
}
func(a *account) withdraw(value float64) (bool, error) {
   if a.balance >= value{
      a.balance = a.balance-value
      return true, nil
   }
   return false, fmt.Errorf("Withdrawal failed, because the request amount %0.2f is higher than balance: %0.2f ", value,a.balance)
}

func main() {
   a := account{}
   a.number = "C21345345345355"
   a.balance = 159

   // we use _ to omit out since we don't need it
   _,err := a.withdraw(200)
   if err!=nil{
      fmt.Println(err)
      return

   }
   fmt.Println("The withdrawal occurred succcessfully")
   fmt.Println("Your new balance is", a)
}
