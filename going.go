/*
  Ben Miller
  2023-09-28
*/

package main

import (
  "fmt"
  "strconv"
  "math/rand"
  "time"
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
  fmt.Println(firstInt)
  fmt.Println(secondInt)
  fmt.Println(firstInt + secondInt); 

  fmt.Println("\n---- Types ----");
  var a int = 10;
  var b int32 = 20;
  var c byte = 15;
  var d float32 = 0.05

  fmt.Println("30: ", int(a) + int(b))
  fmt.Println("35: ", int32(b) + int32(c))
  fmt.Println("25: ", int(a) + int(c))
  fmt.Println("10.05: ", float32(a) + float32(d))
  fmt.Println("1: ", float32(b) * float32(d))
  fmt.Println("300: ", float32(c) / float32(d))


  var mask uint16 = 0xffff;
  var value uint16 = 38;
  fmt.Printf("%08b", mask & value)

  fmt.Println("\n---- Randomness ----")
  ns := rand.NewSource(time.Now().UnixNano())
  generator := rand.New(ns)

  fmt.Println(generator.Intn(100))
  fmt.Println(generator.Intn(100))
  
  fmt.Println(7 * 5 / 3);
}
