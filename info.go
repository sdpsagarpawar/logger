package logger

import (
	"log"
)

// Info prints normal info log
func (l *logger) Info(msg ...interface{}) {
	log.Print(info, msg)
}

// Infof prints info log with format
func (l *logger) Infof(format string, body ...interface{}) {
	log.Printf(info+format, body)
}

// Infoln prints log with println info
func (l *logger) Infoln(body ...interface{}) {
	log.Println(info, body)
}
