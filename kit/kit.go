package kit

import (
	"os"
	"path/filepath"
	"log"
	"io/ioutil"
	"encoding/json"
)

type config struct {
	Mongodb string `json:"mongodb"`
	Port    uint   `json:"port"`
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

	logfile, err := os.Open("logs.log")
	if err != nil {
		log.Fatal(err)
	}
	Logger = log.New(logfile, "fuchuansia|> ", log.Llongfile)
}
