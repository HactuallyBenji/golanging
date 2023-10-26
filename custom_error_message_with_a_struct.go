package main

import (
   "fmt"
)

// struct for error output
type withdrawError struct {
   err string
   value float64
   balance float64
}

type account struct {
   number string
   balance float64
}

func(a *account) withdraw(value float64) (bool, error) {
   if a.balance >= value{
      a.balance = a.balance - value
      return true,nil
   }
   return false, &withdrawError{"Withdraw Error", value, a.balance}
}

 // implement the method for the withdrawError Type
func (e *withdrawError) Error() string {
   return fmt.Sprintf("%s: withdrawal failed because the requested amount %0.2f is higher than balance: %0.2f", e.err, e.value, e.balance)
}

func main() {
   a := account{}
   a.number = "C21345345345355"
   a.balance = 159

   _, err := a.withdraw(200)
   if err != nil{
      fmt.Println(err)
      return
   }

   fmt.Println("The withdrawal occurred succcessfully")
   fmt.Println("Your new balance is", a)
}
