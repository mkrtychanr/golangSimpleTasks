package main

import "fmt"

type Human struct {
	Name    string
	Surname string
}

type Action struct {
	Human
	ActionType string
}

func main() {
	ac := Action{Human{"John", "Doe"}, "Jump"}
	fmt.Println(ac.Name)
}
