package main

import "fmt"
import "github.com/AitorATuin/calibrefs"

func main() {
	calibrefsSrv := calibrefs.NewCalibreFS()
	fmt.Println("Starting calibrefs server ...")
	calibrefsSrv.StartNetListener("tcp", "127.0.0.1")
}
