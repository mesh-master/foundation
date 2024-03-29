package archive

import (
	"archive/tar"
	"bytes"
	"github.com/mesh-master/foundation/pkg/z/platform"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type ttar struct {
	archive
	tarReader *tar.Reader
	tarWriter *tar.Writer
	basedir   platform.Pathname
}

func (t *ttar) handleHeader(hdr *tar.Header) (err error) {
	pathname := t.basedir.ComposePath(hdr.Name).Normalize()
	switch hdr.Typeflag {
	case tar.TypeDir:
		err = t.fs.CreateDir(pathname, platform.UnixPerms(hdr.Mode))
	case tar.TypeReg:
		var fd *os.File
		if fd, err = os.OpenFile(pathname.String(), os.O_CREATE|os.O_WRONLY, os.FileMode(hdr.Mode)); err != nil {
			return
		}
		var buf bytes.Buffer
		if _, err = io.Copy(&buf, t.tarReader); err != nil {
			return
		}
		err = t.fs.Write(fd, buf.Bytes())
	}
	return
}

func (t *ttar) handleRegularFile(path string, info fs.FileInfo) (err error) {
	var (
		file []byte
	)
	hdr := &tar.Header{
		Name: path,
		Size: info.Size(),
		Mode: int64(t.fsPerms),
	}
	if err = t.tarWriter.WriteHeader(hdr); err != nil {
		return
	}
	if file, err = os.ReadFile(path); err != nil {
		return
	}
	if _, err = t.tarWriter.Write(file); err != nil {
		return
	}
	return
}

func (t *ttar) walkFn(path string, info fs.FileInfo, err error) error {
	if err != nil {
		return err
	}
	switch m := info.Mode(); true {
	case m.IsDir():
		return filepath.Walk(path, t.walkFn)
	case m.IsRegular():
		return t.handleRegularFile(path, info)
	}
	return err
}

func (t *ttar) Run() (err error) {
	switch true {
	case t.tarWriter != nil:
		err = filepath.Walk(t.target.String(), t.walkFn)
	case t.tarReader != nil:
		var hdr *tar.Header
		t.basedir = t.target.Dirname()
		for {
			hdr, err = t.tarReader.Next()
			switch err {
			case io.EOF:
				return nil
			case nil:
				if err = t.handleHeader(hdr); err != nil {
					return
				}
			default:
				return err
			}
		}
	}
	return
}
