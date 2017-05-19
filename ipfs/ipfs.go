package ipfs

import (
	"github.com/ipfs/go-ipfs-api"
)

var url string

func SetIPFSURL(v string) {
	url = v
}

func GetIPFSURL() string {
	return url
}

func newIPFSShell() *shell.Shell {
	return shell.NewShell(url)
}
