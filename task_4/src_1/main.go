package main

import (
	_ "embed"
)

//go:generate go build -tags add

//go:embed files/file2.txt
var fileString string

var Animals = []string{
	"cat",
	"dog",
	"hedgehog",
}

func main() {
	for _, anim := range Animals {
		println(">", anim)
	}
	println(fileString)
}
