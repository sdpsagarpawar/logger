# Logger Library

The Logger library provides a simple and flexible logging mechanism for Go applications. It allows you to log messages with different log levels, such as debug, info, warning, error, and panic.

## Installation

To use the Logger library in your Go project, you need to import it:

``` 
go import "github.com/sdpsagarpawar/logger" 
```

## Usage
```
logger := logger.NewLogger()
```

## Logging Levels
The Logger library supports the following log levels:

Debug: Used for detailed debugging information.
Info: Used for general information messages.
Warning: Used to indicate potential issues or warnings.
Error: Used to log error messages.
Panic: Used to log critical errors and panic the application.
Fatal: Used to log fatal errors and exit the application

## Logging Methods
The Logger library provides the following logging methods:

Debug(values ...interface{}): Logs a debug message.

Debugf(format string, values ...interface{}): Logs a formatted debug message.

Debugln(values ...interface{}): Logs a debug message with a newline character.

Info(values ...interface{}): Logs an info message.

Infof(format string, values ...interface{}): Logs a formatted info message.

Infoln(values ...interface{}): Logs an info message with a newline character.

Warning(values ...interface{}): Logs a warning message.

Warningf(format string, values ...interface{}): Logs a formatted warning message.

Warningln(values ...interface{}): Logs a warning message with a newline character.

Error(values ...interface{}): Logs an error message.

Errorf(format string, values ...interface{}): Logs a formatted error message.

Errorln(values ...interface{}): Logs an error message with a newline character.

Panic(values ...interface{}): Logs a panic message and panics the application.

Panicf(format string, values ...interface{}): Logs a formatted panic message and panics the application.

Panicln(values ...interface{}): Logs a panic message with a newline character and panics the application.

Fatal(values ...interface{}): Logs a fatal error message and exits the application.

Fatalf(format string, values ...interface{}): Logs a formatted fatal error message and exits the application.

Fatalln(values ...interface{}): Logs a fatal error message with a newline character and exits the application.

Log(values ...interface{}): Logs a log message.

Logf(format string, values ...interface{}): Logs a formatted log message.

Logln(values ...interface{}): Logs a log message with a newline character.

## Example
Here's an example of how to use the Logger library:
```
package main

import (
	"github.com/sdpsagarpawar/logger"
)

func main() {
	// Create a new logger
	logger := logger.NewLogger()

	// Log messages
	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warning("Warning message")
	logger.Error("Error message")
	logger.Panic("Panic message")
	logger.Fatal("Fatal message")
}


```
This example demonstrates how to create a logger instance and use various logging methods to log messages with different log levels.

## License
This project is licensed under the MIT License.

```

Feel free to customize the README.md file according to your needs, including the installation instructions, usage examples, and licensing information.

```