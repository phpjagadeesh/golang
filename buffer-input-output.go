package main

import (
  "fmt"
  "bufio"
  "os"
)

func main() {
  scaner := bufio.NewScanner(os.Stdin)
  fmt.Println(scaner)

  for scaner.Scan() {
    fmt.Println(scaner.Text())
  }
}