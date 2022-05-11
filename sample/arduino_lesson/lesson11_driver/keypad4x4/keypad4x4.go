package keypad4x4

import (
	"Sm2Go/sample/arduino_lesson/lesson11_driver/pindriver"
	"fmt"
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
	columns      [4]pindriver.Pinif
	rows         [4]pindriver.Pinif
	mapping      [4][4]string
}

var pinDevice pindriver.Pinif

func ConfigurePin(p pindriver.Pinif) {
	pinDevice = p
}

// takes r4 -r1 pins and c4 - c1 pins
func NewDevice(r4, r3, r2, r1, c4, c3, c2, c1 pindriver.Pinif) Device {
	result := &device{}
	result.columns = [4]pindriver.Pinif{c4, c3, c2, c1}
	result.rows = [4]pindriver.Pinif{r4, r3, r2, r1}

	return result
}

// Configure sets the column pins as input and the row pins as output
func (keypad *device) Configure() {
	// PinConfigはコンストラクタ（構造体）を定義

	inputConfig := pindriver.PinConfig{Mode: pindriver.PinInputPullup}
	for i := range keypad.columns {
		keypad.columns[i].Configure(inputConfig)
	}

	outputConfig := pindriver.PinConfig{Mode: pindriver.PinOutput}
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
		// time.Sleep(1 * time.Microsecond)
		// 何に対応するからこのぐらいの値を入れてください
		// 書き換える部分をどのように提供するか、ドライバーのままで動かせないのか
		// 動いている仕組みがどうなっているか？何のために待っているのか？
		// 仕組みをまとめた図

		// fmt.Println("----- row index =", rowIndex, "-----")

		for columnIndex, _ := range keypad.columns {
			columnPin := keypad.columns[columnIndex]
			// time.Sleep(1 * time.Microsecond)
			fmt.Println("column index =", columnIndex, ", columnPin.Get() =", columnPin.Get(), "keypad.inputEnabled =", keypad.inputEnabled)

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
