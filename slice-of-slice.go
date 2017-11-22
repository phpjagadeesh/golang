package main

import "fmt"

func main() {
    slice1 := [][]string{
        []string{"1","2"},
        []string{"3","4"},
    }

    fmt.Println(slice1) 

    // Append New Element to slice
    var newSlice []int
    s:= append(newSlice, 3)
    fmt.Println(s)   
}