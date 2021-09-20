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

var temp float64 = 20
var high_limit float64 = 22
var low_limit float64 = 18

func Heat_Entry() {
	// nothing to do
}

func Heat_Do() {
	temp = temp + 0.1
	fmt.Println(temp)
}

func Heat_Exit() {
	// nothing to do
}

func Unheat_Entry() {
	// nothing to do
}

func Unheat_Do() {
	temp = temp - 0.1
	fmt.Println(temp)
}

func Unheat_Exit() {
	// nothing to do
}

func Stop_Entry() {
	// nothing to do
}

func Stop_Do() {
	// nothing to do
}

func Stop_Exit() {
	// nothing to do
}

func higher_limit_Cond() bool {
	// Please write the conditions under which a state transitions
	if temp >= high_limit {
		return true
	}
	return false
}

func lower_limit_Cond() bool {
	// Please write the conditions under which a state transitions
	if temp < low_limit {
		return true
	}
	return false
}

func power_off_Cond() bool {
	// Please write the conditions under which a state transitions
	if input == "off" {
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
	go func() {
		time.Sleep(1 * time.Millisecond)
		state_transition()
	}()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		time.Sleep(1 * time.Millisecond) // Change according to specifications
		scan(scanner)
		if current_state == Stop {
			fmt.Println("quit")
			break
		}
	}
}
