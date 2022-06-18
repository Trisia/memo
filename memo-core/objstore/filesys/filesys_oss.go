package filesys

import (
	"fmt"
	"github.com/tjfoc/gmsm/sm3"
	"go.uber.org/zap"
	"io"
	"os"
	"path/filepath"
	"strings"
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

func (f *FileSystemBaseOSS) Put(docId, filename string, in io.Reader) (hashName string, written int64, err error) {
	var buf [4096]byte
	if in == nil {
		return "", 0, nil
	}
	_ = os.Mkdir(filepath.Join(f.base, docId), os.ModePerm)

	dst, err := os.CreateTemp("", "doc-*.asset")
	if err != nil {
		return "", 0, err
	}
	defer os.Remove(dst.Name())

	hash := sm3.New()
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
	if err != nil {
		return "", 0, err
	}
	err = dst.Close()
	if err != nil {
		return "", 0, err
	}

	sum := hash.Sum(nil)
	// 文件名为： Hash值.后缀
	hashName = fmt.Sprintf("%02x%s", sum, strings.ToLower(filepath.Ext(filename)))

	tmp, err := os.Open(dst.Name())
	if err != nil {
		return "", 0, err
	}
	defer tmp.Close()
	target, err := os.Create(filepath.Join(f.base, docId, hashName))
	if err != nil {
		return "", 0, err
	}
	defer target.Close()

	written, err = io.Copy(target, tmp)

	return
}

func (f *FileSystemBaseOSS) Get(docId, hashName string, out io.Writer) (written int64, err error) {
	if out == nil || hashName == "" {
		return 0, nil
	}
	p := filepath.Join(f.base, docId, hashName)
	if _, err = os.Stat(p); err != nil {
		return 0, err
	}

	src, err := os.Open(p)
	if err != nil {
		return 0, err
	}
	defer src.Close()
	return io.Copy(out, src)
}

func (f *FileSystemBaseOSS) Del(docId, hashName string) error {
	if hashName == "" {
		return nil
	}
	p := filepath.Join(f.base, docId, hashName)
	zap.L().Debug("删除文件对象", zap.String("file", p))
	return os.Remove(p)
}

func (f FileSystemBaseOSS) DelDoc(docId string) error {
	if docId == "" {
		return nil
	}
	p := filepath.Join(f.base, docId)
	zap.L().Debug("删除文档关联所有资源", zap.String("file", p))
	return os.RemoveAll(p)
}
