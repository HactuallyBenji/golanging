package main

import (
  "fmt"
  "regexp"
)

func main() {
  testString1 := "Hello there"
  testString2 := "123-123-123"
  testString3 := "My ID is 123"
  testString4 := "Testing123"
  testString5 := "Helloinnn"
  testString6 := "Helloi"

  numbersAndLetters, _ := regexp.Compile("^[A-Za-z0-9]*$")

  fmt.Println("Letters and numbers only")

  fmt.Println(numbersAndLetters.MatchString(testString1))
  fmt.Println(numbersAndLetters.MatchString(testString2))
  fmt.Println(numbersAndLetters.MatchString(testString3))
  fmt.Println(numbersAndLetters.MatchString(testString4))

  fmt.Println("\ni then any number of n's")
  // strings that include the letter i followed by zero or more instances of the letter n. 
  iThenMaybeN, _ := regexp.Compile("i[n]*")

  fmt.Println(iThenMaybeN.MatchString(testString1))
  fmt.Println(iThenMaybeN.MatchString(testString2))
  fmt.Println(iThenMaybeN.MatchString(testString5))
  fmt.Println(iThenMaybeN.MatchString(testString6))

  fmt.Println("\ni then at least one n")
  // all strings that include the letter i followed by one or more instances of the letter n.

  iThenSomeN, _ := regexp.Compile("i[n]+")
  fmt.Println(iThenSomeN.MatchString(testString1))
  fmt.Println(iThenSomeN.MatchString(testString2))
  fmt.Println(iThenSomeN.MatchString(testString5))
  fmt.Println(iThenSomeN.MatchString(testString6))
} 
