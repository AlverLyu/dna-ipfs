package jsonrpc

import (
	"github.com/AlverLyu/dna-ipfs/common"
	"github.com/AlverLyu/dna-ipfs/ipfs"
	log4 "github.com/alecthomas/log4go"
)

func init() {
	IpfsHandlerMgr.RegHandler(NewGetFile())
}

type GetFileResponse struct {
	Data []byte `json:"data"`
}

type GetFile struct{}

func NewGetFile() *GetFile { return &GetFile{} }

func (this *GetFile) GetName() string {
	return "getfile"
}

func (this *GetFile) Handle(params map[string]interface{}) (result interface{}, errorCode int) {
	p, ok := params["id"]
	if !ok {
		return nil, common.Err_Params
	}

	id, ok := p.(string)
	if !ok {
		log4.Error("GetFile failed: 'id' is not a string")
		return nil, common.Err_Params
	}

	data, err := ipfs.Get(id)
	if err != nil {
		log4.Error("GetFile failed: %s", err)
		return nil, common.Err_IPFS_ERROR
	}

	log4.Info("GetFile %s: %v", id, data)

	return &GetFileResponse{Data: data}, common.Err_OK
}
