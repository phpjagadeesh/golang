package main

import "fmt"

func do(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Println("integer", v)
    } 
}

func main() {
  do(12)
}