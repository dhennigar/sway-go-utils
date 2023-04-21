// swayctl

// TODO: add function that brings target to current screen if it is open elsewhere (selectable by flag)
// TODO: add functions for switching to / sending to the next open workspace (selectable by flag)

package main

import(
	"os"
	"strings"
	"log"

	"go.i3wm.org/i3/v4"
)

func focus_or_launch(tree i3.Tree, appid string) {
	child := tree.Root.FindChild(func(n *i3.Node) bool {
		return strings.HasSuffix(n.AppID, appid) || strings.HasSuffix(n.WindowProperties.Instance, appid)
	})

	if child != nil {
		i3.RunCommand("[app_id=" + appid + "] focus")
		i3.RunCommand("[instance=" + appid + "] focus")
	} else {
		i3.RunCommand("exec " + appid)
	}
}

func main() {
	tree, err := i3.GetTree()
	if err != nil {
		log.Fatal(err)
	}
	
	command := os.Args[1]

	if command == "focus" {
		target := os.Args[2]
		focus_or_launch(tree, target)		
	}

	if command == "free" {
		//		focus_free_ws()
	}

}
