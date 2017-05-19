package jsonrpc

import (
	"fmt"
	"net/http"
	"time"

	log4 "github.com/alecthomas/log4go"
	"github.com/julienschmidt/httprouter"
)

type JsonRpcServer struct {
	port       string
	httpServer *http.Server
}

func NewJsonRpcServer(port, rpcPath string, readTimeout, writeTimeout int, maxHeaderBytes int) *JsonRpcServer {
	server := &JsonRpcServer{
		port: port,
	}
	httpRouter := httprouter.New()
	httpRouter.POST(rpcPath, IpfsHandlerMgr.Handle)
	httpRouter.NotFound = &RPCNotFound{}

	httpServer := &http.Server{
		Addr:           fmt.Sprintf(":%s", port),
		Handler:        httpRouter,
		ReadTimeout:    time.Second * time.Duration(readTimeout),
		WriteTimeout:   time.Second * time.Duration(writeTimeout),
		MaxHeaderBytes: maxHeaderBytes,
	}
	server.httpServer = httpServer

	return server
}

func (this *JsonRpcServer) Start() {
	log4.Info("JsonRpcServer Start.\n")
	doStart := func() {
		err := this.httpServer.ListenAndServe()
		if err != nil {
			panic(fmt.Errorf("JsonRpcServer ListenAndServe error:%s", err))
		}
	}
	go doStart()
}

type RPCNotFound struct{}

func (this *RPCNotFound) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log4.Debug("Cannot handle:%s\n", r.URL.String())
	w.WriteHeader(http.StatusNotFound)
}
