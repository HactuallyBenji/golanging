package main

import (
   "os"
   "io/ioutil"
   "fmt"
)

func main() {
   // delete directory and all contents
   // defer to end of program
   defer os.RemoveAll("U:/github/golanging/test_directory")

   // create a directory
   err := os.Mkdir("U:/github/golanging/test_directory", 0755) // will throw an error if the directory exists
   // if there is an error then panic to fail program.
   if err != nil {
      panic(err)
   }

  // MkdirAll create a tree of directories
  err = os.MkdirAll("U:/github/golanging/test_directory/nested_directory", 0755)
  // if there is an error then panic to fail program.
  if err != nil {
     panic(err)
  }

  // list content of a directory
  content, err := ioutil.ReadDir("U:/github/golanging")
  // if there is an error then panic to fail program.
  if err != nil {
    panic(err)
  }
  // iterate through content
  for _, item := range content {
    fmt.Println(" ", item.Name(), item.IsDir())
  }
}
