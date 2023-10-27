package main

import (
   "fmt"
   "time"
)

func main() {
   // display current time
   now := time.Now()
   fmt.Println("Today's date and time:", now)
   fmt.Println("Current year:", now.Year())
   fmt.Println("Current month:", now.Month())
   fmt.Println("Current day:", now.Day())
   fmt.Println("Current hour:", now.Hour())
   fmt.Println("Current minute:", now.Minute())
   fmt.Println("Current second:", now.Second())
   fmt.Println("Current nanosecond:", now.Nanosecond())
   fmt.Println("Current location:", now.Location())
   fmt.Println("Current weekday:", now.Weekday())
}
