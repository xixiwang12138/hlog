package tools

import (
	"context"
	"github.com/bwmarrin/snowflake"
)

type snowflakeId struct {
	*snowflake.Node
}

func newSnowflakeId(nodeId int64) *snowflakeId {
	node, err := snowflake.NewNode(nodeId)
	if err != nil {
		panic(err)
	}
	return &snowflakeId{Node: node}
}

func (s *snowflakeId) NextId(ctx context.Context) string {
	return s.Node.Generate().Base64()
}
