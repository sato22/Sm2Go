// This is a test file for testing state transitions

package ledbutton

import (
	"Sm2Go/sample/demo/sm2go"
	"log"
	"testing"
	"time"
)

// output device
type Led struct {
	name    string
	current string // "High" or "Low"
}

type Button struct {
	name    string
	release bool // true → release状態，false → push状態
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

// Button
func (b *Button) Push() {
	b.release = false
	log.Println(b.name, "Push")
}

func (b *Button) Release() {
	b.release = true
	log.Println(b.name, "Release")
}

func (b Button) Get() bool {
	return b.release
}

// define struct
var led = &Led{"led", "Low"}
var leftButton = &Button{"leftButton", true}
var rightButton = &Button{"rightButton", true}

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
	ConfigureDevice(led)
	ConfigureLeftButton(leftButton)
	ConfigureRightButton(rightButton)

	ConfigureLog(logTest)

	env := sm2go.NewTestEnv() // TestEnv構造体

	// goroutine (base.go Task())
	env.Add(sm2go.Continue, func() {
		for {
			Task()
			env.Sleep(10 * time.Millisecond)
		}
	},
	)

	// goroutine (user operation)
	env.Add(sm2go.Done, func() {
		// leftButtonPush
		log.Println("----------------leftButtonPush-------------")
		env.Sleep(1 * time.Second)
		leftButton.Push()
		env.Sleep(500 * time.Millisecond)
		leftButton.Release()

		// rightButtonPush
		log.Println("----------------rightButtonPush-------------")
		env.Sleep(1 * time.Second)
		rightButton.Push()
		env.Sleep(500 * time.Millisecond)
		rightButton.Release()

		// leftButtonPush double
		log.Println("----------------leftButtonPush double-------------")
		env.Sleep(1 * time.Second)
		leftButton.Push()
		env.Sleep(500 * time.Millisecond)
		leftButton.Release()
		env.Sleep(1 * time.Second)
		leftButton.Push()
		env.Sleep(500 * time.Millisecond)
		leftButton.Release()
		env.Sleep(1 * time.Second)
	},
	)

	env.Set(1)
	env.Go()
}
