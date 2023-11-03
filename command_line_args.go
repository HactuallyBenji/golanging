package main

import (
   "fmt"
   "os"
)

func main() {
   args := os.Args
   programName := args[0]
   arguments := args[1:]
   fmt.Println(programName)
   fmt.Println(arguments)
}
