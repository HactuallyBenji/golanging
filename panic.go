package main

import (
   "os"
)

func main() {
   _, err := os.Create("/erronous_path/file.txt")
   if err != nil {
      panic(err)
   }
}
