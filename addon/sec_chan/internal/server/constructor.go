package server

import (
	"github.com/mesh-master/foundation/pkg/z"
	"github.com/mesh-master/foundation/service/net"
)

func NewSecureChanService(eps []z.EndpointInterface, cfg ConfigInterface) z.NetworkServiceInterface {
	svc := new(secChanServer)
	svc.NetworkServiceInterface = net.NewNetworkService(Name, cfg, eps)
	//svc.WithCodec(codec.NewCodec())
	return svc
}
