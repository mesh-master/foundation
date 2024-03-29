//
// Implementation of NetParcel network service along with its runtime registration.

package server

import (
	"github.com/mesh-master/foundation/app/net_parcel/server/ftp"
	"github.com/mesh-master/foundation/internal/ancillary/archive"
	proto "github.com/mesh-master/foundation/internal/autogen/foundation"
	"github.com/mesh-master/foundation/internal/runtime"
	"github.com/mesh-master/foundation/pkg/y/netparcel"
	"github.com/mesh-master/foundation/pkg/z"
	"github.com/mesh-master/foundation/pkg/z/ancillary"
	"github.com/mesh-master/foundation/pkg/z/dictionary"
	"github.com/mesh-master/foundation/pkg/z/platform"
	"google.golang.org/grpc"
)

var Name = proto.NetParcel_ServiceDesc.ServiceName

type serviceUnimpl struct {
	proto.UnimplementedNetParcelServer
}

type netParcel struct {
	z.NetworkServiceInterface
	ancillary.ArchiveOptions
	ftp.FtpImpl
	sessionImpl
	serviceUnimpl
}

type ConfigInterface interface {
	z.ServiceCfgInterface
}

func (s *netParcel) Register(srv *grpc.Server) {
	proto.RegisterNetParcelServer(srv, s)
}

func (s *netParcel) RegisterFtpPostActionHandler(fn netparcel.FtpPostActionHandlerFn, fileExt string) {
	s.FtpImpl.PostActions[fileExt] = fn
}

func (svc *netParcel) handleGzipTarball(ctx dictionary.NetServerContextInterface, path platform.Pathname) (err error) {
	var (
		plat  z.PlatformInterface
		untar z.RunnableInterface
	)
	if ctx.Tenant() != 0 {
		// TODO: retrieve a tenant platform API object from the runtime registry
	} else {
		plat = runtime.Runtime().Platform()
	}
	if untar, err = archive.NewUntar(plat, path, ancillary.GzipCompressor, svc.ArchiveOptions); err != nil {
		return
	}
	if err = untar.Run(); err != nil {
		return
	}
	// Probably there won't be a case that will require to keep the uploaded tarball once it's unpacked.
	// We can remove it.
	err = plat.Remove(path)
	return
}
