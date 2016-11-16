package calibrefs

import "github.com/juju/loggo"

var Logger loggo.Logger = loggo.GetLogger("calibrefs")

func SetLogLevel(level loggo.Level) {
	Logger.SetLogLevel(level)
}
