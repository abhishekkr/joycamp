package proc

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/abhishekkr/gol/golbin"
	"github.com/abhishekkr/gol/golcrypt"
	"github.com/abhishekkr/gol/golenv"
	"github.com/abhishekkr/gol/golhttpclient"
)

var (
	CmdDir = golenv.OverrideIfEnv("JOYCAMP_CMD_DIR", "/tmp")
)

type Proc struct {
	Cmd string
	Src string
}

func init() {
	golhttpclient.SkipSSLVerify = false
	os.MkdirAll(CmdDir, 0755)
}

func downloadCmd(src string) (cmdPath string, err error) {
	cmdPath = path.Join(CmdDir, golcrypt.MD5([]byte(src)))
	if _, err = os.Stat(cmdPath); err == nil {
		return
	}
	req := golhttpclient.HTTPRequest{Url: src}
	body, err := req.Get()
	if err != nil {
		return
	}
	err = ioutil.WriteFile(cmdPath, []byte(body), 0755)
	return
}

func (p *Proc) Run() (err error) {
	if p.Cmd == "" {
		p.Cmd, err = downloadCmd(p.Src)
		if err != nil {
			return
		}
	}
	out, err := golbin.Exec(p.Cmd)
	log.Println(out)
	return err
}
