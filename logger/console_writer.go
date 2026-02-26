package logger

import "fmt"

type ConsoleWriter struct{}

func (c *ConsoleWriter) Write(message string) {
	fmt.Println(message)
}
