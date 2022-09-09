package main

import "fmt"

type Node struct {
	value     string
	neighbors []Node
}

func breathFirstSearch(head Node, target string) int {
	if head.neighbors == nil {
		return 0
	}
	for _, friend := range head.neighbors {
		if friend.value == target {
			fmt.Println("found: ", friend.value, target)
			return 1
		}
	}

	for _, friend := range head.neighbors {

		return 1 + breathFirstSearch(friend, target)
	}
	return 0
}

func main() {
	cn := Node{"Or", []Node{{"Yotam", nil}, {"Andrey", []Node{{"Eldar", []Node{{"Shai", nil}}}}}}}

	fmt.Println(cn)

	fmt.Printf("%d", breathFirstSearch(cn, "Shai"))
}
