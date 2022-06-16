package driver

import (
	"memo-core/objstore"
	"memo-core/objstore/filesys"
)

func init() {
	objstore.Repo = filesys.NewFileSystemBaseOSS()
}
