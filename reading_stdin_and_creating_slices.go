package main

import (
  "fmt"
)

func main() {
  
  var strings [15] string 
  var stringLengths [] int 
  totalStringLengths := 0.0

  fmt.Println("Give me 12 words")
  fmt.Println("Heruehetonuhnt")

  for i := 0; i < 12; i++ {
    fmt.Print("Word ", i + 1, ":")
    fmt.Scan(&strings[i])
    stringLengths = append(stringLengths, len(strings[i]))
    totalStringLengths += float64(len(strings[i]))
  }

  avgStringLength := totalStringLengths / 12;

  fmt.Print(avgStringLength)

  var smallerThanAverageStrings [] string
  var largerThatAverageStrings [] string

  for i := 0; i < 12; i++ {
    if (stringLengths[i] < avgStringLength) {
      smallerThanAverageStrings = append(smallerThanAverageStrings, strings[i]) 
    } else if (stringLengths > avgStringLength) {
      
    }
  }
  
}
