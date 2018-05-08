package main

import (
	"flag"
	"fmt"
	"log"

	proc "github.com/abhishekkr/joycamp/proc"
)

var (
	Cmd = flag.String("cmd", "", "command to be run provided directly")
)

func main() {
	flag.Parse()
	fmt.Println("joycamp~", *Cmd)
	p := proc.Proc{Cmd: *Cmd}
	if err := p.Run(); err != nil {
		log.Println(err)
	}
}
