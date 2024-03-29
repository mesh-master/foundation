package z

import (
	job "github.com/AgentCoop/go-work"
	"github.com/mesh-master/foundation/pkg/z/ancillary/crypto"
	"google.golang.org/grpc"
)

type ClientInterface interface {
	Middleware() ClientMiddlewareInterface
	ServiceName() string
	Endpoint() EndpointInterface
	ConnectTask(j job.JobInterface) (job.Init, job.Run, job.Finalize)
	OnConnect(cc grpc.ClientConnInterface)
	WithDialOption(grpc.DialOption)
	DialOptions() []grpc.DialOption
	Meta() MetaInterface
	WithMeta(MetaInterface)
	// WithOptions sets options for a call.
	WithOptions(any)
	SessionId() SessionId
	Reset()
}

type NetworkClientInterface interface {
	ClientInterface
	ApiKeyAwareInterface
	NetService() NetworkServiceInterface
	BlockCipher() crypto.AEAD_CipherInterface
	WithBlockCipher(crypto.AEAD_CipherInterface)
}

type LocalClientInterface interface {
	ClientInterface
}
