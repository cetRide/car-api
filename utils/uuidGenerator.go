package util

import (
	"strconv"

	"github.com/sony/sonyflake"
)

func GenerateUID() string {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, _ := flake.NextID()

	flakeId := strconv.FormatUint(id, 10)
	return flakeId
}
