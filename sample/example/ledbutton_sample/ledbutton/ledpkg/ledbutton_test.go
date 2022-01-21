// This is a test file for testing state transitions

package ledpkg

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
var led = &Led{"led", "High"}
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

func TestStateTrans(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			time.Sleep(1 * time.Millisecond)
			Task()
		}
	}()
	wg.Wait()
}

func TestDevice(t *testing.T) {
	ConfigureDevice(led)
	ConfigureLeftButton(leftButton)
	ConfigureRightButton(rightButton)

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
		// leftButtonPush
		log.Println("----------------leftButtonPush-------------")
		time.Sleep(1 * time.Second)
		leftButton.Push()
		time.Sleep(500 * time.Millisecond)
		leftButton.Release()

		// rightButtonPush
		log.Println("----------------rightButtonPush-------------")
		time.Sleep(1 * time.Second)
		rightButton.Push()
		time.Sleep(500 * time.Millisecond)
		rightButton.Release()

		// leftButtonPush double
		log.Println("----------------leftButtonPush double-------------")
		time.Sleep(1 * time.Second)
		leftButton.Push()
		time.Sleep(500 * time.Millisecond)
		leftButton.Release()
		time.Sleep(1 * time.Second)
		leftButton.Push()
		time.Sleep(500 * time.Millisecond)
		leftButton.Release()
		time.Sleep(1 * time.Second)
		wg.Done()
	}()
	wg.Wait()
}
