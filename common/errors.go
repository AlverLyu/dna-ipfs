package common

const (
	Err_OK               = 0
	Err_File_Not_EXIST   = 1001
	Err_Params           = 1002
	Err_Method_Not_EXIST = 1003
	Err_IPFS_ERROR       = 1004
	Err_Unknow           = 9999
)

var errDesc = map[int]string{
	Err_OK:               "OK",
	Err_File_Not_EXIST:   "File does not exist",
	Err_Params:           "Invalid parameters",
	Err_Method_Not_EXIST: "Method does not exist",
	Err_IPFS_ERROR:       "IPFS error",
	Err_Unknow:           "Unknow error",
}

func GetErrDesc(errorCode int) string {
	desc, _ := errDesc[errorCode]
	return desc
}
