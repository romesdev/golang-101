package main

import "fmt"

func main () {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

  j := 0
  for j <= 4 {
    fmt.Println(j)
    j++
  }

  // this is a slice
  // has a dynamic length
  colors := []string{"black", "green", "red", "yellow", "white"}

  for i := 0; i < len(colors); i++ {
    fmt.Println(colors[i])
  }

  for k, color := range colors{
    fmt.Println(k, ":",  color)
  }
}
