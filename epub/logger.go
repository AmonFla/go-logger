package logger

import (
	"encoding/json"
	"fmt"
	"log/syslog"
	"sync"
	"time"

	tools "github.com/AmonFla/go-tools/epub"
)

// global package
var l LoggerStruct
var once sync.Once

type ConfigStruct struct {
	AppName            string
	DefaultSyslogLevel string
	Method             string
	RemoteSyslog       struct {
		IP   string
		Port string
		Type string
	}
	NotifyLevel []string
	ConsoleLog  bool
}

type LoggerStruct struct {
	prio      map[string]syslog.Priority
	logwriter *syslog.Writer
	config    ConfigStruct
}

type LoggerMessageStruct struct {
	Date     string
	Level    string
	Message  string
	File     string
	Function string
	Flag     int
}

func init() {
	l.prio = map[string]syslog.Priority{
		"LOG_EMERG":   syslog.LOG_EMERG,
		"LOG_ALERT":   syslog.LOG_ALERT,
		"LOG_CRIT":    syslog.LOG_CRIT,
		"LOG_ERR":     syslog.LOG_ERR,
		"LOG_WARNING": syslog.LOG_WARNING,
		"LOG_NOTICE":  syslog.LOG_NOTICE,
		"LOG_INFO":    syslog.LOG_INFO,
		"LOG_DEBUG":   syslog.LOG_DEBUG,
	}
}

func connect() {
	var err error
	switch l.config.Method {
	case "RemoteSyslog":
		//remoto
		l.logwriter, err = syslog.Dial(l.config.RemoteSyslog.Type, l.config.RemoteSyslog.IP+":"+l.config.RemoteSyslog.Port, l.prio[l.config.DefaultSyslogLevel], l.config.AppName)
	default:
		//local
		l.logwriter, err = syslog.New(l.prio[l.config.DefaultSyslogLevel], l.config.AppName)
	}
	if err != nil {
		panic("enable to start log")
	}
}

func NewLogger(c ConfigStruct) LoggerStruct {
	l.config = c
	once.Do(connect)
	return l
}

func (log LoggerStruct) LogData(level, msg, file, function string, flag int) {
	if tools.StringInArray(level, log.config.NotifyLevel) {
		currentTime := time.Now().Local()
		m, _ := json.Marshal(LoggerMessageStruct{currentTime.Format("2006-01-02 15:04:05.000"), level, msg, file, function, flag})
		switch level {
		case "LOG_EMERG":
			log.logwriter.Emerg(string(m))
		case "LOG_ALERT":
			log.logwriter.Alert(string(m))
		case "LOG_CRIT":
			log.logwriter.Crit(string(m))
		case "LOG_ERR":
			log.logwriter.Err(string(m))
		case "LOG_WARNING":
			log.logwriter.Warning(string(m))
		case "LOG_NOTICE":
			log.logwriter.Notice(string(m))
		case "LOG_DEBUG":
			log.logwriter.Debug(string(m))
		default:
			log.logwriter.Info(string(m))
		}
		if log.config.ConsoleLog {
			fmt.Println(string(m))
		}
	}
}
