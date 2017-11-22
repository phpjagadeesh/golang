// A method is function has defined reciever

package main

import "fmt"


type User struct {
    first, second string
}

func (u User) Greeting() string {
    return fmt.Sprintf("%s\n%s", u.first, u.second)
}

func main() {
    u := User{"first", "second"}
    fmt.Println(u.Greeting())
} 
