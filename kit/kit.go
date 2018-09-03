package kit

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type config struct {
	Mongodb string `json:"mongodb"`
	Port    uint   `json:"port"`
	Token   string `json:"token"`
}

var Config config
var Logger *log.Logger

func init() {
	dir, _ := os.Getwd()
	path := filepath.Join(dir, "config.json")
	cfile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer cfile.Close()
	dat, err := ioutil.ReadAll(cfile)
	if err != nil {
		log.Fatal(err)
	}
	if err = json.Unmarshal(dat, &Config); err != nil {
		log.Fatal(err)
	}

	logfile, err := os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	Logger = log.New(io.MultiWriter(logfile, os.Stdout), "|> ", log.Llongfile|log.Ldate|log.Ltime)
}
