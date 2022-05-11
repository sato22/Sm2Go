// This is a test file for testing state transitions

package pkg

import (
	"Sm2Go/sample/arduino_lesson/lesson11_driver/keypad4x4"
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

// pin
type Pin struct {
	name string
	high bool
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

// Pin
func (p Pin) Configure(interface{}) {}

func (p Pin) Get() bool {
	return p.high
}

func (p *Pin) High() {
	p.high = true
}

func (p *Pin) Low() {
	p.high = false
}

func (p *Pin) Set(h bool) {
	p.high = h
}

var button [16]string = [16]string{"Off", "Off", "Off", "Off", "Off", "Off", "Off", "Off", "Off", "Off", "Off", "Off", "Off", "Off", "Off", "Off"}

// push button
func pushButtonOne() {
	button[0] = "On"
}

func pushButtonTwo() {
	button[1] = "On"
}

func pushButtonThree() {
	button[2] = "On"
}

func pushButtonA() {
	button[3] = "On"
}

func pushButtonFour() {
	button[4] = "On"
}

func pushButtonFive() {
	button[5] = "On"
}

func pushButtonSix() {
	button[6] = "On"
}

func pushButtonB() {
	button[7] = "On"
}

func pushButtonSeven() {
	button[8] = "On"
}

func pushButtonEight() {
	button[9] = "On"
}

func pushButtonNine() {
	button[10] = "On"
}

func pushButtonC() {
	button[11] = "On"
}

func pushButtonAsterisk() {
	button[12] = "On"
}

func pushButtonZero() {
	button[13] = "On"
}

func pushButtonSharp() {
	button[14] = "On"
}

func pushButtonD() {
	button[15] = "On"
}

// releaseButton
func releaseButtonOne() {
	button[0] = "Off"
}

func releaseButtonTwo() {
	button[1] = "Off"
}

func releaseButtonThree() {
	button[2] = "Off"
}

func releaseButtonA() {
	button[3] = "Off"
}

func releaseButtonFour() {
	button[4] = "Off"
}

func releaseButtonFive() {
	button[5] = "Off"
}

func releaseButtonSix() {
	button[6] = "Off"
}

func releaseButtonB() {
	button[7] = "Off"
}

func releaseButtonSeven() {
	button[8] = "Off"
}

func releaseButtonEight() {
	button[9] = "Off"
}

func releaseButtonNine() {
	button[10] = "Off"
}

func releaseButtonC() {
	button[11] = "Off"
}

func releaseButtonAsterisk() {
	button[12] = "Off"
}

func releaseButtonZero() {
	button[13] = "Off"
}

func releaseButtonSharp() {
	button[14] = "Off"
}

func releaseButtonD() {
	button[15] = "Off"
}

func circuitSurv() {
	if button[0] == "On" {
		c4.Set(r4.Get())
		time.Sleep(1 * time.Microsecond)
	} else {
		c4.High()
	}

	if button[1] == "On" {
		c3.Set(r4.Get())
		time.Sleep(1 * time.Microsecond)
	} else {
		c3.High()
	}

	if button[2] == "On" {
		c2.Set(r4.Get())
		time.Sleep(1 * time.Microsecond)
	} else {
		c2.High()
	}

	if button[3] == "On" {
		// fmt.Println("button[3] On")
		c1.Set(r4.Get())
		time.Sleep(1 * time.Microsecond)
	} else {
		c1.High()
	}

	if button[4] == "On" {
		c4.Set(r3.Get())
		time.Sleep(1 * time.Microsecond)
	} else {
		c4.High()
	}

	if button[5] == "On" {
		// fmt.Println("button[5] On")
		c3.Set(r3.Get())
		time.Sleep(1 * time.Microsecond)
	} else {
		// fmt.Println("button[5] Off")
		c3.High()
	}

	if button[6] == "On" {
		c2.Set(r3.Get())
		time.Sleep(1 * time.Microsecond)
	} else {
		c2.High()
	}

	if button[7] == "On" {
		c1.Set(r3.Get())
		time.Sleep(1 * time.Microsecond)
	} else {
		c1.High()
	}

	if button[8] == "On" {
		c4.Set(r2.Get())
		time.Sleep(1 * time.Microsecond)
	} else {
		c4.High()
	}

	if button[9] == "On" {
		c3.Set(r2.Get())
		time.Sleep(1 * time.Microsecond)
	} else {
		c3.High()
	}

	if button[10] == "On" {
		c2.Set(r2.Get())
		time.Sleep(1 * time.Microsecond)
	} else {
		c2.High()
	}

	if button[11] == "On" {
		c1.Set(r2.Get())
		time.Sleep(1 * time.Microsecond)
	} else {
		c1.High()
	}

	if button[12] == "On" {
		c4.Set(r1.Get())
		time.Sleep(1 * time.Microsecond)
	} else {
		c4.High()
	}

	if button[13] == "On" {
		c3.Set(r1.Get())
		time.Sleep(1 * time.Microsecond)
	} else {
		c3.High()
	}

	if button[14] == "On" {
		c2.Set(r1.Get())
		time.Sleep(1 * time.Microsecond)
	} else {
		c2.High()
	}

	if button[15] == "On" {
		c1.Set(r1.Get())
		time.Sleep(1 * time.Microsecond)
	} else {
		c1.High()
	}
}

// define struct
var led = &Led{"led", "Low"}
var r4, r3, r2, r1, c4, c3, c2, c1 = &Pin{"r4", true}, &Pin{"r3", true}, &Pin{"r2", true}, &Pin{"r1", true}, &Pin{"c4", true}, &Pin{"c3", true}, &Pin{"c2", true}, &Pin{"c1", true}

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
		current = OffEmpty
	}
	eod = Entry
}

