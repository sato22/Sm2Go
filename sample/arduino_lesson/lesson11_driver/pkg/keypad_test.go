// This is a test file for testing state transitions

package pkg

import (
	"log"
	"sync"
	"testing"
	"time"
)

/*
	// 一旦Pinのことは考えない(考えたい、考えられたら便利だけど、仮想のPinを設定するとうまくいかない)
	// テストとして丸ごと配布しなきゃいけないなら、もうこちら側でGetKey()などのドライバー内の関数も設定しちゃおう
	// ドライバーは一切触らず、テストで扱っている（パッケージを読み込んで使っている）ドライバー内の関数はこちらで作成する
	// GetKey()とGetIndices()をテスト内で自作して、LEDのON/OFFが移り変わるさまを確認したい
	// テストしたいのは、あくまでステートの遷移！
		OffEmpty →　OffEntered →　On　と移り変わる様が見たい
		ユーザが自作したドライバーが動くかどうかのテストはまた別の領域、別の話
		テストとしては、パスコードが入力されて、合ってたらLEDがついて、間違ってたらつかない（入力に戻る）っていうのが確認したい
	// time.Sleepを入れない構成にする
		これだと、チャネルを用いることで同期処理にする必要がある
		【同期処理】
			・例えば、非同期処理（それぞれが別々に動いていて、干渉しない。Aが起こったことをトリガーにBが起こるのではなく、AとBはそれぞれ独立して動く状態）だと
				pushButtonFive()でbutton[5]をOnに
				→　releaseButtonFive()でbutton[5]をOffに
				っていう処理の間にtime.Sleepしないとbutton[5]がOnになったことを拾ってこられない可能性が高い（まだ実装していないからわからないけど）
				なので、
				pushButtonFive()でbutton[5]をOnに
				→　pushButtonFive()が実行されたということを感知して、GetKey()を実行する
				→　GetKey()処理が終わったら、releaseButtonFive()でbutton[5]をOffに
				という構成にしたい。
		また、time.Sleepで〇秒待つこと（「ユーザが〇秒押し続ける」という設定を行うこと）はソフトウェア上でテストすることではない。
			というか単なるユーザの操作の「設定」で、キーパッドに関するテスト最早関係ないし。
	// 【まとめ】
		・テストしたいのはあくまでステートの遷移であるため、仮想のPinの動作は一旦考えない
			・ユーザが、ドライバーの構造を理解した上でテストを行わなければいけなくなるため
				構造を理解したくないからドライバーがあるのに、ドライバーの動きを理解しなければテストできないのは本末転倒
		・time.Sleepを入れない構成にするように、チャネルを使用して同期させる
			・ステートの遷移のテストとしてtime.Sleepが入ってるのはよくわからない
			・time.Sleepがなくても（関数が値を取ってくるのを待たなくても）いいようにする→同期処理を実装
*/

