package main

import (
  "fmt"
  "errors"
  "math/rand"
  "time"
  "strconv"
)

func findMaxValue(array []int) (int, error) {

  if len(array) == 0 {
    return 0, errors.New("Empty array received. No max value exists") 
  }

  currentMax := array[0]

  for i := 1; i < len(array); i++ {
    if array[i] > currentMax {
      currentMax = array[i]
    }
  }

  return currentMax, nil
}

func findMaxIndex(array []int) (int, error) {
  
  if len(array) == 0 {
    return 0, errors.New("Empty array received. No max value exists")
  }

  maxValueIndex := 0
  currentMax := array[maxValueIndex]

  for i := 0; i < len(array); i++ {
    if array[i] > currentMax {
      currentMax = array[i]
      maxValueIndex = i
    }
  }

  return maxValueIndex, nil
}

func findMinValue(array []int) (int, error) {

  if len(array) == 0 {
    return 0, errors.New("Empty array received. No min value exists") 
  }

  currentMin := array[0]

  for i := 1; i < len(array); i++ {
    if array[i] < currentMin {
      currentMin = array[i]
    }
  }

  return currentMin, nil
}

func findMinIndex(array []int) (int, error) {
  
  if len(array) == 0 {
    return 0, errors.New("Empty array received. No min value exists")
  }

  minValueIndex := 0
  currentMin := array[minValueIndex]

  for i := 0; i < len(array); i++ {
    if array[i] < currentMin {
      currentMin = array[i]
      minValueIndex = i
    }
  }

  return minValueIndex, nil
}

func getPositiveValues(array []int) ([]int, error) {
  
  if len(array) == 0 {
    return make([]int, 0), errors.New("Empty array has no positive values")
  }

  var positiveValues []int

  for i := 0; i < len(array); i++ {
    if array[i] > 0 {
      positiveValues = append(positiveValues, array[i])
    }
  }

  return positiveValues, nil
}

func getNegativeValues(array []int) ([]int, error) {
  
  if len(array) == 0 {
    return make([]int, 0), errors.New("Empty array has no positive values")
  }

  var negativeValues []int

  for i := 0; i < len(array); i++ {
    if array[i] < 0 {
      negativeValues = append(negativeValues, array[i])
    }
  }

  return negativeValues, nil
}

func findLongestSortedSubsequence(array []int) ([]int, error) {
  
  if len(array) == 0 {
    return make([]int, 0), errors.New("Empty array has no subsequence")
  }

  if len(array) == 1 {
    return array, nil
  }

  var longestIncreasingSubsequence []int
  var longentDecreasingSubsequence []int

  previousValue := array[0]

  for i := 0; i < len(array); i++ {
    
  }
}

func main() {
  ns := rand.NewSource(time.Now().UnixNano())
  generator := rand.New(ns)

  var arrayLength string

  fmt.Println("How many random values would you like?")
  fmt.Scan(&arrayLength)

  length, _ := strconv.Atoi(arrayLength)

  var randomArray []int

  for i := 0; i < length; i++ {
    randomArray = append(randomArray, generator.Intn(200) - 100)
  }

  fmt.Println(randomArray)

  maxValue, _ := findMaxValue(randomArray)
  fmt.Println("The maximum value in the array is:", maxValue)

  maxValueIndex, _ := findMaxIndex(randomArray)
  fmt.Println("The maximum value's index in the array is:", maxValueIndex)

  minValue, _ := findMinValue(randomArray)
  fmt.Println("The minimum value in the array is:", minValue)

  minValueIndex, _ := findMinIndex(randomArray)
  fmt.Println("The minimum value's index in the array is:", minValueIndex)

  positiveValues, _ := getPositiveValues(randomArray)
  fmt.Println("The positive values in the array are:", positiveValues)

  negativeValues, _ := getNegativeValues(randomArray)
  fmt.Println("The negative values in the array are:", negativeValues)
 }