func TestDevice01(t *testing.T) {
	// Success

	passcode = "3210"
	keypad := keypad4x4.NewDevice(r4, r3, r2, r1, c4, c3, c2, c1)
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
		pushButtonThree()
		time.Sleep(500 * time.Millisecond)
		releaseButtonThree()

		time.Sleep(500 * time.Millisecond)

		pushButtonTwo()
		time.Sleep(500 * time.Millisecond)
		releaseButtonTwo()

		time.Sleep(500 * time.Millisecond)

		pushButtonOne()
		time.Sleep(500 * time.Millisecond)
		releaseButtonOne()

		time.Sleep(500 * time.Millisecond)

		pushButtonZero()
		time.Sleep(500 * time.Millisecond)
		releaseButtonZero()

		time.Sleep(5 * time.Second)
		wg.Done()
	}()

	// ボタンが押されているかを見るメソッド
	go func() {
		// 回路的な電圧の処理
		for {
			circuitSurv()
		}
	}()

	wg.Wait()
}

func TestDevice02(t *testing.T) {
	// Fail

	passcode = "3210"
	keypad := keypad4x4.NewDevice(r4, r3, r2, r1, c4, c3, c2, c1)
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
		// enter fail(3200)
		pushButtonThree()
		time.Sleep(500 * time.Millisecond)
		releaseButtonThree()

		time.Sleep(500 * time.Millisecond)

		pushButtonTwo()
		time.Sleep(500 * time.Millisecond)
		releaseButtonTwo()

		time.Sleep(500 * time.Millisecond)

		pushButtonZero()
		time.Sleep(500 * time.Millisecond)
		releaseButtonZero()

		time.Sleep(500 * time.Millisecond)

		pushButtonZero()
		time.Sleep(500 * time.Millisecond)
		releaseButtonZero()

		time.Sleep(5 * time.Second)
		wg.Done()
	}()

	// ボタンが押されているかを見るメソッド
	go func() {
		// 回路的な電圧の処理
		for {
			circuitSurv()
		}
	}()

	wg.Wait()
}
