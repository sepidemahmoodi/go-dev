package main

// map ha built in method hayi hastan ke ma mitonim index va value hashono khodemon tarif konim.az har noei mitone bashe 
//masalan inja index ha string va value ha struct person hastan.
// hatman bayad az make estefade konim ke map ha ready to use beshan.in kheyli moheme.

import (
	"fmt"
	"bufio"
	"os"
)

// type Person struct {
// 	Age int
// 	name string
// }

// var a map[string]Person

// func main () {
	
// 	a = make(map[string]Person)
// 	a["sepide"] = Person{28, "sepide"}

// 	fmt.Println(a["sepide"])
	
// }

//duplicate lines program

// func main () {
// 	counts := make(map[string]int)
// 	input := bufio.NewScanner(os.Stdin)
// 	for input.Scan() {
// 		counts[input.Text()]++
// 	}

// 	for line, n := range counts {
// 		if (n > 1) {
// 			fmt.Println(line, n)
// 		}
// 	}
// }

//dup2 read from file

func main () {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _,arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}