// output device
type Led struct {
	name    string
	current string // "High" or "Low"
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

// // Pin

// type Pin struct {
// 	name string
// 	high bool
// }

// func (p Pin) Configure(interface{}) {}

// func (p Pin) Get() bool {
// 	return p.high
// }

// func (p *Pin) High() {
// 	p.high = true
// }

// func (p *Pin) Low() {
// 	p.high = false
// }

// func (p *Pin) Set(h bool) {
// 	p.high = h
// }

var button [16]string = [16]string{"Off", "Off", "Off", "Off", "Off", "Off", "Off", "Off", "Off", "Off", "Off", "Off", "Off", "Off", "Off", "Off"}

var ch chan bool    // pushButtonしたことを通知
var chGet chan bool // GetKey()により押されたボタンの値を取ったことを通知

// push button
func pushButtonOne() {
	log.Println("pushOne")
	button[0] = "On"
	ch <- true
}

func pushButtonTwo() {
	log.Println("pushTwo")
	button[1] = "On"
	ch <- true
}

func pushButtonThree() {
	log.Println("pushThree")
	button[2] = "On"
	ch <- true
}

func pushButtonA() {
	log.Println("pushA")
	button[3] = "On"
	ch <- true
}

func pushButtonFour() {
	button[4] = "On"
	ch <- true
}

func pushButtonFive() {
	button[5] = "On"
	ch <- true
}

func pushButtonSix() {
	button[6] = "On"
	ch <- true
}

func pushButtonB() {
	button[7] = "On"
	ch <- true
}

func pushButtonSeven() {
	button[8] = "On"
	ch <- true
}

func pushButtonEight() {
	button[9] = "On"
	ch <- true
}

func pushButtonNine() {
	button[10] = "On"
	ch <- true
}

func pushButtonC() {
	button[11] = "On"
	ch <- true
}

func pushButtonAsterisk() {
	button[12] = "On"
	ch <- true
}

func pushButtonZero() {
	button[13] = "On"
	ch <- true
	// fmt.Println("push:button[13] = ", button[13])
}

func pushButtonSharp() {
	button[14] = "On"
	ch <- true
}

func pushButtonD() {
	button[15] = "On"
	ch <- true
}

// releaseButton
func releaseButtonOne() {
	<-chGet
	log.Println("release ButtonOne")
	button[0] = "Off"
	ch <- false
}

func releaseButtonTwo() {
	<-chGet
	log.Println("release ButtonTwo")
	button[1] = "Off"
	ch <- false
}

func releaseButtonThree() {
	<-chGet
	log.Println("release ButtonThree")
	button[2] = "Off"
	ch <- false
}

func releaseButtonA() {
	<-chGet
	log.Println("release ButtonA")
	button[3] = "Off"
	ch <- false
}

func releaseButtonFour() {
	<-chGet
	button[4] = "Off"
	ch <- false
}

func releaseButtonFive() {
	<-chGet
	button[5] = "Off"
	ch <- false
}

func releaseButtonSix() {
	<-chGet
	button[6] = "Off"
	ch <- false
}

func releaseButtonB() {
	<-chGet
	button[7] = "Off"
	ch <- false
}

func releaseButtonSeven() {
	<-chGet
	button[8] = "Off"
	ch <- false
}

func releaseButtonEight() {
	<-chGet
	button[9] = "Off"
	ch <- false
}

func releaseButtonNine() {
	<-chGet
	button[10] = "Off"
	ch <- false
}

func releaseButtonC() {
	<-chGet
	button[11] = "Off"
	ch <- false
}

func releaseButtonAsterisk() {
	<-chGet
	button[12] = "Off"
	ch <- false
}

func releaseButtonZero() {
	<-chGet
	button[13] = "Off"
	ch <- false
	log.Println("releaseButtonZero")
	// log.Println("button[13] =", button[13])
}

func releaseButtonSharp() {
	<-chGet
	button[14] = "Off"
	ch <- false
}

func releaseButtonD() {
	<-chGet
	button[15] = "Off"
	ch <- false
}

type device struct {
	inputEnabled bool
	lastColumn   int
	lastRow      int
	mapping      [4][4]string
}

func newDevice() *device {
	result := &device{}
	return result
}

func (keypad *device) Configure() {

	keypad.mapping = [4][4]string{
		{"1", "2", "3", "A"},
		{"4", "5", "6", "B"},
		{"7", "8", "9", "C"},
		{"*", "0", "#", "D"},
	}

	keypad.inputEnabled = true
	keypad.lastColumn = -1
	keypad.lastRow = -1
}

func (keypad *device) GetIndices() (int, int) {
	// log.Println("keypad.inputEnabled = ", keypad.inputEnabled, "keypad.lastColumn = ", keypad.lastColumn, ", keypad.lastRow = ", keypad.lastRow)
	// log.Println("button[13] = ", button[13])

	if button[0] == "On" && keypad.inputEnabled {
		keypad.inputEnabled = false
		keypad.lastRow = 0
		keypad.lastColumn = 0

		return keypad.lastRow, keypad.lastColumn
	} else if button[0] == "Off" &&
		keypad.lastRow == 0 &&
		keypad.lastColumn == 0 &&
		!keypad.inputEnabled {
		keypad.inputEnabled = true
		// log.Println("button[0] == Off, keypad.inputEnabled = true")
	}

	if button[1] == "On" && keypad.inputEnabled {
		keypad.inputEnabled = false
		keypad.lastRow = 0
		keypad.lastColumn = 1

		return keypad.lastRow, keypad.lastColumn
	} else if button[1] == "Off" &&
		keypad.lastRow == 0 &&
		keypad.lastColumn == 1 &&
		!keypad.inputEnabled {
		keypad.inputEnabled = true
		// log.Println("button[1] == Off, keypad.inputEnabled = true")
	}

	if button[2] == "On" && keypad.inputEnabled {
		keypad.inputEnabled = false
		keypad.lastRow = 0
		keypad.lastColumn = 2

		return keypad.lastRow, keypad.lastColumn
	} else if button[2] == "Off" &&
		keypad.lastRow == 0 &&
		keypad.lastColumn == 2 &&
		!keypad.inputEnabled {
		keypad.inputEnabled = true
		// log.Println("button[2] == Off, keypad.inputEnabled = true")
	}

	if button[3] == "On" && keypad.inputEnabled {
		keypad.inputEnabled = false
		keypad.lastRow = 0
		keypad.lastColumn = 3

		return keypad.lastRow, keypad.lastColumn
	} else if button[3] == "Off" &&
		keypad.lastRow == 0 &&
		keypad.lastColumn == 3 &&
		!keypad.inputEnabled {
		keypad.inputEnabled = true
		// log.Println("button[3] == Off, keypad.inputEnabled = true")
	}

	if button[4] == "On" && keypad.inputEnabled {
		keypad.inputEnabled = false
		keypad.lastRow = 1
		keypad.lastColumn = 0

		return keypad.lastRow, keypad.lastColumn
	} else if button[4] == "Off" &&
		keypad.lastRow == 1 &&
		keypad.lastColumn == 0 &&
		!keypad.inputEnabled {
		keypad.inputEnabled = true
		// log.Println("button[4] == Off, keypad.inputEnabled = true")
	}

	if button[5] == "On" && keypad.inputEnabled {
		keypad.inputEnabled = false
		keypad.lastRow = 1
		keypad.lastColumn = 1

		return keypad.lastRow, keypad.lastColumn
	} else if button[5] == "Off" &&
		keypad.lastRow == 1 &&
		keypad.lastColumn == 1 &&
		!keypad.inputEnabled {
		keypad.inputEnabled = true
		// log.Println("button[5] == Off, keypad.inputEnabled = true")
	}

	if button[6] == "On" && keypad.inputEnabled {
		keypad.inputEnabled = false
		keypad.lastRow = 1
		keypad.lastColumn = 2

		return keypad.lastRow, keypad.lastColumn
	} else if button[6] == "Off" &&
		keypad.lastRow == 1 &&
		keypad.lastColumn == 2 &&
		!keypad.inputEnabled {
		keypad.inputEnabled = true
		// log.Println("button[6] == Off, keypad.inputEnabled = true")
	}

	if button[7] == "On" && keypad.inputEnabled {
		keypad.inputEnabled = false
		keypad.lastRow = 1
		keypad.lastColumn = 3

		return keypad.lastRow, keypad.lastColumn
	} else if button[7] == "Off" &&
		keypad.lastRow == 1 &&
		keypad.lastColumn == 3 &&
		!keypad.inputEnabled {
		keypad.inputEnabled = true
		// log.Println("button[7] == Off, keypad.inputEnabled = true")
	}

	if button[8] == "On" && keypad.inputEnabled {
		keypad.inputEnabled = false
		keypad.lastRow = 2
		keypad.lastColumn = 0

		return keypad.lastRow, keypad.lastColumn
	} else if button[8] == "Off" &&
		keypad.lastRow == 2 &&
		keypad.lastColumn == 0 &&
		!keypad.inputEnabled {
		keypad.inputEnabled = true
		// log.Println("button[8] == Off, keypad.inputEnabled = true")
	}

	if button[9] == "On" && keypad.inputEnabled {
		keypad.inputEnabled = false
		keypad.lastRow = 2
		keypad.lastColumn = 1

		return keypad.lastRow, keypad.lastColumn
	} else if button[9] == "Off" &&
		keypad.lastRow == 2 &&
		keypad.lastColumn == 1 &&
		!keypad.inputEnabled {
		keypad.inputEnabled = true
	}

	if button[10] == "On" && keypad.inputEnabled {
		keypad.inputEnabled = false
		keypad.lastRow = 2
		keypad.lastColumn = 2

		return keypad.lastRow, keypad.lastColumn
	} else if button[10] == "Off" &&
		keypad.lastRow == 2 &&
		keypad.lastColumn == 2 &&
		!keypad.inputEnabled {
		keypad.inputEnabled = true
	}

	if button[11] == "On" && keypad.inputEnabled {
		keypad.inputEnabled = false
		keypad.lastRow = 2
		keypad.lastColumn = 3

		return keypad.lastRow, keypad.lastColumn
	} else if button[11] == "Off" &&
		keypad.lastRow == 2 &&
		keypad.lastColumn == 3 &&
		!keypad.inputEnabled {
		keypad.inputEnabled = true
	}

	if button[12] == "On" && keypad.inputEnabled {
		keypad.inputEnabled = false
		keypad.lastRow = 3
		keypad.lastColumn = 0

		return keypad.lastRow, keypad.lastColumn
	} else if button[12] == "Off" &&
		keypad.lastRow == 3 &&
		keypad.lastColumn == 0 &&
		!keypad.inputEnabled {
		keypad.inputEnabled = true
	}

	if button[13] == "On" && keypad.inputEnabled {
		keypad.inputEnabled = false
		keypad.lastRow = 3
		keypad.lastColumn = 1

		return keypad.lastRow, keypad.lastColumn
	} else if button[13] == "Off" &&
		keypad.lastRow == 3 &&
		keypad.lastColumn == 1 &&
		!keypad.inputEnabled {
		keypad.inputEnabled = true
	}

	if button[14] == "On" && keypad.inputEnabled {
		keypad.inputEnabled = false
		keypad.lastRow = 3
		keypad.lastColumn = 2

		return keypad.lastRow, keypad.lastColumn
	} else if button[14] == "Off" &&
		keypad.lastRow == 3 &&
		keypad.lastColumn == 2 &&
		!keypad.inputEnabled {
		keypad.inputEnabled = true
	}

	if button[15] == "On" && keypad.inputEnabled {
		keypad.inputEnabled = false
		keypad.lastRow = 3
		keypad.lastColumn = 3

		return keypad.lastRow, keypad.lastColumn
	} else if button[15] == "Off" &&
		keypad.lastRow == 3 &&
		keypad.lastColumn == 3 &&
		!keypad.inputEnabled {
		keypad.inputEnabled = true
	}

	return -1, -1
}

var flagKey bool = false

func (keypad *device) GetKey() string {
	// log.Println("GetKey()実行")

	flagKey = <-ch
	// log.Println("flagKey =", flagKey)

	if flagKey {
		row, column := keypad.GetIndices()
		// fmt.Println("row =", row, "column =", column)

		if row == -1 && column == -1 {
			return "NoKeyPressed"
		} else {
			chGet <- true // GetKey()を実行したことをrelease関数に通知
			// log.Println("chGet <- true")
			return keypad.mapping[row][column]
		}
	}

	return "NoKeyPressed"
}

// define struct
var led = &Led{"led", "Low"}

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

func TestChannel01(t *testing.T) {
	ch = make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		flag := false
		log.Println("受信前:flag =", flag)
		flag = <-ch
		log.Println("受信前:flag =", flag)

		wg.Done()
	}()

	go func() {
		time.Sleep(2 * time.Second)
		pushButtonOne()
	}()

	wg.Wait()
}

