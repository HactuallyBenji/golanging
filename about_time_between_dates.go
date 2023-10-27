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

   // subtract two times to returns a duration in hours, 
   // minutes, seconds
   diff := now.Sub(customTime)
   fmt.Println("Time between now and custom time:", -diff)

   fmt.Println("Hours between now and custom time:", -diff.Hours())
   fmt.Println("Minutes between now and custom time:", -diff.Minutes())
   fmt.Println("Seconds between now and custom time:", -diff.Seconds())
   fmt.Println("Nanoseconds between now and custom time:", -diff.Nanoseconds())
}
