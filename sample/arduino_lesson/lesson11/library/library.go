package library

import (
	"sync"
	"time"
)

// ・wg.wait系（Go言語を知らないユーザが全く書けない構文）を消したい
// ・time.sleepを別の関数として定義、別ファイルに定義して、オブジェクトのメソッドとして実行できるように。
// ・SESのポスター形式確認
// ・アトミックがどうちゃら…

type Time struct {
	Scale string
}

func (t Time) Sleep(h time.Duration) {
	switch t.Scale {
	case "nanoSec":
		time.Sleep(time.Nanosecond * h)
	case "microSec":
		time.Sleep(time.Microsecond * h)
	case "milliSec":
		time.Sleep(time.Millisecond * h)
	case "sec":
		time.Sleep(time.Second * h)
	}
}

func Go(fn ...func()) {
	var slice []func()

	for _, f := range fn {
		slice = append(slice, f)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	// それぞれにgoroutineを割り当てる
	for i, f := range fn {
		go f()
		if i == len(fn)-1 {
			// 最後→ユーザの操作を表した関数であること前提
			go func() {
				defer wg.Done()
				f()
			}()
		}
	}

	wg.Wait()
}

// func Go(fn1 func(), fn2 func()) {
// 	// 可変長な引数として扱ったほうがいい？
// 	slice := []func(){fn1, fn2}

// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go fn1()
// 	go func() {
// 		defer wg.Done()
// 		fn2()
// 	}()
// 	wg.Wait()
// }
