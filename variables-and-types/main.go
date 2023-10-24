package main

import "fmt"

// simple variable declaration
var hello string

// declare multiple variables
var (
  foo float32
  Bar int // available out of the package 
)

func main () {
  hello = "hello world"
  fmt.Println(hello)


  foo = 42.221
  Bar = 221
  fmt.Println(foo, Bar)

  // omit the type with ':=' operator
  message := "golang"
  fmt.Println(message)
  
  var x, y, z = true, 123, "foo"
  y++
  Bar = Bar + y

  fmt.Println(x, y, z)
  fmt.Println(Bar)

  a, b := 101, 42
  fmt.Println(a, b)
  a, b = b, a
  fmt.Println(a, b)
}