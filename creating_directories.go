package main

import (
   "os"
)

func main() {
   // delete directory and all contents
   os.RemoveAll("U:/github/golanging/test_directory") 

   // create a directory
   err := os.Mkdir("U:/github/golanging/test_directory", 0755) // will throw an error if the directory exists

   // if there is an error then panic to fail program.
   if err != nil {
      panic(err)
   }
}
