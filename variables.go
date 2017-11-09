package main 

import "fmt"

func main() {
  var a bool
  var st string
  var num int
  fmt.Println(a,st,num)

  var initvalue string = "Intialisation"
  fmt.Printf("%s\n", initvalue)

  // Short variable declaration
  // Is only avilable inside function
  // Outside not available
  name := "short variable"
  fmt.Println(name)

  // Multiple assign
  c, num := "multiple asign", 100
  fmt.Println(c, num) 

}