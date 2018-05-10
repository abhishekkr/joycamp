package main

import (
	"log"

	proc "github.com/abhishekkr/joycamp/proc"
)

func main() {
	p := proc.Manager()
	if err := p.Run(); err != nil {
		log.Println(err)
	}
}
