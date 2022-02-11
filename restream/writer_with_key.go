package restream

import (
	"github.com/gwuhaolin/livego/av"
	"github.com/gwuhaolin/livego/protocol/rtmp"
)

type WriterWithKey struct {
	inner *rtmp.VirWriter
	key   string
}

func NewWriterWithKey(inner *rtmp.VirWriter, key string) av.WriteCloser {
	return &WriterWithKey{
		inner: inner,
		key:   key,
	}
}

func (v *WriterWithKey) Info() av.Info {
	info := v.inner.Info()
	info.Key = v.key
	return info
}

func (v *WriterWithKey) Close(err error) {
	v.inner.Close(err)
}

func (v *WriterWithKey) Alive() bool {
	return v.inner.Alive()
}

func (v *WriterWithKey) CalcBaseTimestamp() {
	v.inner.CalcBaseTimestamp()
}

func (v *WriterWithKey) Write(packet *av.Packet) error {
	return v.inner.Write(packet)
}