func TestChannel02(t *testing.T) {
	// Success
	ch = make(chan bool)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for {
			flag := <-ch
			log.Println("flag =", flag)
		}
	}()

	// goroutine (user operation)
	go func() {
		pushButtonOne()
		releaseButtonOne()

		pushButtonTwo()
		releaseButtonTwo()

		pushButtonThree()
		releaseButtonThree()

		pushButtonA()
		releaseButtonA()

		wg.Done()
	}()

	wg.Wait()
}

func TestDevice01(t *testing.T) {
	// Success
	ch = make(chan bool)
	chGet = make(chan bool)

	passcode = "5678"
	keypad := newDevice()
	keypad.Configure()

	ConfigureDevice(led)
	ConfigureKeypad(keypad)
	ConfigureLog(logTest)

	var wg sync.WaitGroup
	wg.Add(1)

	// goroutine (base.go Task())
	go func() {
		for {
			Task() // GetKey()を繰り返し実行
		}
	}()

	// go func() {
	// 	for {
	// 		flagKey = <-ch // pushButtonされるまで待つ（pushButtonからchを受信するまで待ち続ける）
	// 	}
	// }()

	// goroutine (user operation)
	// enter correct
	go func() {
		pushButtonFive()
		releaseButtonFive()

		pushButtonSix()
		releaseButtonSix()

		pushButtonSeven()
		releaseButtonSeven()

		pushButtonEight()
		releaseButtonEight()

		time.Sleep(5 * time.Second)
		wg.Done()
	}()

	wg.Wait()
}

