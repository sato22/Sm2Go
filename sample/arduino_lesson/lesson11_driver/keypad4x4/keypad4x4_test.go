package keypad4x4

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// pin
type Pin struct {
	name string
	high bool // true：high, false：low
}

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

// type Pin struct {
// 	label string
// 	in    chan int
// 	out   chan bool
// }

// func NewPin(s string, v bool) *Pin {
// 	p := &Pin{
// 		label: s,
// 		in:    make(chan int),
// 		out:   make(chan bool),
// 	}

// 	go func(v bool) {
// 		x := v
// 		for {
// 			switch ch := <-p.in; ch {
// 			case 0:
// 				p.out <- x
// 			case 1:
// 				// x = ch
// 				p.out <- x
// 			case 2:
// 				// x = ch
// 				p.out <- x
// 			}
// 		}
// 	}(v)

// 	return p
// }

// func (p *Pin) Get() bool {
// 	p.in <- 0
// 	x := <-p.out
// 	return x
// }

// func (p *Pin) Set(v bool) {
// 	p.out <- v
// 	_ = <-p.out
// }

// func (p *Pin) High() {
// 	p.in <- 1
// 	p.out <- true
// }

// func (p *Pin) Low() {
// 	p.in <- 2
// 	p.out <- false
// }

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

// 回路的な電圧を見る関数
func circuitSurv() {
	if button[0] == "On" {
		c4.Set(r4.Get())
		time.Sleep(1 * time.Nanosecond)
	} else {
		c4.High()
	}

	if button[1] == "On" {
		c3.Set(r4.Get())
		time.Sleep(1 * time.Nanosecond)
	} else {
		c3.High()
	}

	if button[2] == "On" {
		c2.Set(r4.Get())
		time.Sleep(1 * time.Nanosecond)
	} else {
		c2.High()
	}

	if button[3] == "On" {
		// fmt.Println("button[3] On")
		c1.Set(r4.Get())
		time.Sleep(1 * time.Nanosecond)
	} else {
		c1.High()
	}

	if button[4] == "On" {
		c4.Set(r3.Get())
		time.Sleep(1 * time.Nanosecond)
	} else {
		c4.High()
	}

	if button[5] == "On" {
		// fmt.Println("button[5] On")
		c3.Set(r3.Get())
		time.Sleep(1 * time.Nanosecond)
	} else {
		// fmt.Println("button[5] Off")
		c3.High()
	}

	if button[6] == "On" {
		c2.Set(r3.Get())
		time.Sleep(1 * time.Nanosecond)
	} else {
		c2.High()
	}

	if button[7] == "On" {
		c1.Set(r3.Get())
		time.Sleep(1 * time.Nanosecond)
	} else {
		c1.High()
	}

	if button[8] == "On" {
		c4.Set(r2.Get())
		time.Sleep(1 * time.Nanosecond)
	} else {
		c4.High()
	}

	if button[9] == "On" {
		c3.Set(r2.Get())
		time.Sleep(1 * time.Nanosecond)
	} else {
		c3.High()
	}

	if button[10] == "On" {
		c2.Set(r2.Get())
		time.Sleep(1 * time.Nanosecond)
	} else {
		c2.High()
	}

	if button[11] == "On" {
		c1.Set(r2.Get())
		time.Sleep(1 * time.Nanosecond)
	} else {
		c1.High()
	}

	if button[12] == "On" {
		c4.Set(r1.Get())
		time.Sleep(1 * time.Nanosecond)
	} else {
		c4.High()
	}

	if button[13] == "On" {
		c3.Set(r1.Get())
		time.Sleep(1 * time.Nanosecond)
	} else {
		c3.High()
	}

	if button[14] == "On" {
		c2.Set(r1.Get())
		time.Sleep(1 * time.Nanosecond)
	} else {
		c2.High()
	}

	if button[15] == "On" {
		c1.Set(r1.Get())
		time.Sleep(1 * time.Nanosecond)
	} else {
		c1.High()
	}
}

var r4, r3, r2, r1, c4, c3, c2, c1 = &Pin{"r4", true}, &Pin{"r3", true}, &Pin{"r2", true}, &Pin{"r1", true}, &Pin{"c4", true}, &Pin{"c3", true}, &Pin{"c2", true}, &Pin{"c1", true}

func TestPin01(t *testing.T) {
	r4.High()
	if r4.Get() != true {
		t.Errorf("expected_r4_high true, result_r4 %v", r4.high)
	}
}

func TestGetIndices01(t *testing.T) {
	// no pin High()
	keypad := NewDevice(r4, r3, r2, r1, c4, c3, c2, c1)
	keypad.Configure()

	var row, column int

	go func() {
		for {
			row, column = keypad.GetIndices()
			if row != -1 && column != -1 {
				fmt.Println("row, column =", row, column)
			}
		}
	}()

	time.Sleep(500 * time.Millisecond)
	exRow, exColumn := -1, -1
	if row != exRow || column != exColumn {
		t.Errorf("expected_row %v , expected_column %v, result_row %v, result_column %v", exRow, exColumn, row, column)
	}
}

