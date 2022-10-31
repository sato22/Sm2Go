// impl.go
// Please edit this file

package model

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

// ここstateごとに出力するように

{{range .}}
func {{.StateName}}Entry() {
	// nothing to do
}
	

func {{.StateName}}Do() {
	// nothing to do
}
	

func {{.StateName}}Exit() {
	// nothing to do
}
{{end}}

// ここもEventごとに出力するように
{{range .}}
func {{.EventName}}Cond() bool {
	// Please write the conditions under which a state transitions
	return true
}
{{end}}