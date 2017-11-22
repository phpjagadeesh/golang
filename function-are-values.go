// Function can pass as argument

package main

import "fmt"

type fn func(string)

func main() {
    test3 := test()
    test2(test3)
}

func test() string{
    return "JAG"
}

func test2(f fn) string{
  fmt.Println(f)
}