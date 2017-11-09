package main

import "fmt"

type NewFiled struct{
    x int
    y bool
} 

func main() {

  fmt.Println(NewFiled{1,true})

  v := NewFiled{2,false}
  fmt.Println(v.x)
}