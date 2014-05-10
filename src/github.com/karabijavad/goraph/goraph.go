package main

import (
  "os"
  "bufio"
)

type Property struct {
  id int
  data []byte
}

func (p *Property) save() error {
  fo, _ := os.Open("properties.store")
  defer func() { fo.Close() }()

  w := bufio.NewWriter(fo)

  _, err := w.Write(p.data[:5]);

  return err
}

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

func main() {
  n := Node{id: 1}
  p := Property{data: []byte("this is the test data")}

  n.properties = []Property{p}

  n.save()
}
