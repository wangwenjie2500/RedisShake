package writer

import (
	"context"
	"github.com/wangwenjie2500/RedisShake/pkg/entry"
	"github.com/wangwenjie2500/RedisShake/pkg/status"
)

type Writer interface {
	status.Statusable
	Write(entry *entry.Entry)
	StartWrite(ctx context.Context) (ch chan *entry.Entry)
	Close()
}
