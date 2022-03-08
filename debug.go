package logger

import (
	"log"
)

// Debug prints normal Debug log
func (l *logger) Debug(msg ...interface{}) {
	log.Print(debug, msg)
}

// Debugf prints Debug log with format
func (l *logger) Debugf(format string, body ...interface{}) {
	log.Printf(debug+format, body)
}

// Debugln prints log with println Debug
func (l *logger) Debugln(body ...interface{}) {
	log.Println(debug, body)
}
