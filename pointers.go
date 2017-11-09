package main

import "fmt"

func main() {
  var p *int
  var i int = 10;
  p = &i
  fmt.Println(p)
  fmt.Println(*p)
}