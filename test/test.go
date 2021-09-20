package main

import (
	"fmt"
	"io"
	"os"
)

// ファイルの作成，書き込みのテスト

// 末尾に Hello World. を書き込むサンプル実装
func writeHello(w io.Writer) {
	fmt.Fprintln(w, "\nHello World.")
}

func main() {
	fp, err := os.Create("hello.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	ft, err := os.Create("hello_he.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ft.Close()

	fp.WriteString("hello")
	fmt.Fprintln(ft, "test")
	writeHello(ft)

	// 直近で書き込んだ内容をReadするにはSeekでファイルの先頭に戻る必要がある
	fp.Seek(0, 0)

	b := make([]byte, 256)
	for {
		_, err := fp.Read(b)
		if err != nil {
			// 			if err != io.EOF {
			// 				fmt.Println(err)
			// 				return
			// 			}
			break
		}
		fmt.Println(string(b))
	}
}
