package imgstore

import "io"

// ObjectStorageService 对象存储服务接口
type ObjectStorageService interface {
	// Put 存储对象
	// oid: 文件名
	// in: 文件流
	Put(oid string, in io.Reader) (int, error)

	// Get 读取对象
	// oid: 文件名
	// out: 文件流
	Get(oid string, out io.Writer) (int, error)
}
