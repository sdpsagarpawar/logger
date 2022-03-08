package logger

import "log"

// Warning prints normal Warning log
func (l *logger) Warning(msg ...interface{}) {
	log.Print(warning, msg)
}

// Warningf prints Warning log with format
func (l *logger) Warningf(format string, body ...interface{}) {
	log.Printf(warning+format, body)
}

// Warningln prints log with println Warning
func (l *logger) Warningln(body ...interface{}) {
	log.Println(warning, body)
}
