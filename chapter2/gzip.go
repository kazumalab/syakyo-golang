package main

import (
  "compress/gzip"
  "os"
)

func main() {
  file, err := os.Create("text.txt.gz")
  if err != nil {
    panic(err)
  }
  writer := gzip.NewWriter(file)
  writer.Header.Name = "test.text"
  io.WriteString(writer, "gzip.Writer example\n")
  writer.Close()
}
