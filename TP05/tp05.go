package main

import "fmt"

func count(from, to int) {
        defer fmt.Println(from)
        fmt.Println(from)
        if from < to {
                count(from + 1, to)
        } else if (from == to) {
                panic("done!")
        }
}

func main() {
        defer func() {
           fmt.Print(recover())
        }()
        count(0, 10)
        //fmt.Println("done!")
        
}
