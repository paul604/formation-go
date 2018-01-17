package main

import (
	"fmt"
)

type developer interface {
	develop() string
}

const (
	jsHelloWorld string = `alert("Hello World!");`
	goHelloWorld string = `fmt.Printf("Hello, world!")`
	pythonHelloWrold string = `print("Hello, World!")`
)

type frontEndDeveloper struct{}

func (d frontEndDeveloper) develop() string {
	return jsHelloWorld
}

type backEndDeveloper struct{}

func (d backEndDeveloper) develop() string {
	return goHelloWorld
}

type fullStackDeveloper struct{}

func (d fullStackDeveloper) develop() string {
	return pythonHelloWrold
}

var jim = frontEndDeveloper{}
var joe = backEndDeveloper{}
var jules = fullStackDeveloper{}

func main() {
	developers := [3]developer{jim, joe, jules}

	for _, d := range developers {
		// To complete using type assertion:
		// switch ... {
		//	case ...: fmt.Printf("That's written in [language]: ")
		// }
		var msg string;
		switch d.(type) {
		case frontEndDeveloper:
			msg = "That's written in JS : "
		case backEndDeveloper:
			msg = "That's written in GO : "
		case fullStackDeveloper:
			msg = "That's written in Python : "
		default:
			msg ="error"
		}
		fmt.Println(msg + d.develop())
	}
}
