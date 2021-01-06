package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rafecoolz/sdhash"
)

var (
	// VERSION is set by the makefile
	VERSION = "v0.0.0"
	// BUILDDATE is set by the makefile
	BUILDDATE = ""
)

var file string
var compare string
var raw bool
var blockSize int
var version bool

func init() {
	flag.StringVar(&file, "f", "", "path to the `file` to be hashed")
	flag.StringVar(&compare, "c", "", "specifies a `filename` or `digest` whose sdhash value will be compared to a filename specified (-f)")
	flag.BoolVar(&raw, "r", false, "set to get only the hash")
	flag.IntVar(&blockSize, "b", 0, "hashes input files in nKB blocks")
	flag.BoolVar(&version, "version", false, "print version")
	flag.Parse()
}

// Main contains the main code
func Main() {
	if version {
		fmt.Printf("%s %s\n", VERSION, BUILDDATE)
		return
	}
	if file == "" {
		fmt.Fprintf(os.Stderr, "Usage of %s [-f <file>]\n\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Println()
		return
	}
	hash, err := sdhash.Hash(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	if compare != "" {
		hashCompare, err := sdhash.Hash(compare)
		if err != nil {
			fmt.Println(err)
			return
		}
		distance, err := sdhash.DiffFilenames(hashCompare)

		fmt.Printf("%d  %s  %s - %s  %s\n", distance, hash, file, hashCompare, compare)
	} else {
		if raw {
			fmt.Println(hash)
		} else {
			fmt.Printf("%s  %s\n", hash, file)
		}
	}
}

func main() {
	Main()
}
