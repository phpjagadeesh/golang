// Unlike other language no paranthesis needed

package main 

import "fmt"

func main() {
    for i := 1; i < 10;
    i++ {
        fmt.Println("Incement ", i);
    }

    // init and post statement are option
    // But conditional are manadetory
    j := 1;      
    for ; j < 10 ; {
      fmt.Println(j)
      j  = j + 1 
    }

    // Bye Bye to 'while'
    k := 10;
    for k <= 20 {
        fmt.Println(k)
        k++
    }
}