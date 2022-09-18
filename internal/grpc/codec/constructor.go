package codec

import (
	"github.com/go-serv/foundation/internal/ancillary/net"
	"google.golang.org/protobuf/proto"
)

func NewDataFrame(msg proto.Message) *dataFrame {
	df := new(dataFrame)
	df.Message = msg
	df.netw = net.NewWriter()
	return df
}

func NewCodec(name string) *codec {
	c := &codec{}
	c.name = "proto"
	return c
}
