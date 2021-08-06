// Please edit this file
// Please run "gofmt -w /path/to/output.go"
// go run *.go

package main

import (
        "bufio"
        "fmt"
        "os"
        "time"
)

var count int

func Wait_Entry() {
	count = 0
}

func Wait_Do() {
        // nothing to do
}

func Wait_Exit() {
        // nothing to do
}

func Run_Entry() {
        // nothing to do
}

func Run_Do() {
	count++
}

func Run_Exit() {
        // nothing to do
}

func Stop_Entry() {
	fmt.Println(count)
}

func Stop_Do() {
        // nothing to do
}

func Stop_Exit() {
        // nothing to do
}

func push_button_Cond() bool {
        // Please write the conditions under which a state transitions
        if input == "p"{
        	input = "nothing"
        	return true
        }
        return false
}

func scan(scanner *bufio.Scanner) {
        for scanner.Scan() {
                input = scanner.Text()
        }
        if err := scanner.Err(); err != nil {
                fmt.Println(err)
        }
}

func main() {
        scanner := bufio.NewScanner(os.Stdin)
        for {
                time.Sleep(100 * time.Millisecond) // Change according to specifications
                go state_transition()
                go scan(scanner)
                if input == "q" {
                        fmt.Println("quit")
                        break
                }
        }
}