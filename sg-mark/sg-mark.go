package main

import(
	"flag"
	"log"

	"go.i3wm.org/i3/v4"
)

func main() {
	bringFlag := flag.Bool("b", false, "bring marked view to current ws")
	newFlag := flag.Bool("n", false, "mark a new window")
	flag.Parse()

	tree, err := i3.GetTree()
	if err != nil {
		log.Fatal(err)
	}


	marked_window := tree.Root.FindChild(func(n *i3.Node) bool {
		if len(n.Marks) > 0 {
			for i := 0; i <= len(n.Marks); i++ {
				return n.Marks[i] == "mark"
			}
		}
		return false
	})

	if marked_window != nil && !*newFlag {
		if *bringFlag {
			current_ws := tree.Root.FindFocused(func(n *i3.Node) bool {
				return n.Type == i3.WorkspaceNode
			})
			i3.RunCommand("[con_mark=mark] move to workspace " + current_ws.Name)
		}
		i3.RunCommand("[con_mark=mark] focus")	
	} else {
		i3.RunCommand("[con_id=__focused__] mark --add 'mark'")
	}	
}
