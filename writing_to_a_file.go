package main

import (
   "fmt"
   "io/ioutil"
)

func main() {
   // create a slice of bytes (utf-8 code) from an input string
   data := []byte("Hello, world!")

   fmt.Println(data) // display the slice

   err := ioutil.WriteFile("U:/github/golanging/test.txt", data, 0644) 
   fmt.Println(err)

   new_file, err := ioutil.ReadFile("U:/github/golanging/test.txt")

   // feedback message in case of error
   fmt.Println(err)

   // convert the file contents to a string and display it
   fmt.Print(string(new_file))
}
