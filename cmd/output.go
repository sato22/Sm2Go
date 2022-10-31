package sm2go

//go:generate go run $GOFILE
//go:generate gofmt -w ../bitvec8.go
//go:generate gofmt -w ../bitvec16.go
//go:generate gofmt -w ../bitvec32.go
//go:generate gofmt -w ../bitvec64.go

import (
	"bytes"
	"io/ioutil"
	"log"
	"text/template"
)

type Bitvec struct {
	Bits string
	Mask string
}

func output(filename string, param Bitvec) {
	b, err := ioutil.ReadFile("bitvec_tmpl.go")
	if err != nil {
		log.Fatal(err)
	}
	tmpl := string(b)

	t, err := template.New("bitvec").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var buf bytes.Buffer
	if err = t.Execute(&buf, param); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(filename, buf.Bytes(), 0644); err != nil {
		log.Fatalf("writing output: %s", err)
	}
}

func main() {
	output("../bitvec8.go", Bitvec{
		Bits: "8",
		Mask: "0xff",
	})

	output("../bitvec16.go", Bitvec{
		Bits: "16",
		Mask: "0xffff",
	})

	output("../bitvec32.go", Bitvec{
		Bits: "32",
		Mask: "0xffff_ffff",
	})

	output("../bitvec64.go", Bitvec{
		Bits: "64",
		Mask: "0xffff_ffff_ffff_ffff",
	})
}
