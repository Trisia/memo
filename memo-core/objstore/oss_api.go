package objstore

import "io"

// Repo 存储仓库接口，需要驱动程序实现该接口
var Repo ObjectStorageService

// ObjectStorageService 对象存储服务接口
type ObjectStorageService interface {
	// Put 存储对象
	// filename: 原文件名
	// in: 文件流
	// return 存储路径, 写入字节数
	Put(filename string, in io.Reader) (string, int64, error)

	// Get 读取对象
	// loc: 存储路径
	// out: 文件流
	Get(loc string, out io.Writer) (int64, error)

	// Del 删除对象
	Del(loc string) error
}
