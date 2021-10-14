package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sato22/Sm2Go/pkg"
)

func main() {
	infile := flag.String("i", "", "diagrams.net file for State Machine")
	outfile := flag.String("o", "", "Output file (No edit)")
	outfileEdit := flag.String("oe", "", "Output file (Edit)")
	outfileTest := flag.String("ot", "", "Output file (Test)")
	flag.Parse()

	// パッケージ名を指定
	var name string
	fmt.Println("--------Please enter the package name--------")
	fmt.Scan(&name)

	var oline []string
	var oeline []string
	var otline []string

	// 入力ファイル（diagrams.netにて作成）の読み込み
	if *infile != "" {
		if xml, err := ioutil.ReadFile(*infile); err == nil {
			// result := sm2go.Parse(xml)
			oline, oeline, otline = sm2go.WriteAll(xml, name)
		} else {
			panic(err)
		}
	}

	// 出力ファイルの指定（編集しないファイル）
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
				panic(err)
			}
		}
	}

	// 出力ファイルの指定（要編集ファイル）
	if *outfileEdit != "" {
		fileEdit, err := os.Create(*outfileEdit)
		if err != nil {
			panic(err)
		}
		defer fileEdit.Close()

		for _, oe := range oeline{
			b := []byte(oe)
			_, err := fileEdit.Write(b)
			if err != nil {
				panic(err)
			}
		}
	}

	// 出力ファイルの指定（テスト用ファイル）
	if *outfileTest != "" {
		fileTest, err := os.Create(*outfileTest)
		if err != nil {
			panic(err)
		}
		defer fileTest.Close()

		for _, ot := range otline{
			b := []byte(ot)
			_, err := fileTest.Write(b)
			if err != nil {
				panic(err)
			}
		}
	}
}
