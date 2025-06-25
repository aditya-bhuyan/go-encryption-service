package logger

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

type logEntry struct {
	Level     string `json:"level"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

var (
	logFile     *os.File
	logMutex    sync.Mutex
	infoWriter  io.Writer
	errorWriter io.Writer
)

func init() {
	var err error
	logFile, err = os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}

	// Multi-writer for stdout and file
	infoWriter = io.MultiWriter(os.Stdout, logFile)
	errorWriter = io.MultiWriter(os.Stderr, logFile)
}

func logJSON(level, msg string, writer io.Writer) {
	entry := logEntry{
		Level:     level,
		Message:   msg,
		Timestamp: time.Now().Format(time.RFC3339),
	}
	logMutex.Lock()
	defer logMutex.Unlock()
	jsonData, _ := json.Marshal(entry)
	writer.Write(jsonData)
	writer.Write([]byte("\n"))
}

func Info(msg string) {
	logJSON("INFO", msg, infoWriter)
}

func Error(msg string) {
	logJSON("ERROR", msg, errorWriter)
}

// Usage:
// logger.Info("Service started")
// logger.Error("Something went wrong")
