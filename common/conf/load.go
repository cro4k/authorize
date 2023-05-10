package conf

import (
	"flag"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

var config = flag.String("config", "config.yml", "")

func init() {
	flag.Parse()
	//logrus.SetReportCaller(true)
}

func Load(v interface{}) error {
	b, err := ioutil.ReadFile(*config)
	if err != nil {
		return err
	}
	b = []byte(ExpandEnv(string(b)))
	return yaml.Unmarshal(b, v)
}

func MustLoad(v interface{}) {
	if err := Load(v); err != nil {
		logrus.Fatal(err)
	}
}
