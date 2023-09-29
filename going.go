/*
  Ben Miller
  2023-09-28
*/

package main

import (
  "fmt"
  "strconv"
);

func main() {
  fmt.Println("Hello, world!")
  fmt.Println("Enter your first name: ");
  
  var firstName string;
  fmt.Scanln(&firstName);

  fmt.Println("Enter your last name: ");

  var lastName string;
  fmt.Scanln(&lastName);

  fmt.Println("Your first name is " + firstName);
  fmt.Println("Your last name is " + lastName);

  fmt.Println("Welcome, " + firstName + " " + lastName);

  fmt.Println("Enter an integer: ");

  var firstNumber string;
  fmt.Scanln(&firstNumber)

  fmt.Println("Enter a second integer: ");

  var secondNumber string;
  fmt.Scanln(&secondNumber);

  var firstInt int
  var secondInt int

  firstInt, _ = strconv.Atoi(firstNumber);
  secondInt, _ = strconv.Atoi(secondNumber);
  fmt.Println(firstInt + secondInt); 
}
