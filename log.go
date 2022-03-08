package logger

type Logger interface {
	Panic(...interface{})
	Panicf(string, ...interface{})
	Panicln(...interface{})
	Info(...interface{})
	Infof(string, ...interface{})
	Infoln(...interface{})
	Error(...interface{})
	Errorf(string, ...interface{})
	Errorln(...interface{})
	Warning(...interface{})
	Warningf(string, ...interface{})
	Warningln(...interface{})
	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})
	Debug(...interface{})
	Debugf(string, ...interface{})
	Debugln(...interface{})
}

const (
	info    = "INFO: "
	panic   = "PANIC: "
	err     = "ERROR: "
	warning = "WARNING: "
	fatal   = "FATAL: "
	debug   = "DEBUG: "
)

type logger struct{}

func NewLogger() Logger {
	return &logger{}
}
