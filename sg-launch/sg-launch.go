// sg-launch.go

package main

import(
	"strings"
	"log"
	"flag"

	"go.i3wm.org/i3/v4"
)

func bring_to_current(appid string) {

	tree, err := i3.GetTree()
	if err != nil {
		log.Fatal(err)
	}

	current_ws := tree.Root.FindFocused(func(n *i3.Node) bool {
		return n.Type == i3.WorkspaceNode
	})

	i3.RunCommand("[app_id=" + appid + "] move to workspace " + current_ws.Name)
	i3.RunCommand("[instance=" + appid + "] move to workspace " + current_ws.Name)
}

func focus_or_launch(appid string) {

	tree, err := i3.GetTree()
	if err != nil {
		log.Fatal(err)
	}

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

	//	var sendFlag = flag.Bool("s", false, "send focused window to target")
	var bringFlag = flag.Bool("b", false, "bring target to current workspace")
	flag.Parse()
	
	target := flag.Arg(0)
	
	if *bringFlag {
		bring_to_current(target)
	}
	focus_or_launch(target)		
}
