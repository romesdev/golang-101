package main

import "fmt"

func main(){
  // slice
  // slice couldn't be compares
	colors := []string{"black", "green", "red", "yellow", "white"}
  fmt.Println(colors, len(colors), cap(colors))

  // how to init a slice
  var a []int32
  b := []int32{1, 5, 3, 6}
  c := make([]int32, 4) // define length
  d := make([]int32, 0, 10) // define length and capacity
  fmt.Println(d)

  // at this moment the capacity is doubled
  colors = append(colors, "blue")
  fmt.Println(colors, len(colors), cap(colors))

  // arrays
  // the array length belongs to his type
  // how to init a slice
  a := [3]int{1, 5, 20}
  b := [...]int{1, 5, 20}
  // array could be compares
  fmt.Println(a == b) // output: true

  // map {key: value}
  // unsorted
  ages := make(map[string]uint8)
  ages["john"] = 42
  ages["doe"] = 20
  ages["peter"] = 30
  fmt.Println(ages)

  // not found key
  fmt.Println(ages["key"]) // return zero from type of data

  // good way to retrieve data (or not)
  value, ok := ages["john"]
  fmt.Println(value, ok)
  age, status := ages["key"]
  fmt.Println(age, status)
}
