// This is a test file for testing state transitions

package ledbutton

import (
	"sync"
	"testing"
	"time"
	"log"
)

func TestStateTrans(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			time.Sleep(1 * time.Millisecond)
			task()
		}
		wg.Done()
	}()
	wg.Wait()
}

/*
Button型などの構造体を定義
構造体にメソッドを追加，Button.Get()でlog.Println("button pushed")などを実行し，
ボタンが押されたことをコンソール上で確認できるよう構成
goroutineにより，task()を回しながらユーザの操作をまわす
*/

type Led struct{
	name string
	current string // 現在の状態
}

type Button struct{
	name string
	push bool		// falseの時はrelease状態
	on bool			// OnボタンかOffボタンか
}

func (led Led) High() {
	led.current = "High"
	log.Println(led.name , led.current)
}

func (led Led) Low() {
	led.current = "Low"
	log.Println(led.name , led.current)
}

func (button Button) Push(led Led){
	button.push = true
	log.Println(button.name , "Push")

	if button.on{
		led.High()
	}else{
		led.Low()
	}
}

func (button Button) Release(){
	button.push = false
	log.Println(button.name , "Release")
}

func (p Button) Get() bool{
	return p.push
}

var led = Led{"led", "High"}
var leftButton = Button{"leftButton", false, true}
var rightButton = Button{"rightButton", false, false}

func TestDevice(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			// time.Sleep(1 * time.Millisecond)
			task()
		}
	}()

	go func(){
		log.Println("----------------led_test-------------")
		led.High()
		time.Sleep(1 * time.Second)
		led.Low()
		log.Println("----------------button_test-------------")
		// leftButtonPush
		time.Sleep(1 * time.Second)
		leftButton.Push(led)
		time.Sleep(500 * time.Millisecond)
		leftButton.Release()
		// rightButtonPush
		time.Sleep(1 * time.Second)
		rightButton.Push(led)
		time.Sleep(500 * time.Millisecond)
		rightButton.Release()
		// leftButtonPush dubble
		time.Sleep(1 * time.Second)
		leftButton.Push(led)
		time.Sleep(500 * time.Millisecond)
		leftButton.Release()
		time.Sleep(1 * time.Second)
		leftButton.Push(led)
		time.Sleep(500 * time.Millisecond)
		leftButton.Release()
		time.Sleep(1 * time.Second)
		wg.Done()
	}()
	wg.Wait()
}

