package main

import "fmt"

func main() {
    var i bool = true
    if i == true {
        fmt.Println("BOOL Value")
    }
   
    var firstValue  = first(1)
    fmt.Println(firstValue)
    var secondValue = second(100)
    fmt.Println(secondValue)

     
    // Switch
    switch(secondValue) {
      case 1:
      fmt.Println("First")
      case 199:
      fmt.Println("Second")  
    }

 }

func first(j int)(int) {
  if k := j ; j < 10 {
   return k
  }
  return 0
}

func second(m int)(int) {
  if n := m ; m < 100 {
   return n
  }
  return 199     
}