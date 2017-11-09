package main

import "fmt"

func main() {
  defer first()
  second()
  third()
  for i := 1; i <= 4 ; i++ {
    defer fifth()
    fourth()
  }
}

func first() {
  fmt.Println("First")    
}

func second() {
  fmt.Println("Second")  
}

func third() {
  fmt.Println("Third")    
}

func fourth() {
  fmt.Println("Fourth")    
}

func fifth() {
  fmt.Println("Fifth")    
}