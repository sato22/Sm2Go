// Please run "gofmt -w /path/to/output.go"

package main // write package name

import (
	// write about import
	"bufio"
	"fmt"
	// "log"
	"os"
	"time"
)

var count int
var input string

func Wait_Entry() {
	// Processing performed in the "Entry" state of the "Wait" state.
	// nothing to do
	count = 0
}

func Run_Entry() {
	// Processing performed in the "Entry" state of the "Run" state.
	// fmt.Println("Run_Entry() is executed.")

}

func Stop_Entry() {
	// Processing performed in the "Entry" state of the "Stop" state.
	// fmt.Println("Stop_Entry() is executed.")
	fmt.Println(count)
}

func Wait_Do() {
	// Processing performed in the "Do" state of the "Wait" state.
	// fmt.Println("Wait_Do() is executed.")
}

func Run_Do() {
	// Processing performed in the "Do" state of the "Run" state.
	// fmt.Println("Run_Do() is executed.")
	count++
}

func Stop_Do() {
	// Processing performed in the "Do" state of the "Stop" state.
	// fmt.Println("Stop_Do() is executed.")
}

func push_button_Cond() bool {
	// Please write the conditions under which a state transitions
	if input == "p" {
		input = "nothing"
		return true
	}
	return false
}

/*
func touch_display_Cond() bool {
	// Please write the conditions under which a state transitions
	if input == "t" {
		input = "nothing"
		return true
	}
	return false
}
*/

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
