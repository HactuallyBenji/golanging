package main

import (
   "os"
)

func main() {
   // delete directory and all contents
   os.RemoveAll("U:/github/golanging/my_directory")

   // create a directory
   err := os.Mkdir("U:/github/golanging/my_directory", 0755) // will throw an error if the directory exists
   // if there is an error then panic to fail program.
   if err != nil {
      panic(err)
   }

  // MkdirAll create a tree of directories
  err = os.MkdirAll("U:/github/golanging/my_directory/my_nested_directory", 0755)
  // if there is an error then panic to fail program.
  if err != nil {
     panic(err)
  }
}
