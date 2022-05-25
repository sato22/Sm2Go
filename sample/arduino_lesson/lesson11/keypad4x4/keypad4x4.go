package keypad4x4

import (
	"machine"
)

/*
・driverは自作想定
	→　driver自体がテストできるべき
・machine.Pinのinterfaceを持った構造体を定義
*/

// Device is used as 4x4 keypad driver
type Device interface {
	Configure()
	GetKey() string
	GetIndices() (int, int)
}

// device is a driver for 4x4 keypads
type device struct {
	inputEnabled bool
	lastColumn   int
	lastRow      int
	columns      [4]machine.Pin
	rows         [4]machine.Pin
	mapping      [4][4]string // keypadの値として"A"や"#"等の文字を扱うため
}

// takes r4 -r1 pins and c4 - c1 pins
func NewDevice(r4, r3, r2, r1, c4, c3, c2, c1 machine.Pin) Device {
	result := &device{}
	result.columns = [4]machine.Pin{c4, c3, c2, c1}
	result.rows = [4]machine.Pin{r4, r3, r2, r1}

	return result
}

// Configure sets the column pins as input and the row pins as output
func (keypad *device) Configure() {
	inputConfig := machine.PinConfig{Mode: machine.PinInputPullup}
	for i := range keypad.columns {
		keypad.columns[i].Configure(inputConfig)
	}

	outputConfig := machine.PinConfig{Mode: machine.PinOutput}
	for i := range keypad.rows {
		keypad.rows[i].Configure(outputConfig)
		keypad.rows[i].High()
	}

	keypad.mapping = [4][4]string{
		{"1", "2", "3", "A"},
		{"4", "5", "6", "B"},
		{"7", "8", "9", "C"},
		{"*", "0", "#", "D"},
	}

	keypad.inputEnabled = true
	keypad.lastColumn = -1
	keypad.lastRow = -1
}

// GetKey returns the code for the given key.
// The codes start with 0 at the upper left end of the keypad and end with 15 at the lower right end of the keypad
// Example:
// 0	1	2	3
// 4	5	6	7
// 8	9	10	11
// 12	13	14	15
// returns 255 for no keyPressed
func (keypad *device) GetKey() string {
	row, column := keypad.GetIndices()
	if row == -1 && column == -1 {
		return "NoKeyPressed"
	}

	return keypad.mapping[row][column]
}

// GetIndices returns the position of the pressed key
func (keypad *device) GetIndices() (int, int) {
	for rowIndex, rowPin := range keypad.rows {
		rowPin.Low()

		// fmt.Println("----- row index =", rowIndex, "-----")

		for columnIndex, _ := range keypad.columns {
			columnPin := keypad.columns[columnIndex]
			// fmt.Println("column index =", columnIndex, ", columnPin.Get() =", columnPin.Get(), "keypad.inputEnabled =", keypad.inputEnabled)

			if !columnPin.Get() && keypad.inputEnabled {
				// fmt.Println("---if---", "index =", columnIndex, ", columnPin.Get() =", columnPin.Get())

				keypad.inputEnabled = false

				keypad.lastColumn = columnIndex
				keypad.lastRow = rowIndex

				// fmt.Println("keypad.lastColumn =", keypad.lastColumn)
				// fmt.Println("keypad.lastRow =", keypad.lastRow)

				return keypad.lastRow, keypad.lastColumn
			}

			if columnPin.Get() &&
				columnIndex == keypad.lastColumn &&
				rowIndex == keypad.lastRow &&
				!keypad.inputEnabled {
				keypad.inputEnabled = true
				// fmt.Println("----------if-----------     columnPin.Get() =", columnPin.Get(), "columnindex =", columnIndex, "rowindex =", rowIndex, "keypad.inputEnabled =", keypad.inputEnabled)
			}
		}

		rowPin.High()
	}

	return -1, -1
}
