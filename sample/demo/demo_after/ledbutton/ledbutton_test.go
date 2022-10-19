// This is a test file for testing state transitions

package ledbutton

import (
	"Sm2Go/sample/demo_tyukan/sm2go"
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
var butt = &Button{"button", true}

type DebugStruct struct{}

var logTest = DebugStruct{}

func (l DebugStruct) Println(debstr string) {
	log.Println(debstr)
}

func TestDevice(t *testing.T) {
	ConfigureDevice(led)
	ConfigureButton(butt)

	ConfigureLog(logTest)

	env := sm2go.NewTestEnv() // TestEnv構造体

	// goroutine(base.go Task())
	env.Add(sm2go.Continue, func() {
		for {
			time.Sleep(10 * time.Millisecond)
			Task()
		}
	},
	)

	// goroutine(user operation)
	env.Add(sm2go.Done, func() {
		// buttonPush
		env.Sleep(500 * time.Millisecond)
		butt.Push()
		env.Sleep(500 * time.Millisecond)
		butt.Release()

		env.Sleep(500 * time.Millisecond)
		butt.Push()
		env.Sleep(500 * time.Millisecond)
		butt.Release()
		env.Sleep(500 * time.Millisecond)
	},
	)

	env.Set(1)
	env.Go()
}
