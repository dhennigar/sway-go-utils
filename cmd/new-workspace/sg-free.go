package main

import(
	"log"
	"flag"
	"strconv"

	"go.i3wm.org/i3/v4"
)

func contains(s []int64, i int64) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}
	return false
}

func find_free_workspace() int {
	
	workspaces, err := i3.GetWorkspaces()
	if err != nil {
		log.Fatal(err)
	}

	ws_numbers := make([]int64, 0)
	
	for i := 0; i < len(workspaces); i++ {
		ws_numbers = append(ws_numbers, workspaces[i].Num)
	}

	for i := 1; i < 9; i++ {
		if !contains(ws_numbers, int64(i)) {
			return i
		}
	}
	return 0
}

func main() {

	var moveFlag = flag.Bool("m", false, "move the focused view to new ws")
	flag.Parse()

	free_ws := strconv.FormatInt(int64(find_free_workspace()), 10)

	if *moveFlag {
		i3.RunCommand("move to workspace " + free_ws)
	}
	
	i3.RunCommand("workspace " + free_ws)
}
