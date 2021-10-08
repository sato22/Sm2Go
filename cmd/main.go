package main

import (
	"flag"
	"fmt"
	"github.com/rellab/sm2go/sm2go"
	"io/ioutil"
	"os"
)

// ---------------------------------- main ---------------------------------------

func main() {
	infile := flag.String("i", "", "drow.io file for State Machine")
	outfile := flag.String("o", "", "Output file (No edit)")
	outfile_edit := flag.String("oe", "", "Output file (Edit)")
	flag.Parse()

	var result map[string]*sm2go.StateMachine

	if *infile != "" {
		if xml, err := ioutil.ReadFile(*infile); err == nil {
			result = sm2go.Parse(xml)
		} else {
			panic(err)
		}
	}

	if (*outfile != "") && (*outfile_edit != "") {
		file, err := os.Create(*outfile)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		file_edit, err := os.Create(*outfile_edit)
		if err != nil {
			panic(err)
		}
		defer file_edit.Close()

		// write
		for _, v := range result {
			// ------------------------------　output.go　-----------------------------------
			sm2go.Write_package(file)
			sm2go.Write_enum(file, v.States)
			sm2go.Write_event(file, v.Transitions)
			sm2go.Write_init(file, v.Initial)
			// ------------------------------　output_edit.go　-----------------------------------
			sm2go.Write_package_edit(file_edit)
			sm2go.Write_func(file_edit, v.States, v.Events)
			sm2go.Write_main(file_edit)
		}
	}

}
