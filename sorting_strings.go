package main

import (
   "fmt"
   "sort"
)

func main() {
   // define a slice
   words := []string{"camel", "zebra", "horse", "dog", "elephant", "giraffe"}
   fmt.Println("Original slice:", words)

   // sort the values in the slice
   sort.Strings(words)
   fmt.Println("Sorted slice:", words)
}
