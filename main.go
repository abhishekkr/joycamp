package main

import (
	"flag"
	"fmt"
	"log"

	proc "github.com/abhishekkr/joycamp/proc"
)

var (
	Cmd = flag.String("cmd", "", "command to be run provided directly")
	Src = flag.String("src", "", "source to fetch command")
)

func main() {
	flag.Parse()
	fmt.Println("joycamp~", *Cmd)
	p := proc.Proc{Cmd: *Cmd, Src: *Src}
	if err := p.Run(); err != nil {
		log.Println(err)
	}
}
