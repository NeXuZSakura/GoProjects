package main

import (
	_ "embed"
	"fmt"
)

//go:generate mockgen -destination=mock_foo.go -package=generate . Doer

//go:embed file2.txt
var fileString string

var Animals = []string{
	"cat",
	"dog",
	"hedgehog",
}

func main() {
	for _, anim := range Animals {
		fmt.Println(">", anim)
	}
	fmt.Println(fileString)
}

type Doer interface {
	DoSomething(int, string) error
}
