package logger

import (
	"sync"
	"time"
)

type Logger struct {
	writers []Writer
	level   Level
}

var instance *Logger
var once sync.Once

func GetLogger() *Logger {
	once.Do(func() {
		instance = &Logger{
			writers: []Writer{&ConsoleWriter{}}, // default
			level:   INFO,
		}
	})
	return instance
}

func (l *Logger) AddWriter(w Writer) {
	l.writers = append(l.writers, w)
}

func (l *Logger) SetLevel(level Level) {
	l.level = level
}

func (l *Logger) log(msgLevel Level, label string, msg string) {
	if msgLevel < l.level {
		return
	}

	message := time.Now().Format("15:04:05") + " " + label + " " + msg

	for _, w := range l.writers {
		w.Write(message)
	}

}

func (l *Logger) Debug(msg string) { l.log(DEBUG, "DEBUG", msg) }
func (l *Logger) Info(msg string)  { l.log(INFO, "INFO ", msg) }
func (l *Logger) Warn(msg string)  { l.log(WARN, "WARN ", msg) }
func (l *Logger) Error(msg string) { l.log(ERROR, "ERROR", msg) }
