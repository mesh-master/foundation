package net

import (
	"github.com/go-serv/foundation/pkg/z"
	"github.com/go-serv/foundation/pkg/z/ancillary/crypto"
)

type netClient struct {
	z.ClientInterface
	svc         z.NetworkServiceInterface
	blockCipher crypto.AEAD_CipherInterface
	apiKey      []byte
}

func (c *netClient) NetService() z.NetworkServiceInterface {
	return c.svc
}

func (c *netClient) BlockCipher() crypto.AEAD_CipherInterface {
	return c.blockCipher
}

func (c *netClient) WithBlockCipher(cipher crypto.AEAD_CipherInterface) {
	c.blockCipher = cipher
}

func (c *netClient) ApiKey() []byte {
	return c.apiKey
}

func (c *netClient) WithApiKey(key []byte) {
	c.apiKey = key
}