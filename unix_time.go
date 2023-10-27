package main

import (
   "fmt"
   "time"
)

func main() {
   now := time.Now()
   unixtime := now.Unix() // unix time
   unixnanotime := now.UnixNano() // unix nano time

   fmt.Println(now)
   fmt.Println(unixtime)
   fmt.Println(unixnanotime)
}
