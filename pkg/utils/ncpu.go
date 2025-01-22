package utils

import (
	"runtime"

	"github.com/wangwenjie2500/RedisShake/pkg/config"
	"github.com/wangwenjie2500/RedisShake/pkg/log"
)

func SetNcpu() {
	if config.Opt.Advanced.Ncpu != 0 {
		log.Infof("set ncpu to %d", config.Opt.Advanced.Ncpu)
		runtime.GOMAXPROCS(config.Opt.Advanced.Ncpu)
		log.Infof("set GOMAXPROCS to %v", config.Opt.Advanced.Ncpu)
	} else {
		log.Infof("GOMAXPROCS defaults to the value of runtime.NumCPU [%v]", runtime.NumCPU())
	}
}
