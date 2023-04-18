package tools

import (
	"strconv"
	"time"
)

const Epoch uint64 = 1420070400000

func ConvertToTime(flake string) time.Time {
	id, _ := strconv.ParseUint(flake, 10, 64)

	return time.UnixMilli(int64(id>>22 + Epoch))
}