package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

func WriteLogs(clientName, logLine string) error {
	// Get current time
	now := time.Now()
	// Prepare file path
	logPath := fmt.Sprintf("/var/log/remote/logs/%d/%d/%d/%s.log", now.Year(), now.Month(), now.Day(), clientName)

	f, err := os.OpenFile(logPath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		return err
	}
	defer f.Close()
	logger := log.New(f, "", log.LstdFlags)
	// log to append
	logger.Println(logLine)
	return nil
}
