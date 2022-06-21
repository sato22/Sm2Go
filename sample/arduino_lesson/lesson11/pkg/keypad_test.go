// This is a test file for testing state transitions

package pkg

import (
	"Sm2Go/sample/arduino_lesson/lesson11_driver/library"
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
	// log.Println("button Five On")
}

func pushButtonSix() {
	button[6] = "On"
	// log.Println("button Six On")
}

func pushButtonB() {
	button[7] = "On"
}

func pushButtonSeven() {
	button[8] = "On"
	// log.Println("button Seven On")
}

func pushButtonEight() {
	button[9] = "On"
	// log.Println("button Eight On")
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
	// log.Println("button Five Off")
}

func releaseButtonSix() {
	button[6] = "Off"
	// log.Println("button Six Off")
}

func releaseButtonB() {
	button[7] = "Off"
}

func releaseButtonSeven() {
	button[8] = "Off"
	// log.Println("button Seven Off")
}

func releaseButtonEight() {
	button[9] = "Off"
	// log.Println("button Eight Off")
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

type device struct {
	mapping      [4][4]string
	inputEnabled bool
	lastColumn   int
	lastRow      int
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
}

func (keypad *device) GetIndices() (int, int) {
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

/*
// GetIndides()いらない？
	keypad_impl.go自体はGetKey()しか実行されていないし、GetIndices()の挙動はユーザに関係ない
	// GetKey()関数は、あくまで「その位置の値を返す」関数という扱いの方がいいかも
		// （ないかもしれないけど）Keypadの値を変更する時、GetKey()にべた書きだと変更しづらい
		// Configure()内でKeypadの値を定義して、GetKey()でとってくる構成の方が楽
*/

func (keypad *device) GetKey() string {
	row, column := keypad.GetIndices()

	if row == -1 && column == -1 {
		// log.Println("NoKeyPressed")
		return "NoKeyPressed"
	} else {
		// log.Println("row,column =", row, column)
		return keypad.mapping[row][column]
	}
}

/*
// velilogのclkのような構成に
〇msといった処理時間の設定
チャタリングのシミュレーションも
*/

// GetKey() pushButton()を待つ
// releaseBUtton() pushBUtton()を待つ（実際はGetKey()を待っている）
// pushBUtton() releaseButton()を待つ

// goroutineが重要なので、リアルタイムで簡潔に動作テストできることを主張
// データの受け渡し（データの共有？）をチャネルで（少し前にやったやつ）
// ある程度フレームワークのようなものを提供したい

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

func TestDevice01(t *testing.T) {
	// Success

	passcode = "5678"
	keypad := newDevice()
	keypad.Configure()

	// 構造体のフィールドで1単位あたりの秒数を設定（ユーザが自由に単位を設定できるように）
	// nano := library.Time{Scale: "nanoSec"}
	milli := library.Time{Scale: "milliSec"}
	sec := library.Time{Scale: "sec"}

	ConfigureDevice(led)
	ConfigureKeypad(keypad)
	ConfigureLog(logTest)

	// var wg sync.WaitGroup
	// wg.Add(1) // waitGroupがユーザにわかりづらい→syncが見えないように

	// goroutine (base.go Task())
	taskLoop := func() {
		for {
			// 〇秒ごとにTask()を実行（時間の割合で考える）
			Task() // Repeat GetKey()
			milli.Sleep(100)
			// 1サイクルの中でTask()が締める時間の割合
		}
	}

	// goroutine (user operation)
	// correct input (PassCode:5678, Entered:5678)
	push := func() {
		pushButtonFive()
		milli.Sleep(500)
		releaseButtonFive()

		pushButtonSix()
		milli.Sleep(500)
		releaseButtonSix()

		pushButtonSeven()
		milli.Sleep(500)
		releaseButtonSeven()

		pushButtonEight()
		milli.Sleep(500)
		releaseButtonEight()

		sec.Sleep(5) // interval until LED turns Off
		// wg.Done()    // Goを知らない人にも扱えるように→wg.Done()はノイズ
	}

	library.Go(taskLoop, push)

	// wg.Wait()
}

func TestDevice02(t *testing.T) {
	// Fail
	// ch = make(chan bool)
	// chGetKey = make(chan bool)

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
			Task() // Repeat GetKey()
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// goroutine (user operation)
	// wrong input (PassCode:3210, Entered:3212)
	go func() {
		pushButtonThree()
		time.Sleep(500 * time.Millisecond) // ボタンを0.5秒押し続ける
		releaseButtonThree()

		pushButtonTwo()
		time.Sleep(500 * time.Millisecond) // ボタンを0.5秒押し続ける
		releaseButtonTwo()

		pushButtonOne()
		time.Sleep(500 * time.Millisecond) // ボタンを0.5秒押し続ける
		releaseButtonOne()

		pushButtonTwo()
		time.Sleep(500 * time.Millisecond) // ボタンを0.5秒押し続ける
		releaseButtonTwo()

		time.Sleep(5 * time.Second)
		wg.Done()
	}()

	wg.Wait()
}

func TestDevice03(t *testing.T) {
	// Fail
	// Press the same button twice in a row
	// ch = make(chan bool)
	// chGetKey = make(chan bool)

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
			Task() // Repeat GetKey()
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// goroutine (user operation)
	// wrong input (PassCode:123A, Entered:1233)
	go func() {
		pushButtonOne()
		time.Sleep(500 * time.Millisecond) // ボタンを0.5秒押し続ける
		releaseButtonOne()

		pushButtonTwo()
		time.Sleep(500 * time.Millisecond) // ボタンを0.5秒押し続ける
		releaseButtonTwo()

		pushButtonThree()
		time.Sleep(500 * time.Millisecond) // ボタンを0.5秒押し続ける
		releaseButtonThree()

		pushButtonThree()
		time.Sleep(500 * time.Millisecond) // ボタンを0.5秒押し続ける
		releaseButtonThree()

		time.Sleep(5 * time.Second)
		wg.Done()
	}()

	wg.Wait()
}
