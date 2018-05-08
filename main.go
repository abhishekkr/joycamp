package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/abhishekkr/gol/golbin"
)

var (
	Cmd = flag.String("cmd", "", "command to be run provided directly")
)

type Proc struct {
	Command string
}

func main() {
	flag.Parse()
	fmt.Println("joycamp~", *Cmd)
	out, err := golbin.Exec(*Cmd)
	log.Println(out)
	log.Println(err)
}
