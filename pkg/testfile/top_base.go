package modelTest

// Please do not edit this file.

// package name to import

// constは別ファイル
const (
	debug = true
)

type TopState int

const (
	// 親と子で状態名が被らないよう，{接頭語}Onなどで記載
	On TopState = iota
	Off
)

// Eodは別ファイル
type Eod int

// constは別ファイル
const (
	entry Eod = iota
	do
	exit
)

// 接頭語の先頭文字は小文字で（同じパッケージ下にあるため．）
var topeod Eod
var topCurrentState TopState

// main関数やテスト関数で使用するグローバル関数を1つ
func Step() {
	topStep()
}

func topStep() {
	switch topCurrentState {
	case On:
		if topeod == Entry {
			onEntry()
			topeod = Do
		}
		if topeod == Do {
			onDo()
			if time3secCond() {
				topCurrentState = Off
				if debug {
					logger.Println("State is changed: On to Off")
				}
				topeod = Exit
			}
		}
		if topeod == Exit {
			onExit()
			topeod = Entry
		}
	case Off:
		if topeod == Entry {
			offEntry()
			topeod = Do
		}
		if topeod == Do {
			offDo()
			if correctKeyCond() {
				topCurrentState = On
				if debug {
					logger.Println("State is changed: Off to On")
				}
				topeod = Exit
			}
		}
		if topeod == Exit {
			offExit()
			topeod = Entry
		}
	}
}

func init() {
	topCurrentState = Off
	topeod = Entry
}