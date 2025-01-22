package reader

import (
	"context"
	"github.com/wangwenjie2500/RedisShake/pkg/entry"
	"github.com/wangwenjie2500/RedisShake/pkg/status"
)

type Reader interface {
	status.Statusable
	StartRead(ctx context.Context) []chan *entry.Entry
}
