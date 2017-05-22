package conf

import (
	log4 "github.com/alecthomas/log4go"
)

func OpenDefaultLog(filename string) {
	logger := log4.NewFileLogWriter(filename, true)
	logger.SetRotateDaily(true)
	logger.SetRotateMaxBackup(30)
	log4.AddFilter("default", log4.FINE, logger)
}

func OpenCustomLog(filename string) {
	log4.LoadConfiguration(filename)
}
