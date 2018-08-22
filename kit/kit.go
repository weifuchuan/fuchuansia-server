package kit

import (
	"os"
	"path/filepath"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"encoding/json"
)

type config struct {
	Mongodb string `json:"mongodb"`
}

var Config config

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
}
