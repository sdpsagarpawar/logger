package logger

import (
	"log"
)

// Fatal prints normal Fatal log
func (l *logger) Fatal(msg ...interface{}) {
	log.Fatal(fatal, msg)
}

// Fatalf prints Fatal log with format
func (l *logger) Fatalf(format string, body ...interface{}) {
	log.Fatalf(fatal+format, body)
}

// Fatalln prints log with println Fatal
func (l *logger) Fatalln(body ...interface{}) {
	log.Fatalln(fatal, body)
}
