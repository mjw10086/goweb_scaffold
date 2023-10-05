package config

import (
	"path/filepath"

	"github.com/alecthomas/kingpin/v2"
)

var (
	Port    = "8080"
	BaseDir = "."
	LogDir  = "log"
)

func Init() {
	paramParse()
}

func paramParse() {
	port := kingpin.Flag("bind", "port to bind server").Default(Port).Short('b').String()
	baseDir := kingpin.Flag("dir", "base dir to store log and articles").Default(BaseDir).Short('d').String()
	kingpin.Parse()

	Port = *port
	BaseDir = *baseDir
	LogDir = filepath.Join(BaseDir, LogDir)
}
