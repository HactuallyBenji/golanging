package main

import (
   "fmt"
   "time"
)

func main() {
   // display current time
   now := time.Now()
   fmt.Println("Current date and time:", now)

   // create custom time
   customTime := time.Date(
      2025, 05, 15, 15, 20, 00, 0, time.Local)
   fmt.Println("Custom date and time:", customTime)

   // comparisons
   fmt.Println("The custom time is before now:", customTime.Before(now))
   fmt.Println("The custom time is after now:", customTime.After(now))
   fmt.Println("The custom time is equal to now:", customTime.Equal(now))
}
