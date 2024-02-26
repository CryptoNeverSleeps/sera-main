package consensus

import (
	"github.com/seracoin/serad/infrastructure/logger"
	"github.com/seracoin/serad/util/panics"
)

var log = logger.RegisterSubSystem("BDAG")
var spawn = panics.GoroutineWrapperFunc(log)
