package main

import (
   "fmt"
   "os"
)

func main() {
   f, err := os.Open("U:/github/golanging/README.md")
   
   fmt.Println(err)

   //  create a slice of bytes
   b1 := make([]byte, 5)  
   data, err := f.Read(b1) 

   // feedback message in case of error; otherwise nil
   fmt.Println(err) // print error due to reading from the file. nil is no errors

   // display the slice
   fmt.Printf(string(b1[:data]))

   // close the file after completing the operations
   f.Close()
}
