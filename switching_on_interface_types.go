package main

import "fmt"

type AccountOperations interface{
   // Methods
   computeInterest() float64
}

type SavingsAccount struct {
   number string
   balance float64
   interest float64
}

type CheckingAccount struct {
   number string
   balance float64
   interest float64
}

func(a *SavingsAccount) computeInterest() float64{
   return 0.005
}

func(a *CheckingAccount) computeInterest() float64{
   return 0.001
}

func CheckType(i interface{}) {
   switch i.(type) {
   case *SavingsAccount:
      fmt.Println("This is a saving account")
   case *CheckingAccount:
      fmt.Println("This is a checking account")
   default:
      fmt.Println("Unknown account")
   }
}

func main() {
   a := SavingsAccount{}
   a.number = "S21345345345355"
   a.balance = 159

   var ao1 AccountOperations
   ao1 = &a
   fmt.Println("Result for ao1")
   CheckType(ao1)

   b := CheckingAccount{}
   b.number = "C218678678345345355"
   b.balance = 2000

   var ao2 AccountOperations
   ao2 = &b
   fmt.Println("Result for ao2")
   CheckType(ao2)
} 
