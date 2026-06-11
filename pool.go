package xml2json

import (
	"bytes"
	"sync"
)

var bp = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func getBuffer() *bytes.Buffer {
	buf := bp.Get().(*bytes.Buffer)
	buf.Reset()

	return buf
}

func putBuffer(buf *bytes.Buffer) {
	bp.Put(buf)
}
