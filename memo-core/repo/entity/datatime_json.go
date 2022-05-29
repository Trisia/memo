package entity

import (
	"fmt"
	"time"
)

// DateTime 时间格式化
type DateTime time.Time

// 东八 (8小时)
var cstZone = time.FixedZone("CST", 8*3600)

func (t DateTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).In(cstZone).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}
