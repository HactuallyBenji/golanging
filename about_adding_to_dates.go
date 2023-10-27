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

   // time operations
   // subtract two times to returns a duration in hours, minutes, seconds
   diff := now.Sub(customTime)
   fmt.Println("Time between now and custom time:", -diff)

   fmt.Println(customTime.Add(diff))
   fmt.Println(customTime.Add(-diff))
}
