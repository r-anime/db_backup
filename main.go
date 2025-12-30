package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

func main() {
	logLevel := pflag.StringP("log-level", "l", "INFO", "Optional log level (INFO, DEBUG, etc)")
	backupType := pflag.StringP("backup-type", "b", "manual", "enable verbose mode")

	pflag.Parse()

	fmt.Println("Test")
	fmt.Println("logLevel", *logLevel)
	fmt.Println("backupType", *backupType)
}
