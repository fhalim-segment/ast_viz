package main

import (
	"io"
	"log"
	"os"

	"gitlab.com/fawad/ast-viz/astviz"
)

func main() {
	infile := os.Stdin
	outfile := os.Stdout
	defer infile.Close()
	defer outfile.Close()
	bytes, err := io.ReadAll(infile)

	if err != nil {
		log.Fatal("Unable to read input", err)
	}
	v, err := astviz.Load(bytes)
	if err != nil {
		log.Fatal("Unable to process input", err)
	}
	v.Dump(outfile, 0)
}
