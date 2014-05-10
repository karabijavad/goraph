package main

import (
  "github.com/karabijavad/goraph/property"
)

type Node struct {
  id  int
  properties []Property
}

func (n *Node) save() bool {
  result := true
  for _, prop := range n.properties {
    err := prop.save()
    if err != nil { result  = false }
  }
  return result
}

func (n *Node) setProperty(string data) bool {
  result := true
  p := Property{data: []byte(data)}
  return result
}
