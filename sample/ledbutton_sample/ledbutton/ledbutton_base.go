// Please do not edit this file.

package main

import (
// package name to import
)

type State int

const (
	On State = iota
	Off
)

type Eod int

const (
	Entry Eod = iota
	Do
	Exit
)

var eod Eod
var current State

func task() {
	switch current {
	case On:
		if eod == Entry {
			onEntry()
			eod = Do
		}
		if eod == Do {
			onDo()
			if pushRightButtonCond() {
				current = Off
				eod = Entry
			}
		}
		if eod == Exit {
			onExit()
			eod = Entry
		}
	case Off:
		if eod == Entry {
			offEntry()
			eod = Do
		}
		if eod == Do {
			offDo()
			if pushLeftButtonCond() {
				current = On
				eod = Entry
			}
		}
		if eod == Exit {
			offExit()
			eod = Entry
		}
	}
}

func init() {
	current = On
	eod = Entry
}
