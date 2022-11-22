// main file

package main

import (
	"Sm2go/sample/kitchenTimer/kitchen"

	"tinygo.org/x/drivers/buzzer"
	"tinygo.org/x/drivers/examples/ili9341/initdisplay"
	"tinygo.org/x/drivers/ili9341"

	"tinygo.org/x/tinyfont"
	// "tinygo.org/x/tinyfont/examples/initdisplay"
	"tinygo.org/x/tinyfont/freesans"

	"image/color"
	"machine"
	"time"
)

var (
	black = color.RGBA{0x00, 0x00, 0x00, 0xFF}
	white = color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}
	// red   = color.RGBA{255, 0, 0, 255}
	// blue  = color.RGBA{0, 0, 255, 255}
	// green = color.RGBA{0, 255, 0, 255}
)

type DebugStruct struct{}

var logTest = DebugStruct{}

func (l DebugStruct) Println(debstr string) {
	println(debstr)
}

type note struct {
	tone     float64
	duration float64
}

var notes []note = []note{
	{buzzer.C3, buzzer.Quarter},
	{buzzer.Rest, buzzer.Eighth},
	{buzzer.D3, buzzer.Eighth},
	{buzzer.E3, buzzer.Quarter},
	{buzzer.Rest, buzzer.Eighth},
	{buzzer.C3, buzzer.Eighth},
	{buzzer.E3, buzzer.Quarter},
	{buzzer.C3, buzzer.Quarter},
	{buzzer.E3, buzzer.Half},
}

type DebugBuzzer struct{}

var bzr = DebugBuzzer{}

var (
	buzz buzzer.Device
)

func (b DebugBuzzer) Beep() {
	for _, n := range notes {
		buzz.Tone(n.tone, n.duration)
		time.Sleep(10 * time.Millisecond)
	}
	// buzz.Toggle()
}

func (b DebugBuzzer) Mute() {
	buzz.Toggle()
}

var (
	display *ili9341.Device
)

type DebugMonitor struct{}

var moni = DebugMonitor{}

func (m DebugMonitor) PrintVal(str string) {
	display.FillRectangle(0, 0, 320, 240, white)
	tinyfont.WriteLine(display, &freesans.Bold18pt7b, 10, 35, str, black)
}

func main() {
	// ディスプレイの初期設定
	display = initdisplay.InitDisplay()

	// アラームの初期設定
	bzrPin := machine.SPEAKER_PIN
	bzrPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	buzz = buzzer.New(bzrPin)

	// 左、真ん中、右のボタンの初期設定
	btnL := machine.BUTTON_A
	btnL.Configure(machine.PinConfig{Mode: machine.PinInput})

	btnM := machine.BUTTON_B
	btnM.Configure(machine.PinConfig{Mode: machine.PinInput})

	btnR := machine.BUTTON_C
	btnR.Configure(machine.PinConfig{Mode: machine.PinInput})

	// インターフェースで丸め込む
	kitchen.ConfigureLeftButton(btnL)
	kitchen.ConfigureMiddleButton(btnM)
	kitchen.ConfigureRightButton(btnR)

	kitchen.ConfigureMonitor(moni)

	kitchen.ConfigureAlarm(bzr)

	kitchen.ConfigureLog(logTest)

	for {
		kitchen.OneStep()
		time.Sleep(time.Millisecond * 10)
	}
}
