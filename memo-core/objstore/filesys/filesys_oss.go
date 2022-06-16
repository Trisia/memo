package filesys

import (
	"fmt"
	"github.com/tjfoc/gmsm/sm3"
	"io"
	"os"
	"path/filepath"
)

var errInvalidWrite = fmt.Errorf("invalid write result")

type FileSystemBaseOSS struct {
	base string // 基础路径
}

func NewFileSystemBaseOSS() *FileSystemBaseOSS {
	exeLoc, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	base := filepath.Join(exeLoc, "doc_assets/")
	_ = os.MkdirAll(base, os.ModePerm)
	return &FileSystemBaseOSS{base: base}
}

func (f *FileSystemBaseOSS) Put(filename string, in io.Reader) (string, int64, error) {
	var buf [4096]byte
	if in == nil {
		return "", 0, nil
	}
	dst, err := os.CreateTemp("", "doc-*.asset")
	if err != nil {
		return "", 0, err
	}
	defer os.Remove(dst.Name())
	hash := sm3.New()
	var written int64 = 0
	for {
		nr, er := in.Read(buf[:])
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			hash.Write(buf[0:nr])
			if nw < 0 || nr < nw {
				nw = 0
				if ew == nil {
					ew = errInvalidWrite
				}
			}
			written += int64(nw)
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}
	sum := hash.Sum(nil)
	// 文件名为： Hash值.后缀
	loc := fmt.Sprintf("%02x%s", sum, filepath.Ext(filename))
	// 移动文件
	err = os.Rename(filepath.Join(f.base, loc), dst.Name())
	return loc, written, err
}

func (f *FileSystemBaseOSS) Get(loc string, out io.Writer) (int64, error) {
	if out == nil || loc == "" {
		return 0, nil
	}
	src, err := os.Open(filepath.Join(f.base, loc))
	if err != nil {
		return 0, err
	}
	return io.Copy(out, src)
}

func (f *FileSystemBaseOSS) Del(loc string) error {
	if loc == "" {
		return nil
	}
	filepath.Join(f.base, loc)
	return os.Remove(loc)
}
