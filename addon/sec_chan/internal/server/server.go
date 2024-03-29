//
// Implementation of NetParcel network service along with its runtime registration.

package server

import (
	proto "github.com/mesh-master/foundation/internal/autogen/net/sec_chan"
	"github.com/mesh-master/foundation/pkg/z"
	"github.com/mesh-master/foundation/pkg/z/ancillary"
	"google.golang.org/grpc"
)

var Name = proto.SecureChannel_ServiceDesc.ServiceName

type serviceUnimpl struct {
	proto.UnimplementedSecureChannelServer
}

type impl struct{}

type secChanServer struct {
	z.NetworkServiceInterface
	ancillary.ArchiveOptions
	impl
	serviceUnimpl
}

type ConfigInterface interface {
	z.ServiceCfgInterface
}

func (s *secChanServer) Register(srv *grpc.Server) {
	proto.RegisterSecureChannelServer(srv, s)
}
