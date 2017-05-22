package main

import (
	"flag"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/AlverLyu/dna-ipfs/conf"
	"github.com/AlverLyu/dna-ipfs/ipfs"
	"github.com/AlverLyu/dna-ipfs/jsonrpc"
)

var CfgFile string
var LogFile string

func main() {
	binPath, err := os.Executable()
	if err != nil {
		binPath = "."
	}
	binDir := filepath.Dir(binPath)
	flag.StringVar(&CfgFile, "c", binDir+"/dnaipfs.cfg", "The path of config file")
	flag.StringVar(&LogFile, "lc", "", "The path of log config file")
	flag.Parse()

	if len(LogFile) > 0 {
		conf.OpenCustomLog(LogFile)
	} else {
		os.Mkdir(binDir+"/log", os.ModePerm)
		conf.OpenDefaultLog(binDir + "/log/dnaipfs.log")
	}

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
