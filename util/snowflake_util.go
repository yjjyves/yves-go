package util

import (
	"strconv"
	"time"
)
import "github.com/sony/sonyflake"

// 创建 sonyflake 实例
var settings = sonyflake.Settings{
	StartTime: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	MachineID: func() (uint16, error) {
		// 返回机器 ID，可以是 IP 地址哈希或其他唯一标识
		return 1, nil
	},
}

func NewSnowflake() string {
	sf := sonyflake.NewSonyflake(settings)
	if sf == nil {
		panic("NewSnowflake not created")
	}
	id, err := sf.NextID()
	if err != nil {
		panic(err)
	}
	return strconv.FormatUint(id, 16)
}
