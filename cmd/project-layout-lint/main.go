package main

import (
	"flag"
	"fmt"

	log "github.com/sirupsen/logrus"
)

const version = "v0.1.0-rc.3"

func main() {
	showVersion := flag.Bool("version", false, "Prints version information")
	flag.Parse()
	if *showVersion {
		fmt.Printf("project-layout-lint version: %s\n", version)
		return
	}

	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
	fmt.Println("helloworld")
}
