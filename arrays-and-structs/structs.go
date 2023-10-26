package main

import "fmt"

// create a new struct
type Person struct {
  firstName string
  lastName string
  age uint8
  status bool
}

type Player struct {
  firstName string
  lastName string
  age uint8
  team string
}

// reminder: the first letter
// in uppercase meaning that
// the attribute could be accessed external
type Category struct {
  Name string
  Parent *Category
}

func (category Category) hasParent() bool {
  return category.Parent != nil
}

func (category *Category) setParent(parent *Category) {
  category.Parent = parent
}

// rewrite method String()
func (person Person) String() string {
  return fmt.Sprintf("%s %s has %d years old", person.firstName, person.lastName, person.age)
}


func main(){
  person := Person{
    firstName: "Bruno",
    lastName: "Berle",
    age: 29,
    status: true,
  }

  fmt.Println(person.firstName, person.lastName, person.age, person.status)
  fmt.Println(person)

  players := []Player{
    {
      firstName: "Franz",
      lastName: "Wagner",
      age: 22,
      team: "Orlando Magic",
    },
    {
      firstName: "Nikola",
      lastName: "Jokić ",
      age: 28,
      team: "Denver Nuggets",
    },
    {
      firstName: "Devin",
      lastName: "Booker",
      age: 26,
      team: "Phoenix Suns",
    },
    {
      firstName: "Luka",
      lastName: "Dončić",
      age: 24,
      team: "Dallas Mavericks",
    },
  }

  fmt.Println(players)

  newPlayer := Player{
      firstName: "Giannis",
      lastName: "Antetokounmpo",
      age: 28,
      team: "Milwaukee Bucks",
  }

  players = append(players, newPlayer)

  fmt.Println(players)

  shirt := Category{
    Name: "shirt",
    Parent: nil,
  }

  tShirt := Category{
    Name: "t-shirt",
    Parent: &shirt, // pass as reference
  }

  fmt.Println(shirt.hasParent())
  fmt.Println(tShirt.hasParent())

  clothes := Category{
    Name: "clothes",
    Parent: nil,
  }

  shirt.setParent(&clothes)
  fmt.Println(shirt)
  fmt.Println(shirt.hasParent())
}
