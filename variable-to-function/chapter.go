package main

import "fmt"


type Person struct {
    fname string
    nname string
}

func (P Person) outer() {
  fmt.Println("out function", P)
}

func main() {
    
    fmt.Println(`good`)
    p := Person{"Kape", "Town"}
    fmt.Println(p.fname)

    mapVal := map[string]int{
        "new": 10,
        "old": 9,
    }
    fmt.Println(mapVal["new"])

    p.outer()
}