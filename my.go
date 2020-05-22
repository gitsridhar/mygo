package main

import "fmt"
import "strconv"

func main() {
    fmt.Println("hello world")
    for i:= 0; ; i++ {
        fmt.Println("Doing it again " + strconv.Itoa(i))
    }
}

