package jsonrpc

import (
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/AlverLyu/dna-ipfs/common"
	log4 "github.com/alecthomas/log4go"
	"github.com/julienschmidt/httprouter"
)

var IpfsHandlerMgr = NewIpfsRpcHandlerMgr()

type IpfsRpcHandler interface {
	GetName() string
	Handle(parmas map[string]interface{}) (result interface{}, errorCode int)
}

type IpfsRpcHandlerMgr struct {
	handler map[string]IpfsRpcHandler
	lock    sync.RWMutex
}

func NewIpfsRpcHandlerMgr() *IpfsRpcHandlerMgr {
	return &IpfsRpcHandlerMgr{
		handler: make(map[string]IpfsRpcHandler, 0),
	}
}

func (this *IpfsRpcHandlerMgr) Handle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log4.Error("IpfsRpcHandler Handle read body error:%s", err)
		this.writeResponse(common.NewIPFSRpcResponse("", nil, common.Err_Unknow), w)
		return
	}
	req, err := common.NewIPFSRpcRequest(data)
	if err != nil {
		log4.Error("IpfsRpcHandler NewIPFSRpcRequest from:%s error:%s", data, err)
		this.writeResponse(common.NewIPFSRpcResponse("", nil, common.Err_Unknow), w)
		return
	}
	h, ok := this.GetHandler(req.Method)
	if !ok {
		this.writeResponse(common.NewIPFSRpcResponse(req.Id, nil, common.Err_Method_Not_EXIST), w)
		return
	}

	res, errCode := h.Handle(req.Params)
	this.writeResponse(common.NewIPFSRpcResponse(req.Id, res, errCode), w)
}

func (this *IpfsRpcHandlerMgr) writeResponse(rsp *common.IPFSRpcResponse, w http.ResponseWriter) {
	rspData, _ := rsp.Marshal()
	w.WriteHeader(http.StatusOK)
	w.Write(rspData)
}

func (this *IpfsRpcHandlerMgr) RegHandler(handler IpfsRpcHandler) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.handler[handler.GetName()] = handler
}

func (this *IpfsRpcHandlerMgr) GetHandler(name string) (IpfsRpcHandler, bool) {
	this.lock.RLock()
	defer this.lock.RUnlock()
	h, ok := this.handler[name]
	return h, ok
}
