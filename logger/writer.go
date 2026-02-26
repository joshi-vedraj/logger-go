package logger

type Writer interface {
	Write(message string)
}
