package main

import (
   "fmt"
   "sort"
)

func main() {
   // define a slice
   words := []string{"camel", "zebra", "horse", "dog", "elephant", "giraffe"}
   fmt.Println("Original slice:", words)
   fmt.Println("The original values are sorted:", sort.StringsAreSorted(words))

   // sort the values in the slice
   sort.Strings(words)
   fmt.Println("Sorted slice:", words)
   fmt.Println("The values are sorted:", sort.StringsAreSorted(words))
}
