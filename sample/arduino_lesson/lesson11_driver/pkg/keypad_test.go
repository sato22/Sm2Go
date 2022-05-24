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

var button [16]string = [16]string{"Off", "Off", "Off", "Off", "Off", "Off", "Off", "Off", "Off", "Off", "Off", "Off", "Off", "Off", "Off", "Off"}

var ch chan bool       // pushButtonしたことを通知
var chGetKey chan bool // GetKey()により押されたボタンの値を取ったことを通知

// push button
func pushButtonOne() {
	button[0] = "On"
	ch <- true
}

func pushButtonTwo() {
	button[1] = "On"
	ch <- true
}

func pushButtonThree() {
	button[2] = "On"
	ch <- true
}

func pushButtonA() {
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
	<-chGetKey
	button[0] = "Off"
	ch <- false
}

func releaseButtonTwo() {
	<-chGetKey
	button[1] = "Off"
	ch <- false
}

func releaseButtonThree() {
	<-chGetKey
	button[2] = "Off"
	ch <- false
}

func releaseButtonA() {
	<-chGetKey
	button[3] = "Off"
	ch <- false
}

func releaseButtonFour() {
	<-chGetKey
	button[4] = "Off"
	ch <- false
}

func releaseButtonFive() {
	<-chGetKey
	button[5] = "Off"
	ch <- false
}

func releaseButtonSix() {
	<-chGetKey
	button[6] = "Off"
	ch <- false
}

func releaseButtonB() {
	<-chGetKey
	button[7] = "Off"
	ch <- false
}

func releaseButtonSeven() {
	<-chGetKey
	button[8] = "Off"
	ch <- false
}

func releaseButtonEight() {
	<-chGetKey
	button[9] = "Off"
	ch <- false
}

func releaseButtonNine() {
	<-chGetKey
	button[10] = "Off"
	ch <- false
}

func releaseButtonC() {
	<-chGetKey
	button[11] = "Off"
	ch <- false
}

func releaseButtonAsterisk() {
	<-chGetKey
	button[12] = "Off"
	ch <- false
}

func releaseButtonZero() {
	<-chGetKey
	button[13] = "Off"
	ch <- false
}

func releaseButtonSharp() {
	<-chGetKey
	button[14] = "Off"
	ch <- false
}

func releaseButtonD() {
	<-chGetKey
	button[15] = "Off"
	ch <- false
}

type device struct {
	mapping [4][4]string
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
	if button[0] == "On" {
		return 0, 0
	}

	if button[1] == "On" {
		return 0, 1
	}

	if button[2] == "On" {
		return 0, 2
	}

	if button[3] == "On" {
		return 0, 3
	}

	if button[4] == "On" {
		return 1, 0
	}

	if button[5] == "On" {
		return 1, 1
	}

	if button[6] == "On" {
		return 1, 2
	}

	if button[7] == "On" {
		return 1, 3
	}

	if button[8] == "On" {
		return 2, 0
	}

	if button[9] == "On" {
		return 2, 1
	}

	if button[10] == "On" {
		return 2, 2
	}

	if button[11] == "On" {
		return 2, 3
	}

	if button[12] == "On" {
		return 3, 0
	}

	if button[13] == "On" {
		return 3, 1
	}

	if button[14] == "On" {
		return 3, 2
	}

	if button[15] == "On" {
		return 3, 3
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

var pushFlag bool = false

func (keypad *device) GetKey() string {
	pushFlag = <-ch
	row, column := keypad.GetIndices()

	if pushFlag {
		if row == -1 && column == -1 {
			return "NoKeyPressed"
		} else {
			chGetKey <- true // GetKey()を実行したことをrelease関数に通知
			return keypad.mapping[row][column]
		}
	}

	return "NoKeyPressed"
}

// GetKey() pushButton()を待つ
// releaseBUtton() pushBUtton()を待つ（実際はGetKey()を待っている）
// pushBUtton() releaseButton()を待つ

// 図を作図してみよう

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
	ch = make(chan bool)
	chGetKey = make(chan bool)

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
			Task() // Repeat GetKey()
		}
	}()

	// goroutine (user operation)
	// correct input (PassCode:5678, Entered:5678)
	go func() {
		pushButtonFive()
		releaseButtonFive()

		pushButtonSix()
		releaseButtonSix()

		pushButtonSeven()
		releaseButtonSeven()

		pushButtonEight()
		releaseButtonEight()

		time.Sleep(5 * time.Second) // interval until LED turns Off
		wg.Done()
	}()

	wg.Wait()
}

func TestDevice02(t *testing.T) {
	// Fail
	ch = make(chan bool)
	chGetKey = make(chan bool)

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
		}
	}()

	// goroutine (user operation)
	// wrong input (PassCode:3210, Entered:3212)
	go func() {
		pushButtonThree()
		releaseButtonThree()

		pushButtonTwo()
		releaseButtonTwo()

		pushButtonOne()
		releaseButtonZero()

		pushButtonTwo()
		releaseButtonZero()

		time.Sleep(5 * time.Second)
		wg.Done()
	}()

	wg.Wait()
}

func TestDevice03(t *testing.T) {
	// Fail
	// Press the same button twice in a row
	ch = make(chan bool)
	chGetKey = make(chan bool)

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
		}
	}()

	// goroutine (user operation)
	// wrong input (PassCode:123A, Entered:1233)
	go func() {
		pushButtonOne()
		releaseButtonOne()

		pushButtonTwo()
		releaseButtonTwo()

		pushButtonThree()
		releaseButtonThree()

		pushButtonThree()
		releaseButtonA()

		time.Sleep(5 * time.Second)
		wg.Done()
	}()

	wg.Wait()
}
