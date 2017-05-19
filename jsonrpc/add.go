package jsonrpc

import (
	"encoding/json"

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
	_, nameOK := params["name"]
	_, dataOK := params["data"]
	if !nameOK || !dataOK {
		return nil, common.Err_Params
	}

	ipfsData, err := json.Marshal(params)
	id, err := ipfs.Add(ipfsData)
	if err != nil {
		log4.Error("AddFile failed: %s", err)
		return nil, common.Err_IPFS_ERROR
	}

	log4.Info("AddFile succeeded, IPFS ID is %s", id)

	return &AddFileResponse{ID: id}, common.Err_OK
}
