package main

import "fmt"

func main() {
    var message string // declare the variable
    message = "Hello, World!" // assign a value to the variable

    fmt.Println(message)

    var message2, email string // declare two string variables
    message2 = "Hello, World!" // assign a value to one variable
    email = "john@john.com" // assign a value to the second variable

    fmt.Println(message2)
    fmt.Println(email)


    // initialize two string variables in the same statement
    var message3, email2 string = "Hello, World!", "john@john.com"

    fmt.Println(message3)
    fmt.Println(email2)
}
