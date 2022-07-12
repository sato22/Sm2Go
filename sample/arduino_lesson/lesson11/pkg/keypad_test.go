// This is a test file for testing state transitions

package pkg

import (
	"Sm2Go/sample/arduino_lesson/lesson11_driver/library"
	"log"
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

type device struct {
	mapping      [4][4]string
	button       [4][4]string
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

	keypad.button = [4][4]string{
		{"Off", "Off", "Off", "Off"},
		{"Off", "Off", "Off", "Off"},
		{"Off", "Off", "Off", "Off"},
		{"Off", "Off", "Off", "Off"},
	}

	keypad.inputEnabled = true
}

func (keypad *device) GetIndices() (int, int) {
	for rowIndex, buttonArray := range keypad.button {
		for columnIndex, buttonValue := range buttonArray {
			if buttonValue == "On" && keypad.inputEnabled {
				keypad.inputEnabled = false

				keypad.lastColumn = columnIndex
				keypad.lastRow = rowIndex

				return rowIndex, columnIndex
			}

			if buttonValue == "Off" &&
				columnIndex == keypad.lastColumn &&
				rowIndex == keypad.lastRow &&
				!keypad.inputEnabled {
				keypad.inputEnabled = true
			}
		}
	}

	// OFFの場合にkeypad.inputEnabledがtrueになる

	return -1, -1
}

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

// push button
func (keypad *device) pushButtonOne() {
	keypad.button[0][0] = "On"
}

func (keypad *device) pushButtonTwo() {
	keypad.button[0][1] = "On"
}

func (keypad *device) pushButtonThree() {
	keypad.button[0][2] = "On"
}

func (keypad *device) pushButtonA() {
	keypad.button[0][3] = "On"
}

func (keypad *device) pushButtonFour() {
	keypad.button[1][0] = "On"
}

func (keypad *device) pushButtonFive() {
	keypad.button[1][1] = "On"
}

func (keypad *device) pushButtonSix() {
	keypad.button[1][2] = "On"
}

func (keypad *device) pushButtonB() {
	keypad.button[1][3] = "On"
}

func (keypad *device) pushButtonSeven() {
	keypad.button[2][0] = "On"
}

func (keypad *device) pushButtonEight() {
	keypad.button[2][1] = "On"
}

func (keypad *device) pushButtonNine() {
	keypad.button[2][2] = "On"
}

func (keypad *device) pushButtonC() {
	keypad.button[2][3] = "On"
}

func (keypad *device) pushButtonAsterisk() {
	keypad.button[3][0] = "On"
}

func (keypad *device) pushButtonZero() {
	keypad.button[3][1] = "On"
}

func (keypad *device) pushButtonSharp() {
	keypad.button[3][2] = "On"
}

func (keypad *device) pushButtonD() {
	keypad.button[3][3] = "On"
}

// releaseButton
func (keypad *device) releaseButtonOne() {
	keypad.button[0][0] = "Off"
}

func (keypad *device) releaseButtonTwo() {
	keypad.button[0][1] = "Off"
}

func (keypad *device) releaseButtonThree() {
	keypad.button[0][2] = "Off"
}

func (keypad *device) releaseButtonA() {
	keypad.button[0][3] = "Off"
}

func (keypad *device) releaseButtonFour() {
	keypad.button[1][0] = "Off"
}

func (keypad *device) releaseButtonFive() {
	keypad.button[1][1] = "Off"
	// log.Println("button Five Off")
}

func (keypad *device) releaseButtonSix() {
	keypad.button[1][2] = "Off"
	// log.Println("button Six Off")
}

func (keypad *device) releaseButtonB() {
	keypad.button[1][3] = "Off"
}

func (keypad *device) releaseButtonSeven() {
	keypad.button[2][0] = "Off"
	// log.Println("button Seven Off")
}

func (keypad *device) releaseButtonEight() {
	keypad.button[2][1] = "Off"
	// log.Println("button Eight Off")
}

func (keypad *device) releaseButtonNine() {
	keypad.button[2][2] = "Off"
}

func (keypad *device) releaseButtonC() {
	keypad.button[2][3] = "Off"
}

func (keypad *device) releaseButtonAsterisk() {
	keypad.button[3][0] = "Off"
	// keypad.inputEnabled = true
}

func (keypad *device) releaseButtonZero() {
	keypad.button[3][1] = "Off"
}

func (keypad *device) releaseButtonSharp() {
	keypad.button[3][2] = "Off"
}

func (keypad *device) releaseButtonD() {
	keypad.button[3][3] = "Off"
}

/*
// velilogのclkのような構成に
〇msといった処理時間の設定
チャタリングのシミュレーションも
*/

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

	passcode = "5662"
	keypad := newDevice()
	keypad.Configure()

	ConfigureDevice(led)
	ConfigureKeypad(keypad)
	ConfigureLog(logTest)

	env := library.NewTestEnv() // TestEnv構造体

	// goroutine (base.go Task())
	env.Add(library.Continue, func() {
		for {
			// 〇秒ごとにTask()を実行（1サイクルの中でTask()が締める時間の割合で考える）
			Task()
			env.Sleep(time.Millisecond * 10)
		}
	},
	)

	// goroutine (user operation)
	// correct input (PassCode:5662, Entered:5662)
	env.Add(library.Done, func() {
		keypad.pushButtonFive()
		env.Sleep(time.Millisecond * 50)
		keypad.releaseButtonFive()

		env.Sleep(time.Millisecond * 10) // 次のボタンを押すまでのインターバル

		keypad.pushButtonSix()
		env.Sleep(time.Millisecond * 50)
		keypad.releaseButtonSix()

		env.Sleep(time.Millisecond * 10)

		keypad.pushButtonSix()
		env.Sleep(time.Millisecond * 50)
		keypad.releaseButtonSix()

		env.Sleep(time.Millisecond * 10)

		keypad.pushButtonTwo()
		env.Sleep(time.Millisecond * 50) //
		keypad.releaseButtonTwo()

		env.Sleep(time.Second * 1) // interval until LED turns Off
	},
	)

	env.Go(10)
}

func TestDevice02(t *testing.T) {
	// Fail

	passcode = "3210"
	keypad := newDevice()
	keypad.Configure()

	ConfigureDevice(led)
	ConfigureKeypad(keypad)
	ConfigureLog(logTest)

	env := library.NewTestEnv() // TestEnv構造体

	// goroutine (base.go Task())
	env.Add(library.Continue, func() {
		for {
			// 〇秒ごとにTask()を実行（時間の割合で考える）
			Task() // Repeat GetKey()
			env.Sleep(time.Millisecond * 10)
		}
	},
	)

	// goroutine (user operation)
	// correct input (PassCode:3210, Entered:3212)
	env.Add(library.Done, func() {
		keypad.pushButtonThree()
		env.Sleep(time.Millisecond * 50)
		keypad.releaseButtonThree()

		keypad.pushButtonTwo()
		env.Sleep(time.Millisecond * 50)
		keypad.releaseButtonTwo()

		keypad.pushButtonOne()
		env.Sleep(time.Millisecond * 50)
		keypad.releaseButtonOne()

		keypad.pushButtonTwo()
		env.Sleep(time.Millisecond * 50)
		keypad.releaseButtonTwo()
	},
	)

	env.Go(10)
}
