// This is a test file for testing state transitions

package pkg

import (
	"log"
	"sync"
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
}

func TestDevice(t *testing.T) {
	passcode = "0000"
	keypad.Configure()

	ConfigureDevice(led)
	ConfigureKeypad(keypad)
	ConfigureLog(logTest)

	var wg sync.WaitGroup
	wg.Add(1)

	// goroutine (base.go Task())
	go func() {
		for {
			time.Sleep(50 * time.Millisecond)
			Task()
		}
	}()

	// goroutine (user operation)
	go func() {
		// enter correct
		keypad.PushKey(0, 0)
		time.Sleep(500 * time.Millisecond)
		keypad.PushKey(0, 0)
		time.Sleep(500 * time.Millisecond)
		keypad.PushKey(0, 0)
		time.Sleep(500 * time.Millisecond)
		keypad.PushKey(0, 0)
		time.Sleep(4 * time.Second)

		// enter fail
		keypad.PushKey(0, 0)
		time.Sleep(500 * time.Millisecond)
		keypad.PushKey(0, 1)
		time.Sleep(500 * time.Millisecond)
		keypad.PushKey(0, 2)
		time.Sleep(500 * time.Millisecond)
		keypad.PushKey(0, 3)
		time.Sleep(4 * time.Second)

		// enter correct
		keypad.PushKey(0, 0)
		time.Sleep(500 * time.Millisecond)
		keypad.PushKey(0, 0)
		time.Sleep(500 * time.Millisecond)
		keypad.PushKey(0, 0)
		time.Sleep(500 * time.Millisecond)
		keypad.PushKey(0, 0)
		time.Sleep(4 * time.Second)

		wg.Done()
	}()
	wg.Wait()
}