func TestGetIndices02(t *testing.T) {
	// pushButtonFive()
	keypad := NewDevice(r4, r3, r2, r1, c4, c3, c2, c1)
	keypad.Configure()
	var row, column int
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		// ユーザの操作
		pushButtonFive() // button[5] = "On"

		row, column = keypad.GetIndices() // ドライバー
		fmt.Println("row, column =", row, column)

		releaseButtonFive() // button[5] = Off
		fmt.Println("end")
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

func TestGetIndices03(t *testing.T) {
	// push "A200"
	keypad := NewDevice(r4, r3, r2, r1, c4, c3, c2, c1)
	keypad.Configure()
	var row, column int

	var wg sync.WaitGroup
	wg.Add(1)

	// ボタンが押されたかと、押された位置を見続ける
	go func() {
		for {
			row, column = keypad.GetIndices() // 時間を入れるのであれば何秒に一回動いているかを設定？
			if row != -1 && column != -1 {
				fmt.Println("row, column =", row, column)
			}
		}
	}()

	/*
		・ ユーザが自作したドライバーがテストできる環境を
		・ ユーザ視点でやりたいことは？それを踏まえてどのような仕様にするか？
			・ ユーザとしてやりたいこと：ドライバーを自作した際にドライバーのテストをしたい
			・ 現在、Tinygoのドライバーがすべて用意されていない状況
				・ 配付されているドライバーは信頼できる
				・ 配布されていないドライバーは、自作する必要がある
					・ これのテストがしたい
		・ 連続でのボタン入力（5押した後に2押すとか）のテストはいらない…？？
		・ 関数がきちんと値を取ってこれてるかを見たいだけ
			・テストで値を取ってくるには、time.Sleepを用いる必要がある
				・time.Sleep(1 * time.Microsecond) で値を取れるけど、実機に実装した際に致命的
					・ドライバーを書き換えることになるので
					・あと、ユーザがドライバーを自作した際にも、time.Sleepをいれなければテストできない
				・ドライバーを書き換えずにテストしたい…
				・やはりチャネルを使ったやり取りを実装しなければ難しい？
	*/

	// ユーザの操作
	go func() {
		// push A
		pushButtonA()
		time.Sleep(500 * time.Millisecond) // 0.5秒間押す
		releaseButtonA()

		// 意味がない数字を入れない
		// どれくらいの時間ボタン押したら
		// テストの中身を知らない人には理解できない
		time.Sleep(500 * time.Millisecond)

		// push 2
		pushButtonTwo()
		time.Sleep(500 * time.Millisecond) // 0.5秒間押す
		releaseButtonTwo()

		time.Sleep(500 * time.Millisecond)

		// push 0
		pushButtonZero()
		time.Sleep(500 * time.Millisecond) // 0.5秒間押す
		releaseButtonZero()

		time.Sleep(500 * time.Millisecond)

		// push 0
		pushButtonZero()
		time.Sleep(500 * time.Millisecond) // 0.5秒間押す
		releaseButtonZero()

		time.Sleep(500 * time.Millisecond)

		fmt.Println("end")
		wg.Done()
	}()

	// ボタンが押されているかを見る
	go func() {
		// 回路的な電圧の処理
		for {
			circuitSurv()
		}
	}()

	wg.Wait()
}

func TestGetKey01(t *testing.T) {
	// no button pushed
	keypad := NewDevice(r4, r3, r2, r1, c4, c3, c2, c1)
	keypad.Configure()
	var value string

	go func() {
		for {
			value = keypad.GetKey()
			if value != "NoKeyPressed" {
				fmt.Println("value =", value)
				break
			}
		}
	}()

	time.Sleep(500 * time.Millisecond)
	exValue := "NoKeyPressed"
	if value != string(exValue) {
		t.Errorf("expected_value %v , result_value %v", exValue, value)
	}
}

func TestGetKey02(t *testing.T) {
	// pushButtonFive
	keypad := NewDevice(r4, r3, r2, r1, c4, c3, c2, c1)
	keypad.Configure()
	var value string

	var wg sync.WaitGroup
	wg.Add(1)

	// ユーザの操作
	go func() {
		pushButtonFive() // button[5] = "On"

		value = keypad.GetKey() // ドライバー
		fmt.Println("value =", value)

		releaseButtonFive() // button[5] = Off
		fmt.Println("end")
		wg.Done()
	}()

	// ボタンが押されているかを見る
	go func() {
		// 回路的な電圧の処理
		for {
			circuitSurv()
		}
	}()

	wg.Wait()
}

func TestGetKey03(t *testing.T) {
	// pushButton A200
	keypad := NewDevice(r4, r3, r2, r1, c4, c3, c2, c1)
	keypad.Configure()
	var value string

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for {
			value = keypad.GetKey()
			if value != "NoKeyPressed" {
				fmt.Println("value =", value)
			}
		}
	}()

	// ユーザの操作
	go func() {
		// push A
		pushButtonA()
		time.Sleep(500 * time.Millisecond) // 0.5秒間押す
		releaseButtonA()

		time.Sleep(500 * time.Millisecond)

		// push 2
		pushButtonTwo()
		time.Sleep(500 * time.Millisecond) // 0.5秒間押す
		releaseButtonTwo()

		time.Sleep(500 * time.Millisecond)

		// push 0
		pushButtonZero()
		time.Sleep(500 * time.Millisecond) // 0.5秒間押す
		releaseButtonZero()

		time.Sleep(500 * time.Millisecond)

		// push 0
		pushButtonZero()
		time.Sleep(500 * time.Millisecond) // 0.5秒間押す
		releaseButtonZero()

		time.Sleep(500 * time.Millisecond) // time.Sleepで充分待つ

		fmt.Println("end")
		wg.Done()
	}()

	// ボタンが押されているかを見る
	go func() {
		// 回路的な電圧の処理
		for {
			circuitSurv()
		}
	}()

	wg.Wait()
}
