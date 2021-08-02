package main

import (
	"fmt"
)

// golangでenumを実装する場合
type State int

const (
	Wait State = iota
	Run
	Stop
)

type Eod int

const (
	Entry Eod = iota
	Do
	Exit
)

func task1(s State, eod Eod) State {
	// generate hogehoge
	switch s {
	case Wait:
		if eod == Entry {
			eod = Do
		}
		if eod == Do {
			// inputを確認
			if(// ここに，jsonファイルから取り出した条件文){
				s = Run
				fmt.Println("State is changed: Wait to Run")
				eod = Entry
			}
		}
	case Run:
		if eod == Entry {
			eod = Do
		}
		if eod == Do {
			if(// ここに，jsonファイルから取り出した条件文){
				s = Stop
				fmt.Println("State is changed: Run to Stop")
				eod = Entry
			}
		}
	case Stop:
		if eod == Entry {
			eod = Do
		}
		if eod == Do {
			if(// ここに，jsonファイルから取り出した条件文){
				s = Wait
				fmt.Println("State is changed: Stop to Wait")
				eod = Entry
			}
		}
	default:
		fmt.Println("Unknown")
	}
	return s
}

func main() {
	var s State = Wait // initial state
	var e Eod = Entry
	for {
		var input string
		fmt.Scan(&input)
		s = print_state(s, e)
		if input == "exit" {
			fmt.Println("Exit")
			break
		}
	}
}
