package main

import (
   "fmt"
   "os"
)

func main() {
   data := "Hello, world!"
   fmt.Println("data string:", data)

   // create a new file
   f, err := os.Create("U:/github/golanging/myCreatedFile.txt")

   // display the new file
   fmt.Println(err)
   fmt.Println("new file:", f)


   // write the string to the new file
   n, err := f.WriteString(data)
   fmt.Println(err)
   fmt.Println("characters in file:", n)

   // close the file
   f.Close()
}
