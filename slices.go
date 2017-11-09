package main

import "fmt"

func main() {

  // Normal array  
  netvalue := [3]int{1,2,3}
  fmt.Println(netvalue)
  fmt.Println(netvalue[0:2])

  // But netter is slice
  var s = []int{1,23,4}
  fmt.Println(s)

  //struct and slice
  sls := []struct {
     x int
     st bool
  }{
    {1, true},
    {2, false},
  }

  fmt.Println(sls)
  fmt.Println(sls[0])
  fmt.Println(sls[0].x)

  // Slice Range 
  sliceNet := []int{1,2,3,4,5,6,7}
  fmt.Println(sliceNet[:3])
  fmt.Println(sliceNet[2:]) 
}