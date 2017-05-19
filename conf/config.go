package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	log4 "github.com/alecthomas/log4go"
)

var GCfg = NewConfig()

type Config struct {
	Port           string
	RpcPath        string
	ReadTimeout    int
	WriteTimeout   int
	MaxHeaderBytes int
	IPFSURL        string
}

func NewConfig() *Config {
	return &Config{}
}

func (this *Config) Init(file string) {
	err := this.load(file)
	if err == nil {
		return
	}
	log4.Error("Config Init error:%s", err)
	this.Port = "8080"
	this.RpcPath = "/rpc/ipfs"
	this.ReadTimeout = 30
	this.WriteTimeout = 30
	this.MaxHeaderBytes = 104857600
	this.IPFSURL = "http://localhost:5001"
}

func (this *Config) load(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("NewConfig open file:%s error:%s", file, err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return fmt.Errorf("NewConfig read file:%s error:%s", file, err)
	}
	err = json.Unmarshal(data, this)
	if err != nil {
		return fmt.Errorf("json Unmarshal data:s error:%s", data, err)
	}
	return nil
}
