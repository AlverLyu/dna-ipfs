package common

import "encoding/json"

const DefaultIPFSRpcVersion = "2.0"

type IPFSRpcRequest struct {
	Id     string                 `json:"id"`
	Method string                 `json:"method"`
	Params map[string]interface{} `json:"params"`
}

func NewIPFSRpcRequest(data []byte) (*IPFSRpcRequest, error) {
	req := &IPFSRpcRequest{}
	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

type IPFSRpcResponse struct {
	Id          string      `json:"id"`
	Jsonrpc     string      `json:"jsonrpc"`
	ErrorCode   int         `json:"errrorcode"`
	ErrorString string      `json:"errorstring"`
	Result      interface{} `json:"result"`
}

func NewIPFSRpcResponse(id string, result interface{}, errorCode int, jsonrpcVersion ...string) *IPFSRpcResponse {
	jsonrpc := DefaultIPFSRpcVersion
	if len(jsonrpcVersion) > 0 {
		jsonrpc = jsonrpcVersion[0]
	}
	return &IPFSRpcResponse{
		Id:        id,
		Jsonrpc:   jsonrpc,
		Result:    result,
		ErrorCode: errorCode,
	}
}

func (this *IPFSRpcResponse) Marshal() ([]byte, error) {
	this.ErrorString = GetErrDesc(this.ErrorCode)
	return json.Marshal(this)
}
