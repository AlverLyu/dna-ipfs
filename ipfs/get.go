package ipfs

import (
	"io/ioutil"
)

// Get the data with the given @id
func Get(id string) ([]byte, error) {
	s := newIPFSShell()

	r, err := s.Cat(id)
	defer r.Close()

	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
