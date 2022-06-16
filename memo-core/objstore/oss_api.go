package objstore

import "io"

// Repo 存储仓库接口，需要驱动程序实现该接口
var Repo ObjectStorageService

// ObjectStorageService 对象存储服务接口
type ObjectStorageService interface {
	// Put 存储对象
	// docId: 文档ID
	// filename: 原文件名
	// in: 文件流
	// return 存储路径, 写入字节数
	Put(docId, filename string, in io.Reader) (hashName string, written int64, err error)

	// Get 读取对象
	// docId: 文档ID
	// hashName: 文件Hash名称
	// out: 文件流
	// return 如果文件不存在返还 os.ErrNotExist
	Get(docId, hashName string, out io.Writer) (written int64, err error)

	// Del 删除对象
	// docId: 文档ID
	// hashName: 文件Hash名称
	Del(docId, hashName string) error

	// DelDoc 删除文档相关的所有资源
	// docId: 文档ID
	DelDoc(docId string) error
}
