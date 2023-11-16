package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"

	"task3_1/calculations"
)

func main() {
	var err error
	var n int64
	args := os.Args
	switch len(args) {
	case 2:
		n, err = strconv.ParseInt(args[1], 10, 64)
	case 3:
		n, err = strconv.ParseInt(args[2], 10, 64)
	default:
		log.Fatal("Error: Use only two values")
	}
	if err != nil {
		log.Fatal("Error: Incorrect input")
	}
	myflag := flag.Bool("log", false, "log")
	flag.Parse()
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.Info(calculations.Calculate(n, *myflag))
}
