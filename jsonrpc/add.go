package jsonrpc

import (
	"github.com/AlverLyu/dna-ipfs/common"
	"github.com/AlverLyu/dna-ipfs/ipfs"
	log4 "github.com/alecthomas/log4go"
)

func init() {
	IpfsHandlerMgr.RegHandler(NewAddFile())
}

type AddFileResponse struct {
	ID string `json:"id"`
}

type AddFile struct{}

func NewAddFile() *AddFile { return &AddFile{} }

func (this *AddFile) GetName() string {
	return "addfile"
}

func (this *AddFile) Handle(params map[string]interface{}) (result interface{}, errorCode int) {
	name, nameOk := params["name"]
	data, dataOK := params["data"]
	if !nameOk || !dataOK {
		return nil, common.Err_Params
	}

	id, err := ipfs.Add([]byte(data.(string)))
	if err != nil {
		log4.Error("AddFile failed: %s", err)
		return nil, common.Err_IPFS_ERROR
	}

	v, ok := name.(string)
	if !ok {
		log4.Error("AddFile failed! 'name' is not a string")
		return nil, common.Err_Params
	}
	log4.Info("AddFile %s %s", id, v)

	return &AddFileResponse{ID: id}, common.Err_OK
}
