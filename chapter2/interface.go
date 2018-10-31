package main

import (
  "fmt"
)

type Talker interface {
  Talk()
}

type Greeter struct {
  name String
}

func (g Greeter) Talk() {
  fmt.Printf("Hello, my name is %s\n", g.name)
}

func main() {
  var talker Talker
  talker = &Greeter{"kazuma"}
  talker.Talk()
}
