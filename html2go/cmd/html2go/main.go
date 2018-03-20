package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/nobonobo/vecty-samples/html2go/trans"
)

func main() {
	var inputFname string
	var withBody bool
	flag.BoolVar(&withBody, "b", false, "with body element")
	flag.Parse()
	inputFname = flag.Arg(0)
	var input io.ReadCloser = os.Stdin
	if inputFname != "" {
		fp, err := os.Open(inputFname)
		if err != nil {
			log.Fatalln(err)
		}
		defer fp.Close()
		input = fp
	}
	if err := trans.Trans(os.Stdout, input, withBody); err != nil {
		log.Fatalln(err)
	}
}
