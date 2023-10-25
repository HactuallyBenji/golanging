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

// create a struct type
type SavingsAccount struct {
   number string
   balance float64
   interest float64
}

// implement method from interface1
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
   var ao1 AccountOperations

   // a implements the method of interface AccountOperations
   ao1 = &a
   fmt.Println("ao1 info:")
   ao1.displayInfo()

   var uo1 UserOperations
   // a also implements the methods of interface UserOperations
   uo1 = &a
   // execute the account number change
   uo1.changeANumber("2345353453")
   fmt.Println("updated ao1 info:")
   ao1.displayInfo() // we can see the updated information
}
