package main

import (
	"fmt"

	"github.com/AitorATuin/calibrefs"
	"github.com/juju/loggo"
)

func main() {
	calibrefs.SetLogLevel(loggo.DEBUG)
	calibrefsSrv := calibrefs.NewCalibreFS()
	calibrefsSrv.Start(calibrefsSrv)
	calibrefsSrv.Debuglevel = 10
	calibrefs.Logger.Infof("Starting calibrefs server")
	err := calibrefsSrv.StartNetListener("tcp", "127.0.0.1:5640")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}
}
