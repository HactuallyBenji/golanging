package main

import (
   "fmt"
   "sort"
)

func main() {
   // define a slice
   numbers := []int{67, 18, 62, 60, 25, 64, 75, 5, 17, 55}
   fmt.Println("Original Numbers:", numbers)

   // use the sort.Ints function to sort the values in the slice
   sort.Ints(numbers)

   fmt.Println("Sorted Numbers:", numbers)
}
