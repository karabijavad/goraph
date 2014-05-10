package main

import (
  "github.com/karabijavad/goraph/node"
)



func main() {
  n := Node{id: 1}
  n.setProperty("javad")

  n.save()
}

