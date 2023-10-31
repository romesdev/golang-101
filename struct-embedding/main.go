package main

import "fmt"

type Base struct {
    x int
}

func (b Base) describe() string {
    return fmt.Sprintf("base with b=%d", b.x)
}

type Container struct {
    Base
    y int
}

func main() {

    // one way to initialize
    // co := Container{
    //     Base: {
    //         x: 1,
    //     },
    //     y: 4,
    // }

    co := Container{}
    co.x = 0
    co.y = 42

    fmt.Printf("co={x: %d, y: %d}\n", co.x, co.y)

    fmt.Println("also works:", co.Base.x)

    fmt.Println(co.describe())
}
