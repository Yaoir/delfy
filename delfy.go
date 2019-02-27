package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/user"
	"time"
	)

var datafile []string = []string {
	"/usr/share/games/delfy/delfy.responses",
	"/usr/share/delfy/delfy.responses",
	"/usr/local/share/delfy/delfy.responses",
	"delfy.responses",
	}

// standard input
var stdin *bufio.Reader

// for random numbers
var randomgen *rand.Rand

// data file in a slice of strings
var then []string

const cmdname string = "delfy"

func err_exit(msg string) {
	fmt.Fprintf(os.Stderr,"%s: %s\n",cmdname,msg)
	os.Exit(1)
}

func init_delfy() {
//
	var err error
	var inputfile *os.File
	var usr *user.User

	// init random number generator
        randomgen = rand.New(rand.NewSource(time.Now().UnixNano()))

	// open standard input
	stdin = bufio.NewReader(os.Stdin)

	// find the data file (delfy.responses)

	for i := 0; i < len(datafile); i++ {
		inputfile, err = os.Open(datafile[i])
		if err == nil { break }	// found it
	}

	if err != nil {
		// none of the system directories or the current directory had it, so try $HOME/.delfy
		usr, err = user.Current()
		if err != nil { err_exit("cannot identify current user") }
		fname := usr.HomeDir + "/.delfy/delfy.responses"
		inputfile, err = os.Open(fname)
	}

	if err != nil { err_exit("cannot open delfy.responses for input") }

	fh := bufio.NewReader(inputfile)

	for {
		s, err := fh.ReadString('\n')
		if err == io.EOF { break; }
		if err != nil { err_exit("cannot read from data file") }
		then = append(then,s)
	}
	inputfile.Close()
}

func main() {
//
	init_delfy()
	fmt.Printf("--\n")

	for {
		fmt.Printf("If ")
		s, err := stdin.ReadString('\n')
		if len(s) < 1 || (len(s) == 1 && s[0] == '\n') || err == io.EOF { break; }
		if err != nil { err_exit("cannot read from data file") }
		fmt.Printf("Then %s",then[randomgen.Int() % len(then)])
		fmt.Printf("--\n")
	}
}
