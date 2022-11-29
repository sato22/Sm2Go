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

/*
【interface内の関数の定義について（hoge_impl.go, main.go　で関係を持たせる）（2022/11/26）】
// １．仕様に基づいて、テストしたい内容を書きだす
// ２．hoge_impl.goにて、テストしたい仕様に合わせて、関数のインターフェースを定義
	// キッチンタイマーの場合、ディスプレイに表示する文字のフォントや文字の色に関する仕様はない
	// （例えば、10秒前になると文字が赤くなるとか）ので、ディスプレイに関する関数は表示に関するもののみで最小限に留める。
	// （ディスプレイの表示が更新されるかどうかは確認しなければならないため）
*/

/*
【tiny display について（2022/11/26）】
・Goで書かれたディスプレイのシミュレーター
Sm2go由来の構造だと、どうしてもStep()関数がでてくるしmainだけで完結しないので難しい…？
*/

/*
【個ゼミメモ_やること（2022/11/26）】
・実機での動作の確認
	・表示が更新されるか、文字がかぶって表示されないか
・キッチンタイマーの拡張機能を考える
	・現在使用中のm5stackはtinygoがwifi周りのドライバーを実装していないため、WioTerminalを使用する予定
	・ネットワークを使ってどこかのサイトの値を参照して、それを用いて何かするとかだと面白い？
	・Sm2Goがあれば、機能の追加も簡単にできること（拡張性の高さ）を伝えたい。
・タイマー割込みについて調べる
	・現在は、「一回for文を実行するごとに〇ms経過すると想定」して、〇回for文がまわったら一秒経過したことになるという考え方で実装している。
		・これは厳密には1秒ではない。Step()関数の処理時間も明確に決まっていないし、処理時間が長いときもあれば処理時間が短い時間もある。
			・それを一様に〇msとみなして実装していること
			・一秒とみなすカウントの数も適当に決定していること（Step()関数の実際の実行時間を図ってから設定していない）
				以上の主に二点の原因から、どうしても一秒とはズレが生じている。
		・これを解決するには、「タイマー割込み機能」を実装する必要がある？

	・キッチンタイマーのソースコードとか参考に？
	・clkを基に一秒経ったことを通知する？「interrupt」や「タイマー割込み」で要検索
		・参考URL：　https://edn.itmedia.co.jp/edn/articles/1509/25/news014.html
*/

/*
【来週(12/5)以降？】
・WioTerminalで同じものを作成し、実機が違ってもimplで定義したインターフェースの中身を変えれば簡単に対応できるんですよ～ということを伝えたい。
	・これはWioTerminalを入手次第。11/28以降に届く予定？なので、12月以降の個ゼミでかな。

個人的に
・githubの整理、mdファイル（Sm2goの説明書）の作成
*/

var (
	black = color.RGBA{0x00, 0x00, 0x00, 0xFF}
	white = color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}
	// red   = color.RGBA{255, 0, 0, 255}
	// blue  = color.RGBA{0, 0, 255, 255}
	// green = color.RGBA{0, 255, 0, 255}
)

// ターミナルエミュレータ用にPrintln関数を置き換え
type DebugStruct struct{}

var logTest = DebugStruct{}

func (l DebugStruct) Println(debstr string) {
	println(debstr)
}

// ブザーを鳴らす際の音階の設定
type note struct {
	tone     float64
	duration float64
}

var notes []note = []note{
	{buzzer.C3, buzzer.Whole},
}

// main文用にBeep()関数, Mute()関数を定義
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
}

func (b DebugBuzzer) Mute() {
	buzz.Toggle()
}

// main文用にPrintVal()関数(hoge_impl.goにて定義)を再定義
var (
	display *ili9341.Device
)

type DebugMonitor struct{}

var moni = DebugMonitor{}

var label *Label = NewLabel(10, 50)

func (m DebugMonitor) PrintVal(str string) {
	// 一度文字の上に白い画像を配置することで文字を消去してから再び表示するため、チラつく
	display.FillRectangle(0, 0, 320, 240, white)
	tinyfont.WriteLine(display, &freesans.Bold18pt7b, 0, 135, str, black)

	// ↓ 以下の形式だとチラつかないらしい？

	// label.FillScreen(white)
	// tinyfont.WriteLine(label, &freesans.Bold18pt7b, 0, 135, str, black)
	// display.DrawRGBBitmap(0, 100, label.Buf, label.W, label.H)
	// time.Sleep(1 * time.Second)
}

// Label周りは、sm2goパッケージに移行…？
// いやそれだとsm2goに関するものだけじゃなくなるからユーザにわかりにくいか…
// 作るなら別にLabelパッケージを？

type Label struct {
	Buf  []uint16
	W, H int16
}

func NewLabel(w, h int16) *Label {
	return &Label{
		Buf: make([]uint16, w*h),
		W:   w,
		H:   h,
	}
}

func (l *Label) Display() error {
	return nil
}

func (l *Label) Size() (x, y int16) {
	return l.W, l.H
}

func (l *Label) SetPixel(x, y int16, c color.RGBA) {
	l.Buf[x+y*l.W] = RGBATo565(c)
}

func (l *Label) FillScreen(c color.RGBA) {
	for i := range l.Buf {
		l.Buf[i] = RGBATo565(c)
	}
}

func RGBATo565(c color.RGBA) uint16 {
	r, g, b, _ := c.RGBA()

	return uint16((r & 0xF800) +
		((g & 0xFC00) >> 5) +
		((b & 0xF800) >> 11))
}

func main() {
	// ディスプレイの初期設定
	display = initdisplay.InitDisplay()
	display.FillScreen(white)

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
		time.Sleep(time.Millisecond * 50)
		// 5ms (Step()) + 50ms (time.Sleep()) = 一周当たり55msかかるとする
	}
}
