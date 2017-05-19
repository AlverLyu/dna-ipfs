package ipfs

import (
	"bytes"
	"errors"
)

// Add @data to IPFS as a single file and return the IPFS ID
func Add(data []byte) (string, error) {
	s := newIPFSShell()
	if s == nil {
		return "", errors.New("Failed creating IPFS api shell")
	}
	return s.Add(bytes.NewBuffer(data))
}
