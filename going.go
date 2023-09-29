/*
  Ben Miller
  2023-09-28
*/

package main

import "fmt"

func main() {
  fmt.Println("Enter your first name: ");
  
  var firstName string;
  fmt.Scanln(&firstName);

  fmt.Println("Enter your last name: ");

  var lastName string;
  fmt.Scanln(&lastName);

  fmt.Println("Your first name is " + firstName);
  fmt.Println("Your last name is " + lastName);

  fmt.Println("Welcome, " + firstName + " " + lastName);
}
