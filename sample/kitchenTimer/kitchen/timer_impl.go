package kitchen

import (
	"fmt"
	"time"
)

type DebugLogger interface {
	Println(string)
}

var logger DebugLogger

func ConfigureLog(p DebugLogger) {
	logger = p
}

type Switch interface {
	Get() bool
}

var buttonL Switch
var buttonM Switch
var buttonR Switch

func ConfigureLeftButton(s Switch) {
	buttonL = s
}

func ConfigureMiddleButton(s Switch) {
	buttonM = s
}

func ConfigureRightButton(s Switch) {
	buttonR = s
}

// interface Monitor
// ディスプレイに文字が表示され、表示が更新されることを確認するためのインターフェース
type Monitor interface {
	PrintVal(string)
}

var display Monitor

func ConfigureMonitor(m Monitor) {
	display = m
}

// interface Alarm
// アラームがなることを確認するためのインターフェース
type Alarm interface {
	Beep()
	Mute()
}

var alarm Alarm

func ConfigureAlarm(a Alarm) {
	alarm = a
}

var mm int = 0
var ss int = 0

var inputEnabled bool = true

func alarmonEntry() {
	alarm.Beep()
}

func alarmonDo() {
	// nothing to do
}

func alarmonExit() {
	alarm.Mute()
}

func countdownEntry() {
	// nothing to do
}

func countdownDo() {
	if buttonR.Get() {
		time.Sleep(time.Second * 1)
		ss--
		if ss == -1 {
			if mm == 0 {
				ss = 0
			} else {
				ss = 59
				mm--
			}
		}

		if ss%2 == 1 {
			display.PrintVal(fmt.Sprintf("%d : %d", mm, ss))
		} else {
			display.PrintVal(fmt.Sprintf("%d   %d", mm, ss))
		}
	}
}

func countdownExit() {
	// nothing to do
}

func timersetEntry() {
	// nothing to do
}

func timersetDo() {
	if !buttonL.Get() {
		mm++
		if mm == 60 {
			mm = 0
		}
		display.PrintVal(fmt.Sprintf("%d : %d", mm, ss))
	} else if !buttonM.Get() {
		ss++
		if ss == 60 {
			ss = 0
		}
		display.PrintVal(fmt.Sprintf("%d : %d", mm, ss))
	}
}

func timersetExit() {
	// nothing to do
}

func notsetEntry() {
	display.PrintVal("00 : 00")
}

func notsetDo() {
	// nothing to do
}

func notsetExit() {
	// nothing to do
}

func pushStartStopButtonCond() bool {
	if !buttonR.Get() && inputEnabled {
		inputEnabled = false
		return true
	} else if buttonR.Get() {
		inputEnabled = true
	}
	return false
}

func countdownEndCond() bool {
	return (ss == 0 && mm == 0)
}

func pushTimerButtonCond() bool {
	return (!buttonL.Get() || !buttonM.Get())
}
