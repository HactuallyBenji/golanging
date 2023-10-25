package main

import "fmt"

// create first interface
type AccountOperations interface{
   // Methods
   computeInterest() float64
   displayInfo()
}

// create second interface
type UserOperations interface{
   changeANumber(number string)
}

// create a third interface which uses the first and second interface
type BankingOperations interface{
   AccountOperations
   UserOperations
}

// create a struct type
type SavingsAccount struct {
   number string
   balance float64
   interest float64
}

// implement method from interface 1
func(a *SavingsAccount) computeInterest() float64{
   return 0.001
}

// implement method from interface 2
func(a *SavingsAccount) changeANumber(number string) {
   a.number=number
}

// implement method from interface 1
func(a *SavingsAccount) displayInfo() {
   fmt.Println(a.number)
   fmt.Println(a.balance)
   fmt.Println(a.interest)
}

func main() {
   // create a SavingsAccount variable
   a := SavingsAccount{}
   a.number = "S21345345345355"
   a.balance = 159

   // create a variable of type BankingOperations
   var ao1 BankingOperations
   // implement the methods in BankingOperations
   ao1 = &a 
   ao1.displayInfo() 
   fmt.Println(ao1.computeInterest())
}
