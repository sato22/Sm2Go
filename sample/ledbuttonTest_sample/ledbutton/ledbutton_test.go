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

// output device
type Led struct{
	name string
	current string // 現在の状態
}

type Button struct{
	name string
	push bool		// falseの時はrelease状態
}

func (l *Led) High() {
	l.current = "High"
	log.Println(l.name , l.current)
}

func (l *Led) Low() {
	l.current = "Low"
	log.Println(l.name , l.current)
}

func (b *Button) Push(){
	b.push = true
	log.Println(b.name , "Push")
}

func (b *Button) Release(){
	b.push = false
	log.Println(b.name , "Release")
}

func (b Button) Get() bool{
	return b.push
}

var led = &Led{"led", "High"}
var leftButton = &Button{"leftButton", false}
var rightButton = &Button{"rightButton", false}

func TestDevice(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			time.Sleep(50 * time.Millisecond)
			task()
			// log.Println("test_task()")
		}
	}()

	go func(){
		log.Println("----------------button_test-------------")
		// leftButtonPush
		time.Sleep(1 * time.Second)
		leftButton.Push()
		time.Sleep(500 * time.Millisecond)
		leftButton.Release()
		
		// rightButtonPush
		time.Sleep(1 * time.Second)
		rightButton.Push()
		time.Sleep(500 * time.Millisecond)
		rightButton.Release()

		// leftButtonPush dubble
		time.Sleep(1 * time.Second)
		leftButton.Push()
		time.Sleep(500 * time.Millisecond)
		leftButton.Release()
		time.Sleep(1 * time.Second)
		leftButton.Push()
		time.Sleep(500 * time.Millisecond)
		leftButton.Release()
		time.Sleep(1 * time.Second)
		wg.Done()
	}()
	wg.Wait()
}

