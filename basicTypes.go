package main

import "fmt"
import "math/cmplx"

func main() {
  var number complex128;
  number = cmplx.Sqrt(-10+21i)
  fmt.Println(number)
}