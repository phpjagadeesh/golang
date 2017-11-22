// channel

package main

import "fmt"

func main() {
    ch := make(chan int, 2)
    ch <- 1
    ch <- 2
    fmt.Println(<-ch)
    fmt.Println(<-ch)

    ch1 := make(chan int, 5)
    for i := 0; i < 4; i++ {
      ch1 <- 10 * i
    }
    for j := range ch1 {
    fmt.Println(j)
    } 
} 