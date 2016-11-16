package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/AitorATuin/calibrefs"
	"github.com/juju/loggo"
)

var addr = flag.String("addr", ":5640", "network address")

func main() {
	flag.Parse()
	calibrefs.SetLogLevel(loggo.DEBUG)
	calibrefsSrv := calibrefs.NewCalibreFS()
	calibrefsSrv.Dotu = true
	calibrefsSrv.Id = "calibrefs"
	if !calibrefsSrv.Start(calibrefsSrv) {
		fmt.Fprintf(os.Stderr, "Unable to start CalibreFS Server\n")
		os.Exit(1)
	}
	calibrefsSrv.Debuglevel = 10
	calibrefs.Logger.Infof("Starting calibrefs server")
	err := calibrefsSrv.StartNetListener("tcp", *addr)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}
}
