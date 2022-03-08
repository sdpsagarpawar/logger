package logger

import "log"

// Error prints normal error log
func (l *logger) Error(msg ...interface{}) {
	log.Print(err, msg)
}

// Errorf prints Error log with format
func (l *logger) Errorf(format string, body ...interface{}) {
	log.Printf(err+format, body)
}

// Errorln prints log with println Error
func (l *logger) Errorln(body ...interface{}) {
	log.Println(err, body)
}
