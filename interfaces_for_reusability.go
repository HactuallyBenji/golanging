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

func main() {
   a := SavingsAccount{}
   a.number="S21345345345355"
   a.balance=159

   var ao1 AccountOperations
   ao1 = &a
   fmt.Println("savings interest:", ao1.computeInterest())

   b := CheckingAccount{}
   b.number = "C218678678345345355"
   b.balance = 2000

   var ao2 AccountOperations
   ao2 = &b
   fmt.Println("checking interest:", ao2.computeInterest())
}
