package proc

import (
	"log"

	"github.com/abhishekkr/gol/golbin"
)

type Proc struct {
	Cmd string
}

func (p *Proc) Run() error {
	out, err := golbin.Exec(p.Cmd)
	log.Println(out)
	return err
}
