package rpc

import (
	"github.com/seracoin/serad/infrastructure/logger"
	"github.com/seracoin/serad/util/panics"
)

var log = logger.RegisterSubSystem("RPCS")
var spawn = panics.GoroutineWrapperFunc(log)
