// interface

package main

import "fmt"

type Address struct {
    first, second string
}

type list interface {
    Display()
}

func (A Address) Display() {
    fmt.Println(A.first , A.second)
}

func main() {
  i := Address{"first", "second"}
  i.Display()
}