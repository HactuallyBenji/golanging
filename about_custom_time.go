package main

import (
   "fmt"
   "time"
)

func main() {
   // create custom time
   customTime := time.Date(
      2025, 05, 15, 15, 20, 00, 0, time.Local)
   fmt.Println("Custom date and time:", customTime)

   fmt.Println("Custom year:", customTime.Year())
   fmt.Println("Custom month:", customTime.Month())
   fmt.Println("Custom day:", customTime.Day())
   fmt.Println("Custom weekday:", customTime.Weekday())
   fmt.Println("Custom hour:", customTime.Hour())
   fmt.Println("Custom minute:", customTime.Minute())
   fmt.Println("Custom second:", customTime.Second())
   fmt.Println("Custom nanosecond:", customTime.Nanosecond())
   fmt.Println("Custom location:", customTime.Location())
}
