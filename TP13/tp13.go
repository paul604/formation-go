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
	pythonHelloWorld string = `print("Hello, World!")`
)

//<editor-fold desc="frontEndDeveloper">
type frontEndDeveloper struct {

}

func (frontEndDeveloper) develop() string {
	return jsHelloWorld
}
//</editor-fold>

//<editor-fold desc="backEndDeveloper">
type backEndDeveloper struct {

}

func (backEndDeveloper) develop() string {
	return goHelloWorld
}
//</editor-fold>

//<editor-fold desc="fullStackDeveloper">
type fullStackDeveloper struct {

}

func (fullStackDeveloper) develop() string {
	return pythonHelloWorld
}
//</editor-fold>

var jim = frontEndDeveloper{}// To complete: jim is a frontEndDeveloper
var joe = backEndDeveloper{}// To complete: joe is a backEndDeveloper
var jules = fullStackDeveloper{}// To complete: jules is fullStackDeveloper

func main() {
	developers := [3]developer{jim, joe, jules}

	for _, d := range developers {
		fmt.Println(d.develop())
	}
}
