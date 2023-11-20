package main

//go:generate go build -tags add

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
