package main

import (
	"github.com/flowdev/fdialog/cobracmd"
	"github.com/flowdev/fdialog/uimain"
	"log"
)

func main() {
	log.Default().SetFlags(0) // use simple logger without date, time, file, ...
	err := uimain.RegisterEverything()
	if err != nil {
		log.Printf("FATAL: %v", err)
	}
	cobracmd.Execute()
}
