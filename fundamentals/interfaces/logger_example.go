package main

import "fmt"

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {
	fmt.Printf("[CONSOLE] %s\n", message)
}

type FileLogger struct{}

func (f FileLogger) Log(message string) {
	fmt.Printf("[FILE] %s\n", message)
}

type DatadogLogger struct{}

func (d DatadogLogger) Log(message string) {
	fmt.Printf("[DATADOG] %s\n", message)
}

// processOrder depends on the logger beahviour and not the specific implementation.
func processOrder(logger Logger, orderID string) {
	logger.Log("Processing order: " + orderID)
}

func loggerExample() {
	console := ConsoleLogger{}
	file := FileLogger{}
	datadog := DatadogLogger{}

	processOrder(console, "ORD-1001")
	processOrder(file, "ORD-1001")
	processOrder(datadog, "ORD-1001")
}