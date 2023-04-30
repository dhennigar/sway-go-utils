package main

import (
	"flag"
	"log"
	"strconv"

	"go.i3wm.org/i3/v4"
)

func contains(s []int, i int) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}
	return false
}

func find_new_workspace() int {

	workspaces, err := i3.GetWorkspaces()
	if err != nil {
		log.Fatal(err)
	}

	ws_numbers := make([]int, 0)

	for i := 0; i < len(workspaces); i++ {
		ws_numbers = append(ws_numbers, int(workspaces[i].Num))
	}

	for i := 1; i < 9; i++ {
		if !contains(ws_numbers, i) {
			return i
		}
	}
	return 0
}

func main() {

	var moveFlag = flag.Bool("m", false, "move the focused view to new ws")
	flag.Parse()

	new_ws := strconv.Itoa(find_new_workspace())

	if *moveFlag {
		i3.RunCommand("move to workspace " + new_ws)
	}

	i3.RunCommand("workspace " + new_ws)
	
}
