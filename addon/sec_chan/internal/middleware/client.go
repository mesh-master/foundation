package middleware

import (
	"github.com/mesh-master/foundation/addon/sec_chan/internal/codec"
	"github.com/mesh-master/foundation/addon/sec_chan/x"
	"github.com/mesh-master/foundation/internal/autogen/foundation"
	"github.com/mesh-master/foundation/internal/autogen/net/sec_chan"
	"github.com/mesh-master/foundation/pkg/z"
	"github.com/mesh-master/foundation/pkg/z/ancillary/crypto"
	"github.com/mesh-master/foundation/pkg/z/dictionary"
	"google.golang.org/protobuf/proto"
)

func attachApiKey(client z.NetworkClientInterface, req z.RequestInterface, cipher crypto.AEAD_CipherInterface) {
	meta := req.Meta().Dictionary().(dictionary.NetRequestInterface)
	if v, has := req.ServiceReflection().Get(foundation.E_AuthType); has {
		if v.(foundation.AuthType) == foundation.AuthType_ApiKey {
			rawKey := client.ApiKey()
			if cipher != nil {
				encKey := cipher.Encrypt(rawKey, nil)
				meta.SetApiKey(encKey)
			} else {
				meta.SetApiKey(rawKey)
			}
		}
	}
}

func ClientReqHandler(next z.NextHandlerFn, ctx z.NetContextInterface, req z.RequestInterface) (err error) {
	var (
		msg proto.Message
		df  x.DataFrameInterface
	)
	client := ctx.(z.NetClientContextInterface).Client()
	hasTls := client.Endpoint().(z.NetEndpointInterface).TlsConfig() != nil
	if !hasTls {
		msg = req.Data().(proto.Message)
		df = codec.NewDataFrame(msg)
		encOff := req.MethodReflection().Bool(sec_chan.E_EncOff)
		if !encOff {
			cipher := ctx.(z.NetClientContextInterface).Client().BlockCipher()
			df.WithBlockCipher(cipher)
			attachApiKey(client, req, cipher)
		}
		req.WithData(df)
	} else {
		attachApiKey(client, req, nil)
	}
	_, err = next(req, nil)
	return
}

func ClientResHandler(next z.NextHandlerFn, ctx z.NetContextInterface, res z.ResponseInterface) (err error) {
	var (
		df x.DataFrameInterface
	)
	client := ctx.(z.NetClientContextInterface).Client()
	hasTls := client.Endpoint().(z.NetEndpointInterface).TlsConfig() != nil
	if !hasTls {
		encOff := res.MethodReflection().Bool(sec_chan.E_EncOff)
		if !encOff {
			cipher := ctx.(z.NetClientContextInterface).Client().BlockCipher()
			df = res.Data().(x.DataFrameInterface)
			df.WithBlockCipher(cipher)
			if err = df.Decrypt(); err != nil {
				return
			}
		}
	}
	_, err = next(nil, res)
	return
}
