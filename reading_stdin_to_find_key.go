package main

import (
  "fmt"
)

func main() {
  hasFoundKey := false

  wordFrequencies := map[string]int {
    "hello": 12,
    "world": 22,
    "android": 3,
    "Golang": 63,
    "Google": 32,
  }

  for !hasFoundKey {
    
    var userRequestedKey string   

    fmt.Println("Please enter a key")
    fmt.Scan(&userRequestedKey)

    wordFrequency, isInMap := wordFrequencies[userRequestedKey]

    if isInMap {
      hasFoundKey = true
      fmt.Println("The value at the requested key is:", wordFrequency)
    } else {
      fmt.Println("We could not find the requested key")
    }

  }
}
