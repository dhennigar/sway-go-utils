// sg-launch.go

package main

import (
	"flag"
	"log"

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

	i3.RunCommand("[app_id=" + appid + "] focus")
	i3.RunCommand("[instance=" + appid + "] focus")
}

func send_to_target(appid string) {

	tree, err := i3.GetTree()
	if err != nil {
		log.Fatal(err)
	}

	child := tree.Root.FindChild(func(n *i3.Node) bool {
		return n.AppID == appid || n.WindowProperties.Instance == appid
	})

	parent := child.FindParent()

	if child != nil && parent != nil {
		i3.RunCommand("[con_id=__focused__] move to workspace " + parent.Name)
		i3.RunCommand("workspace " + parent.Name)
	}
}

func focus_or_launch(appid string) {

	tree, err := i3.GetTree()
	if err != nil {
		log.Fatal(err)
	}

	child := tree.Root.FindChild(func(n *i3.Node) bool {
		return n.AppID == appid || n.WindowProperties.Instance == appid
	})

	if child != nil {
		i3.RunCommand("[app_id=" + appid + "] focus")
		i3.RunCommand("[instance=" + appid + "] focus")
	} else {
		i3.RunCommand("exec " + appid)
	}
}

func main() {

	var bringFlag = flag.Bool("b", false, "bring target to current workspace")
	var sendFlag = flag.Bool("s", false, "send focused window to target's workspace")
	flag.Parse()

	target := flag.Arg(0)

	switch {
	case *bringFlag:
		bring_to_current(target)
	case *sendFlag:
		send_to_target(target)
	default:
		focus_or_launch(target)
	}
}
