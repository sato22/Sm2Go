// Please edit this file

package keypad_easy

import (
// package name to import
)

type DebugLogger interface {
Println(string)
}

var logger DebugLogger

func ConfigureLog(p DebugLogger) {
logger = p
}

func onEntry() {
// nothing to do
}

func onDo() {
// nothing to do
}

func onExit() {
// nothing to do
}

func offEntry() {
// nothing to do
}

func offDo() {
// nothing to do
}

func offExit() {
// nothing to do
}

func time3secCond() bool {
// Please write the conditions under which a state transitions
return true
}

func correctKeyCond() bool {
// Please write the conditions under which a state transitions
return true
}

// Please edit this file

package keypad_easy

import (
// package name to import
)

type DebugLogger interface {
Println(string)
}

var logger DebugLogger

func ConfigureLog(p DebugLogger) {
logger = p
}

func offenteredEntry() {
// nothing to do
}

func offenteredDo() {
// nothing to do
}

func offenteredExit() {
// nothing to do
}

func offemptyEntry() {
// nothing to do
}

func offemptyDo() {
// nothing to do
}

func offemptyExit() {
// nothing to do
}

func keyFailedCond() bool {
// Please write the conditions under which a state transitions
return true
}

func keyEnteredCond() bool {
// Please write the conditions under which a state transitions
return true
}

