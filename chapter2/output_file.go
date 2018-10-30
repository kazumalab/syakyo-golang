package main

import {
  "os"
}

func main() {
  file, err := os.Create("test.txt")
  if err != nil {
    panic(err)
  }
  file.write([]byte("os.File example\n"))
  file.Close()
}
