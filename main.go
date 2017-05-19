package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/AlverLyu/dna-ipfs/conf"
	"github.com/AlverLyu/dna-ipfs/ipfs"
	"github.com/AlverLyu/dna-ipfs/jsonrpc"
	log4 "github.com/alecthomas/log4go"
)

var CfgFile string
var LogFile string

func init() {
	flag.StringVar(&CfgFile, "cf", "./etc/dnaipfs.json", "The path of config file")
	flag.StringVar(&LogFile, "lf", "./etc/log4go.xml", "The path of log config file")
}

func main() {
	log4.LoadConfiguration(LogFile)

	conf.GCfg.Init(CfgFile)
	ipfs.SetIPFSURL(conf.GCfg.IPFSURL)

	rpcServer := jsonrpc.NewJsonRpcServer(conf.GCfg.Port, conf.GCfg.RpcPath, conf.GCfg.ReadTimeout, conf.GCfg.WriteTimeout, conf.GCfg.MaxHeaderBytes)
	rpcServer.Start()

	waitToExit()
}

func waitToExit() {
	exit := make(chan bool, 0)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		<-sc
		close(exit)

	}()
	<-exit
}
