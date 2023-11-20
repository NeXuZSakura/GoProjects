package main

//go:generate go build -tags add

var Animals = []string{
	"cat",
	"dog",
	"hedgehog",
}

func main() {
	for _, anim := range Animals {
		println(">", anim)
	}
}
