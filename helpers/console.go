package helpers

import (
	"fmt"
	"time"
)

// LogInfo :nodoc:
func LogInfo(message string) {
	time := time.Now().Format("2006-01-02 15:04:05")

	fmt.Printf("%s [info] %s\n", time, message)
}

// LogError :nodoc:
func LogError(err error) {
	time := time.Now().Format("2006-01-02 15:04:05")

	fmt.Printf("%s [error] %s\n", time, err.Error())
}
