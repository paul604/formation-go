package main

import (
	"fmt"
)

type book struct {
	title, author string
        year int
}

var bookList = []book {
	{"Computing Machinery and Intelligence", "Alan Turing", 1950},
	{"A relational model for large shared data banks", "Edgar F. Codd", 1970},
	{"The C Programming Language", "Brian Kernighan", 1978},
	{"Design Patterns", "Erich Gamma", 1994},
	{"The Art of Computer Programming", "Donald E. Knuth", 1968},
	{"Gödel, Escher, Bach: An Eternal Golden Braid", "Douglas R. Hofstadter", 1999},
	{"Domain-Driven Design: Tackling Complexity in the Heart of Software", "Eric Evans", 2003},
	{"Hackers, Heroes of the Computer Revolution", "Steven Levy", 1984},
	{"Modern Operating Systems", "Andrew S. Tanenbaum", 1992},
	{"The Algorithm Design Manual", "Steven Skiena", 1997},
}

func main() {
	selectedAuthor := "Andrew S. Tanenbaum"
	var selectedBooks []book
	for _, b := range bookList {
		// Add books of selectedAuthor to
		// list selectedBooks
                if isBookAuthor(b, selectedAuthor) {
                        selectedBooks = append(selectedBooks, b)
                }
	}
	fmt.Printf("%+v\n", selectedBooks)
}

func isBookAuthor(b book, author string) bool {
	// Check if author matches book's author
        return b.author==author
}
