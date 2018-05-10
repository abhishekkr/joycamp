package proc

import (
	"encoding/json"
	"flag"
	"fmt"
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

	Cfg  = flag.String("cfg", "", "json config to map all")
	Cmd  = flag.String("cmd", "", "command to be run provided directly")
	Src  = flag.String("src", "", "source to fetch command")
	Args = flag.String("args", "", "args to be passed to command")
	Env  = flag.String("env", "{\"mypath\": \"/tmp\", \"yourpath\": \"/mnt\" }", "env to be set for command")
)

type EnvMap map[string]string

type Proc struct {
	Cmd    string
	Src    string
	Args   string
	EnvMap EnvMap
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
	out, err := golbin.ExecWithEnv(fmt.Sprintf("%s %s", p.Cmd, p.Args), p.EnvMap)
	log.Println(out)
	return err
}

func cfgManager() *Proc {
	cfgBytes, err := ioutil.ReadFile(*Cfg)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("joycamp~", string(cfgBytes))

	var proc Proc
	if err = json.Unmarshal(cfgBytes, &proc); err != nil {
		log.Fatalln("corrupted cfg value")
	}
	return &proc
}

func argManager() *Proc {
	fmt.Println("joycamp~", *Cmd)

	var env EnvMap
	if err := json.Unmarshal([]byte(*Env), &env); err != nil {
		log.Fatalln("corrupted env value")
	}

	return &Proc{
		Cmd:    *Cmd,
		Src:    *Src,
		Args:   *Args,
		EnvMap: env,
	}
}

func Manager() *Proc {
	flag.Parse()
	if *Cfg == "" {
		return argManager()
	}
	return cfgManager()
}
