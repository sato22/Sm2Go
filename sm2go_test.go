package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

/*
func TestState(t *testing.T) {
	// State構造体の確認
	state := State{
		Name:      "Wait",
		EntryProg: "wait_test_EntryProg",
		DoProg:    "wait_test_DoProg",
		ExitProg:  "wait_test_ExitProg",
	}

	// State構造体を表示
	fmt.Println(state)
}

func TestEvent(t *testing.T) {
	// Event構造体の確認
	event := Event{
		Name: "push_button",
		Cond: "ecrobot_get_touch_sensor(NXT_PORT_TOUCH)==1",
	}

	// Event構造体を表示
	fmt.Println(event)
}

func TestTransition(t *testing.T) {
	// Transition構造体の確認
	current_state := State{
		Name:      "Wait",
		EntryProg: "wait_test_EntryProg",
		DoProg:    "wait_test_DoProg",
		ExitProg:  "wait_test_ExitProg",
	}

	next_state := State{
		Name:      "Run",
		EntryProg: "run_test_EntryProg",
		DoProg:    "run_test_DoProg",
		ExitProg:  "run_test_ExitProg",
	}

	eve := Event{
		Name: "push_button",
		Cond: "ecrobot_get_touch_sensor(NXT_PORT_TOUCH)==1",
	}

	trans := Transition{
		src:   &current_state,
		dest:  &next_state,
		event: &eve,
	}

	// Transition構造体を表示
	fmt.Println(trans)
}

func TestWriteEnum(t *testing.T) {
	// write_enum関数のテスト
	current_state := State{
		Name:      "Wait",
		EntryProg: "wait_test_EntryProg",
		DoProg:    "wait_test_DoProg",
		ExitProg:  "wait_test_ExitProg",
	}

	next_state := State{
		Name:      "Run",
		EntryProg: "run_test_EntryProg",
		DoProg:    "run_test_DoProg",
		ExitProg:  "run_test_ExitProg",
	}

	eve := Event{
		Name: "push_button",
		Cond: "ecrobot_get_touch_sensor(NXT_PORT_TOUCH)==1",
	}

	trans1 := Transition{
		src:   &current_state,
		dest:  &next_state,
		event: &eve,
	}

	trans2 := Transition{
		src:   &next_state,
		dest:  &current_state,
		event: &eve,
	}

	trans_list := []Transition{trans1, trans2}

	write_enum(trans_list)
}

func TestRemove(t *testing.T) {
	// remove関数のテスト
	state_wait := State{
		Name:      "Wait",
		EntryProg: "wait_test_EntryProg",
		DoProg:    "wait_test_DoProg",
		ExitProg:  "wait_test_ExitProg",
	}

	state_run := State{
		Name:      "Run",
		EntryProg: "run_test_EntryProg",
		DoProg:    "run_test_DoProg",
		ExitProg:  "run_test_ExitProg",
	}

	state_stop := State{
		Name:      "Stop",
		EntryProg: "stop_test_EntryProg",
		DoProg:    "stop_test_DoProg",
		ExitProg:  "stop_test_ExitProg",
	}

	eve1 := Event{
		Name: "push_button",
		Cond: "ecrobot_get_touch_sensor(NXT_PORT_TOUCH)==1",
	}

	eve2 := Event{
		Name: "touch_display",
		Cond: "get_touch_sensor(NXT_PORT_TOUCH)==1",
	}

	// wait → run (eve1)
	trans1 := Transition{
		src:   &state_wait,
		dest:  &state_run,
		event: &eve1,
	}

	// run → stop (eve1)
	trans2 := Transition{
		src:   &state_run,
		dest:  &state_stop,
		event: &eve1,
	}

	// stop → run (eve2)
	trans3 := Transition{
		src:   &state_stop,
		dest:  &state_run,
		event: &eve2,
	}

	trans_list := []Transition{trans1, trans2, trans3}
	fmt.Println(trans_list)

	trans_list = remove(trans_list, "push_button")
	fmt.Println(trans_list)
}

func TestWriteFunc(t *testing.T) {
	// write_func関数のテスト
	state_wait := State{
		Name:      "Wait",
		EntryProg: "wait_test_EntryProg",
		DoProg:    "wait_test_DoProg",
		ExitProg:  "wait_test_ExitProg",
	}

	state_run := State{
		Name:      "Run",
		EntryProg: "run_test_EntryProg",
		DoProg:    "run_test_DoProg",
		ExitProg:  "run_test_ExitProg",
	}

	state_stop := State{
		Name:      "Stop",
		EntryProg: "stop_test_EntryProg",
		DoProg:    "stop_test_DoProg",
		ExitProg:  "stop_test_ExitProg",
	}

	eve1 := Event{
		Name: "push_button",
		Cond: "ecrobot_get_touch_sensor(NXT_PORT_TOUCH)==1",
	}

	eve2 := Event{
		Name: "touch_display",
		Cond: "get_touch_sensor(NXT_PORT_TOUCH)==1",
	}

	// wait → run (eve1)
	trans1 := Transition{
		src:   &state_wait,
		dest:  &state_run,
		event: &eve1,
	}

	// run → stop (eve1)
	trans2 := Transition{
		src:   &state_run,
		dest:  &state_stop,
		event: &eve1,
	}

	// stop → run (eve2)
	trans3 := Transition{
		src:   &state_stop,
		dest:  &state_run,
		event: &eve2,
	}

	trans_list := []Transition{trans1, trans2, trans3}

	write_func(trans_list)
}

func TestWrite(t *testing.T) {
	// write_enum, write_func関数のテスト
	state_wait := State{
		Name:      "Wait",
		EntryProg: "wait_test_EntryProg",
		DoProg:    "wait_test_DoProg",
		ExitProg:  "wait_test_ExitProg",
	}

	state_run := State{
		Name:      "Run",
		EntryProg: "run_test_EntryProg",
		DoProg:    "run_test_DoProg",
		ExitProg:  "run_test_ExitProg",
	}

	state_stop := State{
		Name:      "Stop",
		EntryProg: "stop_test_EntryProg",
		DoProg:    "stop_test_DoProg",
		ExitProg:  "stop_test_ExitProg",
	}

	eve1 := Event{
		Name: "push_button",
		Cond: "ecrobot_get_touch_sensor(NXT_PORT_TOUCH)==1",
	}

	eve2 := Event{
		Name: "touch_display",
		Cond: "get_touch_sensor(NXT_PORT_TOUCH)==1",
	}

	// wait → run (eve1)
	trans1 := Transition{
		src:   &state_wait,
		dest:  &state_run,
		event: &eve1,
	}

	// run → stop (eve1)
	trans2 := Transition{
		src:   &state_run,
		dest:  &state_stop,
		event: &eve1,
	}

	// stop → run (eve2)
	trans3 := Transition{
		src:   &state_stop,
		dest:  &state_run,
		event: &eve2,
	}

	trans_list := []Transition{trans1, trans2, trans3}

	write_package()
	write_enum(trans_list)
	write_func(trans_list)
}
*/

