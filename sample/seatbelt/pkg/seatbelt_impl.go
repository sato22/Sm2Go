// Please edit this file

package pkg

import "time"

// debuglog
type DebugLogger interface {
	Println(string)
}

var logger DebugLogger

func ConfigureLog(p DebugLogger) {
	logger = p
}

// interface
type Device interface {
	High()
	Low()
}

type Vehicle interface {
	GetSensor() bool
	GetSeatbelt() bool
	DriveSpeedSensor()
}

var dev Device
var car Vehicle

// configure
func ConfigureDevice(d Device) {
	dev = d
}

func ConfigureVehicle(v Vehicle) {
	car = v
}

func fastenEntry() {
	dev.Low()
}

func fastenDo() {
	// nothing to do
}

func fastenExit() {
	// nothing to do
}

func unfastenEntry() {
	// dev.Low()
}

func unfastenDo() {
	car.DriveSpeedSensor() // carのspeedを見続ける
}

func unfastenExit() {
	// nothing to do
}

func onengineEntry() {
	// nothing to do
}

func onengineDo() {
	// nothing to do
}

func onengineExit() {
	// nothing to do
}

func onalarmEntry() {
	dev.High()
}

func onalarmDo() {
	car.DriveSpeedSensor() // carのspeedを見続ける
}

func onalarmExit() {
	// nothing to do
}

func fastEntry() {
	// nothing to do
}

func fastDo() {
	car.DriveSpeedSensor() // carのspeedを見続ける
}

func fastExit() {
	// nothing to do
}

func time3secCond() bool {
	time.Sleep(time.Second * 3)
	return car.GetSensor()
}

func speedSensorOnCond() bool {
	// log.Println("speedSensorOnCond(), car.GetSensor() =", car.GetSensor())
	return car.GetSensor()
}

func fastenSeatbeltCond() bool {
	// log.Println("fastenSeatbeltCond(), car.GetSeatbelt() =", car.GetSeatbelt())
	return car.GetSeatbelt()
}

func speedSensorOffCond() bool {
	// log.Println("speedSensorOffCond(), car.GetSensor() =", car.GetSensor())
	return !car.GetSensor()
}

func unfastenSeatbeltCond() bool {
	// log.Println("unfastenSeatbeltCond(), car.GetSeatbelt() =", car.GetSeatbelt())
	return !car.GetSeatbelt()
}
