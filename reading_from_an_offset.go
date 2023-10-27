package main

import (
   "fmt"
   "os"
)

func main() {
   f, err := os.Open("U:/github/golanging/README.md")
   fmt.Println(err)

   // skip the first 100 bytes
   s, err := f.Seek(100, 0)
   fmt.Println(err) // nil if no isssues

   // display the offset
   fmt.Println(s)

   // read 5 bytes starting from the offset
   data := make([]byte, 20)
   n, err := f.Read(data)

   fmt.Println(err) // nil if no issues
   fmt.Println("Bytes read", n)
   fmt.Println("Reading starting from byte", s, ":", string(data[:n]))

   // close the file
   f.Close()
}
