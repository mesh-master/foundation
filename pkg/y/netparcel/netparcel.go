package netparcel

import (
	proto "github.com/mesh-master/foundation/internal/autogen/foundation"
	"github.com/mesh-master/foundation/pkg/z"
	"github.com/mesh-master/foundation/pkg/z/dictionary"
	"github.com/mesh-master/foundation/pkg/z/platform"
)

type (
	SessionRequest         = proto.Session_Request
	SessionResponse        = proto.Session_Response
	FtpNewSessionRequest   = proto.Ftp_NewSession_Request
	FtpNewSessionResponse  = proto.Ftp_NewSession_Response
	FtpPostActionHandlerFn func(ctx dictionary.NetServerContextInterface, pathname platform.Pathname) error
)

type NetParcelServiceInterface interface {
	RegisterFtpPostActionHandler(fn FtpPostActionHandlerFn, fileExt string)
}

type NetParcelClientInterface interface {
	z.NetworkClientInterface
	Ping(payload uint64) (uint64, error)
	SecureSession(*SessionRequest) (*SessionResponse, error)
	FtpNewSession(*FtpNewSessionRequest) (*FtpNewSessionResponse, error)
	FtpTransferDir(target platform.Pathname, recursive bool, temp bool) error
	FtpTransferFile(target platform.Pathname, temp bool, postAction bool) error
}
