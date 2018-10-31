package main

import (
  "os"
  "fmt"
  "net"

)

func main() {
  conn, err := net.Dial("tcp", "ascii.jp:80")
  if err != nil {
    panic(err)
  }
  io.WriteString(conn, "GET / HTTP1.0\r\nHost: ascii.jp\r\n\r\n")
  io.Copy(os.Stdout, conn)
}
