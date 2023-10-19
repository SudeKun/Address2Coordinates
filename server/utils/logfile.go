package utils

import "os"

var logFilePath = "permanent/log"

func SetupLogging() *os.File {
	logFile, _ := os.Create(logFilePath)
	// Set the log output to the log file
	return logFile
}
