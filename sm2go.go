package parser

import (
	"fmt"
)

/*
・1つ1つのステートを変換
→ ステートの中にステートの構造を変換
→ XML形式を変換
*/

type StateMent struct {
	States      []State      `json:"State"`
	Events      []Event      `json:"Event"`
	Transitions []Transition `json:"Transition"`
}

type State struct {
	Name string `json:"name"`
}

type Event struct {
	Name string `json:"name"` // イベント名
}

type Transition struct {
	Src   string `json:"src"`   // 現在のステート
	Dest  string `json:"dest"`  // 遷移先
	Event string `json:"event"` // イベント（矢印）
}

// var transition_list []Transition

func remove(transition_list []Transition, event_name string) []Transition {
	ret := make([]Transition, len(transition_list))
	i := 0

	for _, transition := range transition_list {
		if event_name != transition.Event {
			ret[i] = transition
			i++
		}
	}

	return ret[:i]
}

func write_package() {
	fmt.Println("// Please run \"gofmt -w /path/to/output.go\"")
	fmt.Println("// go run *.go")
	fmt.Println("")

	fmt.Println("package // write package name")
	fmt.Println("")

	fmt.Println("import (")
	fmt.Println("// write about import")
	fmt.Println(")")
	fmt.Println("")
}

func write_enum(state_list []State) {
	// enum宣言
	fmt.Println("type State int")
	fmt.Println("const (")
	for index, state := range state_list {
		if index == 0 {
			fmt.Println(state.Name, "State = iota")
			continue
		}
		fmt.Println(state.Name)
	}
	fmt.Println(")")
	fmt.Println("")

	fmt.Println("type Eod int")
	fmt.Println("const (")
	fmt.Println("Entry Eod = iota\nDo\nExit\n)")
	fmt.Println("")
}

func write_event(transition_list []Transition) {
	// 状態ごとの関数を作成
	for len(transition_list) != 0 {
		now_event := transition_list[0].Event
		fmt.Printf("func %s (state State, eod Eod) {\n", now_event)
		fmt.Println("switch state {")

		for _, transition := range transition_list {
			if now_event == transition.Event {
				fmt.Printf("case %s:\n", transition.Src)
				// Entry状態での動作を記述
				fmt.Println("if eod == Entry {")
				fmt.Printf("%s_Entry()\n", transition.Src)
				fmt.Println("eod = Do")
				fmt.Println("}") // if eod == Entry

				// Do状態での動作を記述
				fmt.Println("if eod == Do {")
				fmt.Printf("%s_Do()\n", transition.Src)
				fmt.Printf("if %s_Cond() {\n", transition.Event)
				fmt.Printf("state = %s\n", transition.Dest)
				fmt.Printf("fmt.Println(\"State is changed: %s to %s\")\n", transition.Src, transition.Dest)
				fmt.Println("eod = Entry")
				fmt.Println("}") // if event_Cond()
				fmt.Println("}") // if eod == Do
			}
			transition_list = remove(transition_list, now_event) // 表示したtransitionをリストから削除
		}
		fmt.Println("}") // switch state
		fmt.Println("}") // func

		fmt.Println("")
	}
}

func write_func(state_list []State, event_list []Event) {
	write_package()

	for _, state := range state_list {
		fmt.Printf("func %s_Entry() {\n", state.Name)
		fmt.Printf("// Processing performed in the \"Entry\" state of the \"%s\" state.\n", state.Name)
		fmt.Println("}")
		fmt.Println("")
	}

	for _, state := range state_list {
		fmt.Printf("func %s_Do() {\n", state.Name)
		fmt.Printf("// Processing performed in the \"Do\" state of the \"%s\" state.\n", state.Name)
		fmt.Println("}")
		fmt.Println("")
	}

	for _, event := range event_list {
		fmt.Printf("func %s_Cond() bool {\n", event.Name)
		fmt.Printf("// Please write the conditions under which a state transitions\n")
		fmt.Println("return true")
		fmt.Println("}")
		fmt.Println("")
	}

}

func write_main(state_list []State) {
	fmt.Println("/*")

	fmt.Println("func main() {")
	for _, state := range state_list {
		fmt.Printf("var %s State = %s\n", state.Name, state.Name)
	}
	fmt.Println("var eod Eod = Entry")
	fmt.Println("}")
	fmt.Println("*/")
}
