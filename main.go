package main

import (
	"fmt"
	"log"

	"github.com/mattn/go-shellwords"
)

func main() {
	cmd := "$(sleep 1; sleep 1)"
	fmt.Println("Commandline:", cmd)
	args, err := shellwords.Parse(cmd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Args:", args)
	for _, arg := range args {
		fmt.Println(arg)
	}
}
