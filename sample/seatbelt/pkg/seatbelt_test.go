// This is a test file for testing state transitions

package pkg

import (
	"Sm2Go/sample/seatbelt/library"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"testing"
	"time"
)

//log.Println()
type DebugStruct struct{}

var logTest = DebugStruct{}

func (l DebugStruct) Println(debstr string) {
	log.Println(debstr)
}

// output device
type Led struct {
	name    string
	current string // "High" or "Low"
}

// Led
func (l *Led) High() {
	l.current = "High"
	log.Println(l.name, l.current)
}

func (l *Led) Low() {
	l.current = "Low"
	log.Println(l.name, l.current)
}

// vahicle
type vehicle struct {
	speed       float64
	speedSlice  []float64
	sliceIndex  int
	seatBelt    bool // trueだとfasten
	speedSensor bool // trueだと20km/h以上
}

func newCar() *vehicle {
	return &vehicle{
		speed:       0,
		speedSlice:  []float64{},
		sliceIndex:  0,
		seatBelt:    false,
		speedSensor: false,
	}
}

func (car *vehicle) readSpeed() {
	file, err := os.Open("carSpeed.csv") // 車速データをos.Openで開く
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	for {
		row, err := r.Read() // csvを1行ずつ読み込む
		if err == io.EOF {
			break // 読み込むべきデータが残っていない場合，Readはnil, io.EOFを返すため、ここでループを抜ける
		}

		for i, v := range row {
			switch i {
			case 0:
				f, err := strconv.ParseFloat(v, 64)
				if err != nil {
					fmt.Printf("%s\n", err.Error())
					return
				}
				car.speedSlice = append(car.speedSlice, f)
			}
		}
	}
}

func (car vehicle) GetSensor() bool {
	return car.speedSensor
}

func (car vehicle) GetSeatbelt() bool {
	return car.seatBelt
}

func (car *vehicle) FastenSeatbelt() {
	car.seatBelt = true
}

func (car *vehicle) UnFastenSeatbelt() {
	car.seatBelt = false
}

func (car *vehicle) DriveSpeedSensor() {
	if car.speed >= 20 {
		car.speedSensor = true
	} else {
		car.speedSensor = false
	}
}

func (car *vehicle) Accelerator() {
	car.speed = car.speedSlice[car.sliceIndex]
	car.sliceIndex++
	log.Println("Accelerator: car.speed =", car.speed)
}

// define struct
var led = &Led{"led", "Low"}

// init
func init() {
	if led.current == "High" {
		current = OnAlarm
	} else if led.current == "Low" {
		current = UnFasten
	}
	eod = Entry
}

func TestDevice01(t *testing.T) {
	car := newCar()
	car.readSpeed()

	ConfigureDevice(led)
	ConfigureVehicle(car)
	ConfigureLog(logTest)

	env := library.NewTestEnv() // TestEnv構造体

	env.Add(library.Continue, func() {
		for {
			Task()
			env.Sleep(10 * time.Millisecond)
		}
	},
	)

	// 車速とシートベルトするしないは別ルーチン
	// 車速はテスト毎で不変で、ユーザがシートベルトを付ける外すの時点が変わるというイメージ
	// 車速が羅列されたファイルを〇msごとに読み込んで車速を変えていくのが良さそう（車速のファイルは自作して）

	env.Add(library.Done, func() {
		for i := 0; i < len(car.speedSlice); i++ {
			env.Sleep(time.Second * 1)
			car.Accelerator()
		}
	})

	// Unfasten → OnAlarm への遷移確認
	env.Add(library.Done, func() {
		env.Sleep(time.Second * 10)

		car.FastenSeatbelt()

		env.Sleep(time.Second * 10)
	})

	env.Go(1)
}

func TestDevice02(t *testing.T) {
	car := newCar()
	car.readSpeed()

	ConfigureDevice(led)
	ConfigureVehicle(car)
	ConfigureLog(logTest)

	env := library.NewTestEnv() // TestEnv構造体

	env.Add(library.Continue, func() {
		for {
			Task()
			env.Sleep(10 * time.Millisecond)
		}
	},
	)

	// // Unfasten → OnAlarm への遷移確認
	env.Add(library.Done, func() {
		for i := 0; i < len(car.speedSlice); i++ {
			env.Sleep(time.Second * 1)
			car.Accelerator()
		}
	})

	// Unfasten → OnAlarm への遷移確認
	env.Add(library.Done, func() {
		env.Sleep(time.Second * 10)

		car.FastenSeatbelt()

		env.Sleep(time.Second * 3)

		car.UnFastenSeatbelt()

		env.Sleep(time.Second * 10)
	})

	env.Go(1)
}