func TestDevice02(t *testing.T) {
	// Fail
	ch = make(chan bool)
	chGet = make(chan bool)

	passcode = "3210"
	keypad := newDevice()
	keypad.Configure()

	ConfigureDevice(led)
	ConfigureKeypad(keypad)
	ConfigureLog(logTest)

	var wg sync.WaitGroup
	wg.Add(1)

	// goroutine (base.go Task())
	go func() {
		for {
			Task()
		}
	}()

	go func() {
		for {
			flagKey = <-ch // pushButtonされるまで待つ（pushButtonからchを受信するまで待ち続ける）
		}
	}()

	// goroutine (user operation)
	// enter fail(3200)
	go func() {
		pushButtonThree()
		releaseButtonThree()

		pushButtonTwo()
		releaseButtonTwo()

		pushButtonZero()
		releaseButtonZero()

		log.Println("pushButtonZero 一回目完了")

		pushButtonOne()
		releaseButtonOne()

		time.Sleep(5 * time.Second)
		wg.Done()
	}()

	wg.Wait()
}

func TestDevice03(t *testing.T) {
	// Success
	// 全体の挙動確認用
	ch = make(chan bool)
	chGet = make(chan bool)

	passcode = "123A"
	keypad := newDevice()
	keypad.Configure()

	ConfigureDevice(led)
	ConfigureKeypad(keypad)
	ConfigureLog(logTest)

	var wg sync.WaitGroup
	wg.Add(1)

	// goroutine (base.go Task())
	go func() {
		for {
			Task() // GetKey()を繰り返し実行
		}
	}()

	go func() {
		for {
			flagKey = <-ch // pushButtonされるまで待つ（pushButtonからchを受信するまで待ち続ける）
		}
	}()

	// goroutine (user operation)
	go func() {
		// enter correct

		pushButtonOne()
		releaseButtonOne()

		pushButtonTwo()
		releaseButtonTwo()

		pushButtonThree()
		releaseButtonThree()

		pushButtonA()
		releaseButtonA()

		time.Sleep(5 * time.Second)
		wg.Done()
	}()

	wg.Wait()
}
