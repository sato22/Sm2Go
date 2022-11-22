package modelTest

// Please do not edit this file.

// package name to import

type GrandchildState int

const (
	OffEntered GrandchildState = iota
	OffEmpty
)

var grandChildEod Eod
var grandchildCurrentState GrandchildState

func GrandchildStep() {
	switch grandchildCurrentState {
	case OffEntered:
		if grandchildeod == Entry {
			offenteredEntry()
			grandchildeod = Do
		}
		if grandchildeod == Do {
			offenteredDo()
			if keyFailedCond() {
				grandchildCurrentState = OffEmpty
				if debug {
					logger.Println("State is changed: OffEntered to OffEmpty")
				}
				grandchildeod = Exit
			}
		}
		if grandchildeod == Exit {
			offenteredExit()
			grandchildeod = Entry
		}
	case OffEmpty:
		if grandchildeod == Entry {
			offemptyEntry()
			grandchildeod = Do
		}
		if grandchildeod == Do {
			offemptyDo()
			if keyEnteredCond() {
				grandchildCurrentState = OffEntered
				if debug {
					logger.Println("State is changed: OffEmpty to OffEntered")
				}
				grandchildeod = Exit
			}
		}
		if grandchildeod == Exit {
			offemptyExit()
			grandchildeod = Entry
		}
	}
}

func init() {
	grandchildCurrentState = OffEmpty
	grandchildeod = Entry
}
