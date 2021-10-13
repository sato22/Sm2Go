package main

import (
	"flag"
	"io/ioutil"
	"os"
	"github.com/sato22/Sm2Go/pkg"
)

func main() {
	infile := flag.String("i", "", "drow.io file for State Machine")
	outfile := flag.String("o", "", "Output file (No edit)")
	outfileEdit := flag.String("oe", "", "Output file (Edit)")
	flag.Parse()

	if *infile != "" {
		if xml, err := ioutil.ReadFile(*infile); err == nil {
			result := sm2go.Parse(xml)
		} else {
			panic(err)
		}
	}

	if *outfile != "" {
		file, err := os.Create(*outfile)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		
		for _, o := range oline {
			b := []byte(o)
			_, err := file.Write(b)
			if err != nil {
				return err
			}
		}
	}

	if *outfileEdit != "" {
		fileEdit, err := os.Create(*outfileEdit)
		if err != nil {
			panic(err)
		}
		defer fileEdit.Close()
		

		for _, o := range oeline{
			b := []byte(oe)
			_, err := file.Write(b)
			if err != nil {
				return err
			}
		}
	}
}
