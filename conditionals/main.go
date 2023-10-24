package main

import (
	"fmt"
	"os"
	"log"
)

// var (
// 	heads, tails int
// )

// func coinFlip(face string) {
// 	switch face {
//     case "heads":
//       heads++

//     case "tails": 
//       tails++

//     default: 
//       fmt.Println("undefined")
// 	}
// }

func main () {
	a, b :=  12, 14

	if a > b {
		fmt.Println("a is greater than b")
	} else if a < b {
		fmt.Println("b is greater than a")
	} else {
		fmt.Println("a and b is equal")
	}

  x, y :=  120, 11
  switch {
  case x > y:
    fmt.Println("x is greater than y")
  case x < y:
    fmt.Println("y is greater than x")
  default: 
    fmt.Println("x and y is equal")
  }

	file, err := os.Open("hello.txt")

	if err != nil {
		log.Panic(err)	
	}

	data := make([]byte, 100)
	if _, err := file.Read(data); err != nil {
		log.Panic(err)
	}

	fmt.Println(string(data))
}