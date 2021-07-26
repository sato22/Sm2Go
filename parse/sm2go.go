package parser

import (
	"fmt"
)

/*
・1つ1つのステートを変換
→ ステートの中にステートの構造を変換
→ XML形式を変換
*/

type State struct {
	Name      string `json:"name"`
	EntryProg string `json:"entryProg"` // 状態ごとに何をするか
	DoProg    string `json:"doProg"`
	ExitProg  string `json:"exitProg"`
}

type Event struct {
	Name string `json:"name"` // イベント名
	Cond string `json:"cond"` // if文内の条件文（jsonファイルからとってくる）
}

type Transition struct {
	src   *State // 現在のステート？
	dest  *State // 遷移先
	event *Event // イベント（矢印）
}

var transition_list []Transition

func remove(transition_list []Transition, event_name string) []Transition {
	ret := make([]Transition, len(transition_list))
	i := 0

	for _, transition := range transition_list {
		if event_name != transition.event.Name {
			ret[i] = transition
			i++
		}
	}

	return ret[:i]
}

func write_package() {
	fmt.Println("// Please run \"gofmt -w /path/to/output.go\"")
	fmt.Println("")

	fmt.Println("package // input package name")
	fmt.Println("")

	fmt.Println("import (")
	fmt.Println("// write about import")
	fmt.Println(")")
	fmt.Println("")
}

func write_enum(transition_list []Transition) {
	// enum宣言
	fmt.Println("type State int")
	fmt.Println("const (")
	for index, transition := range transition_list {
		if index == 0 {
			fmt.Println(transition.src.Name, "State = iota")
			continue
		}
		fmt.Println(transition.src.Name)
	}
	fmt.Println(")")
	fmt.Println("")

	fmt.Println("type Eod int")
	fmt.Println("const (")
	fmt.Println("Entry Eod = iota\nDo\nExit\n)")
	fmt.Println("")
}

func write_func(transition_list []Transition) {
	// 状態ごとの関数を作成
	for len(transition_list) != 0 {
		now_event := transition_list[0].event.Name
		fmt.Println("func", now_event, "(state State, eod Eod) {")
		fmt.Println("switch state {")

		for _, transition := range transition_list {
			if now_event == transition.event.Name {
				fmt.Println("case", transition.src.Name, ":")
				// Entry状態での動作を記述
				fmt.Println("if eod == Entry {")
				fmt.Println(transition.src.EntryProg)
				fmt.Println("eod = Do")
				fmt.Println("}") // if eod == Entry

				// Do状態での動作を記述
				fmt.Println("if eod == Do {")
				fmt.Println(transition.src.DoProg)
				fmt.Println("if", transition.event.Cond, "{")
				fmt.Println("state = ", transition.dest.Name)
				fmt.Println("fmt.Println(\"State is changed:", transition.src.Name, "to", transition.dest.Name, "\")")
				fmt.Println("eod = Entry")
				fmt.Println("}") // if　transition.event.cond
				fmt.Println("}") // if eod == Do
			}
			transition_list = remove(transition_list, now_event) // 表示したtransitionをリストから削除
		}
		fmt.Println("}") // switch state
		fmt.Println("}") // func

		fmt.Println("")
	}
}
