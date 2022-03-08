package logger

import "log"

//Panic use to panic with normal message
func (l *logger) Panic(msg ...interface{}) {
	log.Panic(panic, msg)
}

//Panicf used to panic with formatted message
func (l *logger) Panicf(format string, body ...interface{}) {
	log.Panicf(panic+format, body)
}

//Panicln used to panic with println
func (l *logger) Panicln(body ...interface{}) {
	log.Panicln(panic, body)
}
