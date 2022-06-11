package dto

type DocSearch struct {
	Limit   int    `form:"limit"`   // 每页大小，默认30
	Offset  int    `form:"offset"`  // 偏移量
	Keyword string `form:"keyword"` // 关键字
}
