package property

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

  w := bufio.NewWriter(fo)

  _, err := w.Write(p.data[:5]);
  fo.Close()
  return err
}

