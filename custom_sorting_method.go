package main

import (
   "fmt"
   "sort"
)

// create an alias type
type mytype []string

// implement the Len method
func (s mytype) Len() int {
   return len(s)
}

 // implement the Swap method
func (s mytype) Swap(i, j int) {
   s[i], s[j] = s[j], s[i]
}

// implement the Less method
func (s mytype) Less(i, j int) bool {
   return len(s[i]) < len(s[j])
}

func main() {
   // create a slice of strings
   fruits := []string{"pear", "passionfruit", "mango", "banana", "fig"}
   fmt.Println("Original slice:", fruits)

   // create a mytype variable
   myfruits := mytype(fruits)

   sort.Sort(myfruits)
   fmt.Println("Sorted by length:", myfruits) 
}
