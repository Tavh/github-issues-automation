package logs

import (
	"log"
	"os"
	"strings"
)

var debugActive bool

func Init() {
	debugActive = os.Getenv("DEBUG") != ""
}

func Debug(msg string, args ...any) {
	if debugActive {
		log.Printf(strings.Join([]string{"[DEBUG]", msg}, " "), args...)
	}
}

func Info(msg string, args ...any) {
	log.Printf(strings.Join([]string{"[INFO]", msg}, " "), args...)
}

func Error(err error) {
	log.Printf("[ERROR] %+v\n", err)
}
