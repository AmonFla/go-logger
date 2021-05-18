package main

import (
	logger "github.com/AmonFla/go-logger/epub"
)

func main() {
	c := logger.ConfigStruct{
		AppName:            "app-test",
		DefaultSyslogLevel: "LOG_NOTICE",
		Method:             "Local",
		RemoteSyslog: logger.RemoteSyslogStruct{
			IP:   "127.0.0.1",
			Port: "514",
			Type: "tcp"},
		NotifyLevel: []string{"LOG_EMERG", "LOG_ALERT", "LOG_CRIT", "LOG_ERR", "LOG_WARNING", "LOG_NOTICE", "LOG_INFO", "LOG_DEBUG"},
		ConsoleLog:  true}

	logger.Init(c)
	l := logger.NewLogger()
	l.LogData("LOG_CRIT", "esto es una prueba", "main.go", "main", 12)

}
