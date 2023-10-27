package main

import (
   "bufio"
   "fmt"
   "os"
)

func main() {
   f, err := os.Open("U:/github/golanging/README.md")
   br := bufio.NewReader(f) // create a buffered reader
   data, err := br.Peek(5) // read 5 bytes
   fmt.Println(err)

   // display the peeked data
   fmt.Printf(string(data))
}
