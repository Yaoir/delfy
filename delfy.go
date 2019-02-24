package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
	)

// data file, containing "then" clauses
const datafile = "delfy.lib"

// standard input
var stdin *bufio.Reader

// datafile in a slice of strings
var then []string

/* replacement for C library random() function */

var randomgen *rand.Rand

func srandom() {
//
        randomgen = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func random() int {
//
        return randomgen.Int()
}

func init_delfy() {
	
	// init random number generator
	srandom()

	// open standard input
	stdin = bufio.NewReader(os.Stdin)

// TODO: standard location for data file. /usr/local/share/... ?
// option: specify data file location

	inputfile, err := os.Open(datafile)

	if err != nil {
		fmt.Fprintf(os.Stderr,"Cannot open %s\n",datafile);
		os.Exit(2)
	}

	fh := bufio.NewReader(inputfile)

	for {
		s, err := fh.ReadString('\n')
		if err == io.EOF { break; }
		if err != nil {
			fmt.Fprintf(os.Stderr,"Cannot read from data file\n")
			os.Exit(2)
		}
		then = append(then,s)
	}
}

func print_phrase() {
	fmt.Printf("Then %s",then[random() % len(then)])
}

func main() {
//
	init_delfy()
	fmt.Printf("--\n")

	for {
		fmt.Printf("If ")
		s, err := stdin.ReadString('\n')
		if len(s) < 1 || (len(s) == 1 && s[0] == '\n') || err == io.EOF { break; }
		if err != nil {
			fmt.Fprintf(os.Stderr,"Cannot read from data file\n")
			os.Exit(2)
		}
		print_phrase()
		fmt.Printf("--\n")
	}
}
