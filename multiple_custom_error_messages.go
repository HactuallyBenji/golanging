package main

import (
   "fmt"
)

type withdrawError struct {
   err string
   value float64
   balance float64
}

// implement the method for the withdrawError Type
func (e *withdrawError) Error() string {
   return fmt.Sprintf("%s: withdrawal failed because the requested amount %0.2f is higher than balance: %0.2f.", e.err,e.value, e.balance)
}

func (e *withdrawError) balanceNegativeorZero() bool {
   return e.balance <= 0
}

func (e *withdrawError) AmountNegativeorZero() bool {
   return e.value <= 0
}

func (e *withdrawError) InsufficientFunds() bool {
   return e.balance - e.value < 0
}

type account struct {
   number string
   balance float64
}

func(a *account) withdraw(value float64) (bool,error) {
   if a.balance <=0 {
      return false, &withdrawError{"Withdrawal Error", value, a.balance}
   }
   if value <=0 {
      return false, &withdrawError{"Withdrawal Error", value, a.balance}
   }
   if a.balance >= value{
      a.balance = a.balance-value
      return true, nil
   }
   return false, &withdrawError{"Withdrawal Error", value, a.balance}
}

func main() {
   a := account{}
   a.number = "C21345345345355"
   a.balance = -100

   _, err := a.withdraw(46)
   if err != nil{
      if err, ok := err.(*withdrawError); ok {
         if err.AmountNegativeorZero() {
            fmt.Println("Amount to be withdrawn is negative")
         }
         if err.balanceNegativeorZero() {
            fmt.Println("Balance is negative")
         }
         if err.InsufficientFunds(){
            fmt.Println("Insufficient funds")
         }
         return
      }
   }

   fmt.Println("The withdrawal occurred succcessfully")
   fmt.Println("Your new balance is ")
   fmt.Println(a)
}
