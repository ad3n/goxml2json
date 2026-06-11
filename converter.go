package xml2json

import (
	"bytes"
	"io"
)

func Convert(r io.Reader, ps ...plugin) (*bytes.Buffer, error) {
	root := &Node{}
	err := NewDecoder(r, ps...).Decode(root)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	e := NewEncoder(buf, ps...)
	err = e.Encode(root)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
