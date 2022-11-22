package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	sm2go "github.com/sato22/Sm2Go/pkg"
)

func main() {
	infile := flag.String("i", "", "diagrams.net file for State Machine")
	flag.Parse()

	count := 0

	m := map[string]string{"1": "one", "2": "two", "3": "three"}

	var oline []string
	var oeline []string
	var otline []string
	var omline []string
	var osline []string

	// ディレクトリ生成
	if err := os.MkdirAll(fmt.Sprintf("%s", os.Args[3]), 0777); err != nil {
		panic(err)
	}

	if *infile != "" {
		if xml, err := ioutil.ReadFile(*infile); err == nil {
			result := sm2go.Parse(xml)

			for key, v := range result {
				base, err := os.Create(fmt.Sprintf("%s/%s_base.go", os.Args[3], os.Args[3]))
				if err != nil {
					fmt.Println(err)
					return
				}
				defer base.Close()

				impl, err := os.Create(fmt.Sprintf("%s/%s_impl.go", os.Args[3], os.Args[3]))
				if err != nil {
					fmt.Println(err)
					return
				}
				defer impl.Close()

				test, err := os.Create(fmt.Sprintf("%s/%s_test.go", os.Args[3], os.Args[3]))
				if err != nil {
					fmt.Println(err)
					return
				}
				defer test.Close()

				oline, oeline, otline, omline, osline = sm2go.WriteAll(xml, os.Args[3], key, v, m, count)

				// package_base.go
				for _, o := range oline {
					b := []byte(o)
					_, err := base.Write(b)
					if err != nil {
						panic(err)
					}
				}

				// package_impl.go
				for _, oe := range oeline {
					b := []byte(oe)
					_, err := impl.Write(b)
					if err != nil {
						panic(err)
					}
				}

				// package_test.go
				for _, ot := range otline {
					b := []byte(ot)
					_, err := test.Write(b)
					if err != nil {
						panic(err)
					}
				}

				count++
			}

		} else {
			panic(err)
		}
	}

	// sm2goディレクトリ生成
	if err := os.Mkdir("sm2go", 0777); err != nil {
		fmt.Println(err)
	}

	// sm2go.go生成
	sm, err := os.Create("sm2go/sm2go.go")
	if err != nil {
		panic(err)
	}
	defer sm.Close()

	// main.go生成
	main, err := os.Create("main.go")
	if err != nil {
		panic(err)
	}
	defer main.Close()

	// sm2go.go
	for _, os := range osline {
		b := []byte(os)
		_, err := sm.Write(b)
		if err != nil {
			panic(err)
		}
	}

	// main.go
	for _, om := range omline {
		b := []byte(om)
		_, err := main.Write(b)
		if err != nil {
			panic(err)
		}
	}
}

// // 入力ファイル（diagrams.netにて作成）の読み込み
// if *infile != "" {
// 	if xml, err := ioutil.ReadFile(*infile); err == nil {
// 		oline, oeline, otline, omline, osline = sm2go.WriteAll(xml, os.Args[3])
// 	} else {
// 		panic(err)
// 	}
// }

// // ディレクトリ生成
// if err := os.MkdirAll(fmt.Sprintf("%s", os.Args[3]), 0777); err != nil {
// 	panic(err)
// }

// // package_base.go生成
// base, err := os.Create(fmt.Sprintf("%s/%s_base.go", os.Args[3], os.Args[3]))
// if err != nil {
// 	panic(err)
// }
// defer base.Close()

// // package_impl.go生成
// impl, err := os.Create(fmt.Sprintf("%s/%s_impl.go", os.Args[3], os.Args[3]))
// if err != nil {
// 	panic(err)
// }
// defer impl.Close()

// // package_test.go生成
// test, err := os.Create(fmt.Sprintf("%s/%s_test.go", os.Args[3], os.Args[3]))
// if err != nil {
// 	panic(err)
// }
// defer test.Close()

// // sm2go.go生成
// if err := os.Mkdir("sm2go", 0777); err != nil {
// 	fmt.Println(err)
// }

// sm, err := os.Create("sm2go/sm2go.go")
// if err != nil {
// 	panic(err)
// }
// defer sm.Close()

// // main.go生成
// main, err := os.Create("main.go")
// if err != nil {
// 	panic(err)
// }
// defer main.Close()

// // package_base.go
// for _, o := range oline {
// 	b := []byte(o)
// 	_, err := base.Write(b)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// // package_impl.go
// for _, oe := range oeline {
// 	b := []byte(oe)
// 	_, err := impl.Write(b)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// // package_test.go
// for _, ot := range otline {
// 	b := []byte(ot)
// 	_, err := test.Write(b)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// // sm2go.go
// for _, os := range osline {
// 	b := []byte(os)
// 	_, err := sm.Write(b)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// // main.go
// for _, om := range omline {
// 	b := []byte(om)
// 	_, err := main.Write(b)
// 	if err != nil {
// 		panic(err)
// 	}
// }
