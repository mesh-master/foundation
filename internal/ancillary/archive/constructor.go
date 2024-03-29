package archive

import (
	"archive/tar"
	"compress/gzip"
	"github.com/mesh-master/foundation/pkg/z/ancillary"
	"github.com/mesh-master/foundation/pkg/z/platform"
	"os"
)

func NewTar(target platform.Pathname, comp ancillary.CompressorType, opts ancillary.ArchiveOptions) (*ttar, error) {
	var (
		fd  *os.File
		err error
	)
	un := new(ttar)
	un.target = target
	un.ArchiveOptions = opts
	if fd, err = os.OpenFile(target.String(), os.O_RDWR, os.FileMode(0555)); err != nil {
		return nil, err
	}
	switch comp {
	case ancillary.GzipCompressor:
		if un.compWriter, err = gzip.NewWriterLevel(fd, un.GzipLevel); err != nil {
			return nil, err
		}
	}
	return un, nil
}

func NewUntar(fs platform.FilesystemInterface, target platform.Pathname, comp ancillary.CompressorType, opts ancillary.ArchiveOptions) (*ttar, error) {
	var (
		fd       *os.File
		err      error
		gzReader *gzip.Reader
	)
	un := new(ttar)
	un.fs = fs
	un.target = target
	un.ArchiveOptions = opts
	if fd, err = os.OpenFile(target.String(), os.O_RDONLY, os.FileMode(0444)); err != nil {
		return nil, err
	}
	switch comp {
	case ancillary.GzipCompressor:
		if gzReader, err = gzip.NewReader(fd); err != nil {
			return nil, err
		}
		gzReader.Multistream(un.GzipMultistream)
		un.tarReader = tar.NewReader(gzReader)
	}
	return un, nil
}