func TestWriteStopWatch(t *testing.T) {
	// jsonファイル読み込みのテスト

	// state.json読み込み
	states, err := ioutil.ReadFile("./stop_watch/stop_watch.json")
	if err != nil {
		panic(err.Error())
	}

	var sm StateMent
	err = json.Unmarshal(states, &sm)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("------------------------------　output.go　-----------------------------------")
	write_package()
	write_enum(sm.States)
	write_event(sm.Transitions)
	write_init(sm.Initial)
	fmt.Println("------------------------------　output_edit.go　-----------------------------------")
	write_package_edit()
	write_func(sm.States, sm.Events)
	write_main()
}

/*
func TestWriteHeater(t *testing.T) {
	// jsonファイル読み込みのテスト

	// state.json読み込み
	states, err := ioutil.ReadFile("./heater/heater.json")
	if err != nil {
		panic(err.Error())
	}

	var sm StateMent
	err = json.Unmarshal(states, &sm)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("------------------------------　output.go　-----------------------------------")
	write_package()
	write_enum(sm.States)
	write_event(sm.Transitions)
	write_init(sm.Initial)
	fmt.Println("------------------------------　output_edit.go　-----------------------------------")
	write_package_func()
	write_func(sm.States, sm.Events)
	write_scan()
	write_main()
}
*/
