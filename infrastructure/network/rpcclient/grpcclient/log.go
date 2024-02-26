package grpcclient

import (
	"github.com/seracoin/serad/infrastructure/logger"
	"github.com/seracoin/serad/util/panics"
)

var log = logger.RegisterSubSystem("RPCC")
var spawn = panics.GoroutineWrapperFunc(log)
