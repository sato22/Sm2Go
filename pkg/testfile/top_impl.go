package modelTest

// Please edit this file

// package name to import

type DebugLogger interface {
	Println(string)
}

var logger DebugLogger

func ConfigureLog(p DebugLogger) {
	logger = p
}
func alarmonEntry() {
	// nothing to do
}

func alarmonDo() {
	// nothing to do
}

func alarmonExit() {
	// nothing to do
}

func countdownEntry() {
	// nothing to do
}

func countdownDo() {
	// nothing to do
}

func countdownExit() {
	// nothing to do
}

func timersetEntry() {
	// nothing to do
}

func timersetDo() {
	// nothing to do
}

func timersetExit() {
	// nothing to do
}

func notsetEntry() {
	// nothing to do
}

func notsetDo() {
	// nothing to do
}

func notsetExit() {
	// nothing to do
}

func pushStartStopButtonCond() bool {
	// Please write the conditions under which a state transitions
	return true
}

func countdownEndCond() bool {
	// Please write the conditions under which a state transitions
	return true
}

func pushTimerButtonCond() bool {
	// Please write the conditions under which a state transitions
	return true
}
