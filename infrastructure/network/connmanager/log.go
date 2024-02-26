package connmanager

import (
	"github.com/seracoin/serad/infrastructure/logger"
	"github.com/seracoin/serad/util/panics"
)

var log = logger.RegisterSubSystem("CMGR")
var spawn = panics.GoroutineWrapperFunc(log)
