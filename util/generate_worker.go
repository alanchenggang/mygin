package util

import (
	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func init() {
	node, _ = snowflake.NewNode(1)
}

// 生成 64 位的 雪花 ID
func GenID() int64 {
	return node.Generate().Int64()
}
