package main

import "fmt"

func main() {
    pow := []int{1,2,3,4,5,6,7}

    for i, v := range pow {
        fmt.Println(i,v)
    }

    second := []int{40,39,30,50}
    for _,v := range second {
        fmt.Println(v)
    }
}