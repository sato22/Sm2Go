// Please do not edit this file.

package ko

// package name to import

const (
	debug = true
)

type State int

const (
	OffEntered State = iota
	OffEmpty
)

type Eod int

const (
	Entry Eod = iota
	Do
	Exit
)

var eod Eod
var currentState State

/*
【エラー】
package Sm2Go/sample/arduino_lesson/lesson11_composite/keypad_eee/ko
	imports Sm2Go/sample/arduino_lesson/lesson11_composite/keypad_eee
	imports Sm2Go/sample/arduino_lesson/lesson11_composite/keypad_eee/ko: import cycle not allowed

import cycle（循環参照，AでBをインポートし，BでAをインポートすると発生する．）により，実行不可能．
解決方法はinterface，変数は共有する必要があるため，koパッケージのInit()とStep()を
keypad_eeeパッケージでインポートしなくてもいいように，interfaceを活用しなければならない？
*/

type Child struct{}

func NewChild() *Child {
	return new(Child)
}

func (c *Child) Step() {
	switch currentState {
	case OffEmpty:
		if eod == Entry {
			offemptyEntry()
			eod = Do
		}
		if eod == Do {
			offemptyDo()
			if keyEnteredCond() {
				currentState = OffEntered
				if debug {
					logger.Println("State is changed: OffEmpty to OffEntered")
				}
				eod = Exit
			}
		}
		if eod == Exit {
			offemptyExit()
			eod = Entry
		}
	case OffEntered:
		if eod == Entry {
			offenteredEntry()
			eod = Do
		}
		if eod == Do {
			offenteredDo()
			if keyFailedCond() {
				currentState = OffEmpty
				if debug {
					logger.Println("State is changed: OffEntered to OffEmpty")
				}
				eod = Exit
			}
		}
		if eod == Exit {
			offenteredExit()
			eod = Entry
		}
	}
}

func (c *Child) Init() {
	currentState = OffEmpty
	eod = Entry
}
