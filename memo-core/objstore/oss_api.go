package objstore

import "io"

// ObjectStorageService 对象存储服务接口
type ObjectStorageService interface {
	// Put 存储对象
	// filename: 原文件名
	// in: 文件流
	// return 存储路径
	Put(filename string, in io.Reader) (loc string, error error)

	// Get 读取对象
	// loc: 存储路径
	// out: 文件流
	Get(loc string, out io.Writer) (int, error)

	// Del 删除对象
	Del(loc string) error
}
