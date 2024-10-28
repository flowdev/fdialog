package main

import (
	"github.com/flowdev/fdialog/cmd"
	"log"
)

func main() {
	log.Default().SetFlags(0) // use simple logger without date, time, file, ...
	cmd.Execute()
}
