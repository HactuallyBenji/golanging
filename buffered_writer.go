package main

import (
   "bufio"
   "fmt"
   "os"
   "io/ioutil"
)

func main() {
   data :="Hello, world!!!"
   fmt.Println("original string:", data)

   // create a new file
   f, err := os.Create("U:/github/golanging/buffered_writing.md")
   fmt.Println(err)

   // create a buffered writer that we can use to write data to the new file
   bw := bufio.NewWriter(f)

   // write the data to the buffer writer
   n, err := bw.WriteString(data)

   // display the number of bytes written
   fmt.Println("bytes written:", n)

   // flush flushes/submits the data to the underlying io.Writer
   bw.Flush()

   newFile, err := ioutil.ReadFile("U:/github/golanging/buffered_writing.md")
   // feedback message in case of error
    fmt.Println(err)
   // convert the file contents to a string and display it
   fmt.Print("file contents:", string(newFile))

   f.Close()
}
