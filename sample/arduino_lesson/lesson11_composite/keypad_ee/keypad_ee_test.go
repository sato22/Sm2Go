// This is a test file for testing state transitions

package keypad_easy

import (
	"Sm2Go/sample/arduino_lesson/lesson11_composite/sm2go"
	"log"
	"testing"
	"time"
)

// output device
type Led struct {
	name    string
	current string // "High" or "Low"
}

// keypad 4x4
type Keypad4x4 struct {
	inputEnabled bool
	lastColumn   int
	lastRow      int
	mapping      [4][4]uint8
}

// Led
func (l *Led) High() {
	l.current = "High"
	log.Println(l.name, l.current)
}

func (l *Led) Low() {
	l.current = "Low"
	log.Println(l.name, l.current)
}

//Keypad
func (keypad *Keypad4x4) Configure() {
	keypad.mapping = [4][4]uint8{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
		{12, 13, 14, 15},
	}

	keypad.inputEnabled = true
	keypad.lastColumn = -1
	keypad.lastRow = -1
}

func (keypad *Keypad4x4) GetKey() uint8 {
	if keypad.inputEnabled {
		keypad.inputEnabled = false
		return keypad.mapping[keypad.lastRow][keypad.lastColumn]
	} else {
		return NoKeyPressed
	}
}

func (keypad *Keypad4x4) PushKey(row, column int) {
	keypad.lastColumn = column
	keypad.lastRow = row
	keypad.inputEnabled = true
}

// define struct
var led = &Led{"led", "Low"}
var keypad = &Keypad4x4{}

//log.Println()
type DebugStruct struct{}

var logTest = DebugStruct{}

func (l DebugStruct) Println(debstr string) {
	log.Println(debstr)
}

// init
func init() {
	if led.current == "High" {
		current = On
	} else if led.current == "Low" {
		current = Off
	}
	eod = Entry
	eodSmall = Entry
}

func TestDevice(t *testing.T) {
	passcode = "0000"
	keypad.Configure()

	ConfigureDevice(led)
	ConfigureKeypad(keypad)
	ConfigureLog(logTest)

	env := sm2go.NewTestEnv() // TestEnv構造体

	// goroutine (base.go Task())
	env.Add(sm2go.Continue, func() {
		for {
			time.Sleep(50 * time.Millisecond)
			Task()
		}
	},
	)

	env.Add(sm2go.Continue, func() {
		for {
			time.Sleep(50 * time.Millisecond)
			Task2()
		}
	},
	)

	// goroutine (user operation)
	env.Add(sm2go.Done, func() {
		// enter correct
		keypad.PushKey(0, 0)
		env.Sleep(500 * time.Millisecond)
		keypad.PushKey(0, 0)
		env.Sleep(500 * time.Millisecond)
		keypad.PushKey(0, 0)
		env.Sleep(500 * time.Millisecond)
		keypad.PushKey(0, 0)
		env.Sleep(4 * time.Second)

		// enter fail
		keypad.PushKey(0, 0)
		env.Sleep(500 * time.Millisecond)
		keypad.PushKey(0, 1)
		env.Sleep(500 * time.Millisecond)
		keypad.PushKey(0, 2)
		env.Sleep(500 * time.Millisecond)
		keypad.PushKey(0, 3)
		env.Sleep(4 * time.Second)

		// enter correct
		keypad.PushKey(0, 0)
		env.Sleep(500 * time.Millisecond)
		keypad.PushKey(0, 0)
		env.Sleep(500 * time.Millisecond)
		keypad.PushKey(0, 0)
		env.Sleep(500 * time.Millisecond)
		keypad.PushKey(0, 0)
		env.Sleep(4 * time.Second)
	},
	)

	env.Set(1)
	env.Go()
}

// テスト関数は二ついらない．
// 領域ごとにテスト関数が生成されるようになっているため，修正．

// This is a test file for testing state transitions

// package keypad_easy

// import (
// "log"
// "testing"
// "time"
// "sync"
// )

// type DebugStruct struct{}
// var logTest = DebugStruct{}
// func (l DebugStruct) Println(debstr string) {
// log.Println(debstr)
// }
// func TestDevice(t *testing.T) {
// env := sm2go.NewTestEnv() // TestEnv構造体

// // goroutine(base.go Task())
// env.Add(sm2go.Continue, func() {
// for {
// time.Sleep(10 * time.Millisecond)
// Task()
// }
// },
// )
// // goroutine(user operation)
// env.Add(sm2go.Done, func() {
// },
// )

// env.Set(1)
// env.Go()
// }
