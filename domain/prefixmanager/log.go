package prefixmanager

import (
	"github.com/seracoin/serad/infrastructure/logger"
	"github.com/seracoin/serad/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)
