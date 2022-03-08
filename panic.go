package logger

import "log"

type Logger interface {
	Panic(interface{})
	Panicf(string, interface{})
	Panicln(interface{})
}

type logger struct{}

func NewLogger() Logger {
	return &logger{}
}

//Panic use to panic with normal message
func (l *logger) Panic(msg interface{}) {
	log.Panic(msg)
}

//Panicf used to panic with formatted message
func (l *logger) Panicf(format string, body interface{}) {
	log.Panicf(format, body)
}

//Panicln used to panic with println
func (l *logger) Panicln(msg interface{}) {
	log.Panicln(msg)
}
