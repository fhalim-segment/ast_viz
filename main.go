package main

import (
	"flag"
	"io"
	"log"
	"os"

	"gitlab.com/fawad/ast-viz/astviz"
	"gitlab.com/fawad/ast-viz/journeyviz"
)

func main() {
	procesAst := flag.Bool("a", false, "Whether the input is an AST. Defaults to processing it as a Journey")
	flag.Parse()
	infile := os.Stdin
	outfile := os.Stdout
	defer infile.Close()
	defer outfile.Close()
	bytes, err := io.ReadAll(infile)

	if err != nil {
		log.Fatal("Unable to read input", err)
	}
	if *procesAst {
		v, err := astviz.Load(bytes)
		if err != nil {
			log.Fatal("Unable to process AST input", err)
		}
		v.Dump(outfile, 0)
	} else {
		v, err := journeyviz.Load(bytes)
		if err != nil {
			log.Fatal("Unable to process Journey input", err)
		}
		v.Dump(outfile, 0)
	}

}
