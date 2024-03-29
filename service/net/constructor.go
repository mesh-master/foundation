package net

import (
	"github.com/mesh-master/foundation/internal/service"
	"github.com/mesh-master/foundation/pkg/z"
)

func NewNetworkService(name string, cfg z.ServiceCfgInterface, endpoints []z.EndpointInterface) *netService {
	s := &netService{}
	s.ServiceInterface = service.NewService(name, cfg, endpoints)
	return s
}

func newTcpEndpoint(hostname string, port int) tcpEndpoint {
	ep := tcpEndpoint{}
	ep.EndpointInterface = service.NewEndpoint()
	ep.hostname = hostname
	ep.port = port
	return ep
}

func NewTcp4Endpoint(hostname string, port int) *tcp4Endpoint {
	e := &tcp4Endpoint{newTcpEndpoint(hostname, port)}
	return e
}

func NewTcp6Endpoint(hostname string, port int) *tcp6Endpoint {
	e := &tcp6Endpoint{newTcpEndpoint(hostname, port), nil}
	return e
}

//func NewConfig(webProxy *WebProxyConfig) *cfg {
//	cfg := new(cfg)
//	cfg.WebProxyConfig = webProxy
//	return cfg
//}
