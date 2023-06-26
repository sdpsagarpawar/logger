package logger

import (
	"log"
	"os"
)

// LogLevel represents the log levels
type LogLevel int

const (
	// LogLevelDebug represents the debug log level
	LogLevelDebug LogLevel = iota

	// LogLevelInfo represents the info log level
	LogLevelInfo

	// LogLevelWarning represents the warning log level
	LogLevelWarning

	// LogLevelError represents the error log level
	LogLevelError
)

// Logger represents a logger instance
type Logger struct {
	debugLogger   *log.Logger
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
}

// NewLogger creates a new logger instance
func NewLogger() *Logger {
	logger := &Logger{}

	// Create the debug logger
	logger.debugLogger = log.New(os.Stdout, "[DEBUG] ", log.Ldate|log.Ltime)

	// Create the info logger
	logger.infoLogger = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)

	// Create the warning logger
	logger.warningLogger = log.New(os.Stdout, "[WARNING] ", log.Ldate|log.Ltime)

	// Create the error logger
	logger.errorLogger = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime)

	return logger
}

// Debug prints a debug message
func (logger *Logger) Debug(values ...interface{}) {
	logger.debugLogger.Println(values...)
}

// Debugf prints a formatted debug message
func (logger *Logger) Debugf(format string, values ...interface{}) {
	logger.debugLogger.Printf(format, values...)
}

// Debugln prints a debug message with a newline character
func (logger *Logger) Debugln(values ...interface{}) {
	logger.debugLogger.Println(values...)
}

// Warning prints a warning message
func (logger *Logger) Warning(values ...interface{}) {
	logger.warningLogger.Println(values...)
}

// Warningf prints a formatted warning message
func (logger *Logger) Warningf(format string, values ...interface{}) {
	logger.warningLogger.Printf(format, values...)
}

// Warningln prints a warning message with a newline character
func (logger *Logger) Warningln(values ...interface{}) {
	logger.warningLogger.Println(values...)
}

// Error prints an error message
func (logger *Logger) Error(values ...interface{}) {
	logger.errorLogger.Println(values...)
}

// Errorf prints a formatted error message
func (logger *Logger) Errorf(format string, values ...interface{}) {
	logger.errorLogger.Printf(format, values...)
}

// Errorln prints an error message with a newline character
func (logger *Logger) Errorln(values ...interface{}) {
	logger.errorLogger.Println(values...)
}

// Fatal prints a fatal error message and exits the program
func (logger *Logger) Fatal(values ...interface{}) {
	logger.errorLogger.Fatal(values...)
}

// Fatalf prints a formatted fatal error message and exits the program
func (logger *Logger) Fatalf(format string, values ...interface{}) {
	logger.errorLogger.Fatalf(format, values...)
}

// Fatalln prints a fatal error message with a newline character and exits the program
func (logger *Logger) Fatalln(values ...interface{}) {
	logger.errorLogger.Fatalln(values...)
}

// Log prints a log message
func (logger *Logger) Log(values ...interface{}) {
	logger.infoLogger.Println(values...)
}

// Logf prints a formatted log message
func (logger *Logger) Logf(format string, values ...interface{}) {
	logger.infoLogger.Printf(format, values...)
}

// Logln prints a log message with a newline character
func (logger *Logger) Logln(values ...interface{}) {
	logger.infoLogger.Println(values...)
}

// Panic logs a panic message and panics the application
func (logger *Logger) Panic(values ...interface{}) {
	logger.errorLogger.Panic(values...)
}

// Panicf logs a formatted panic message and panics the application
func (logger *Logger) Panicf(format string, values ...interface{}) {
	logger.errorLogger.Panicf(format, values...)
}

// Panicln logs a panic message with a newline character and panics the application
func (logger *Logger) Panicln(values ...interface{}) {
	logger.errorLogger.Panicln(values...)
}

// Info prints an info message
func (logger *Logger) Info(values ...interface{}) {
	logger.infoLogger.Println(values...)
}

// Infof prints a formatted info message
func (logger *Logger) Infof(format string, values ...interface{}) {
	logger.infoLogger.Printf(format, values...)
}

// Infoln prints an info message with a newline character
func (logger *Logger) Infoln(values ...interface{}) {
	logger.infoLogger.Println(values...)
}
