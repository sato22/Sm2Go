// Please do not edit this file.

package ledbutton

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

func Task() {
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
if debug {
logger.Println("State is changed: On to Off")
}
eod = Exit
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
if pushLightButtonCond() {
current = On
if debug {
logger.Println("State is changed: Off to On")
}
eod = Exit
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
