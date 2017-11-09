package main

import "fmt"

func main() {
  fmt.Println(add(1));
}

func add(x int)(k, m int, name string) {
    k = x + 10
    m = k + 10
    name = "Name is ok"
    return 
}