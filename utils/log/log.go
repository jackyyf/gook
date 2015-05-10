package log

import (
	"fmt"
	"github.com/astaxie/beego/config"
	_log "github.com/astaxie/beego/logs"
)

var logger *_log.BeeLogger

func init() {
	logger = _log.NewLogger(16384)
	logger.EnableFuncCallDepth(true)
	logger.SetLogFuncCallDepth(3)
	conf, err := config.NewConfig("ini", "conf/log.ini")
	if err != nil {
		logger.SetLogger("console", `{"level":7}`)
		return
	}
	level := conf.DefaultInt("level", _log.LevelTrace)
	if name := conf.String("filename"); name != "" {
		print(name)
		logger.SetLogger("file", fmt.Sprintf(`{"filename":"%q","level":%d}`, name, level))
	} else {
		logger.SetLogger("console", fmt.Sprintf(`{"level":%d}`, level))
	}
}

func Crit(format string, v ...interface{}) {
	logger.Critical(format, v...)
}

func Emerg(format string, v ...interface{}) {
	logger.Emergency(format, v...)
}

func Alert(format string, v ...interface{}) {
	logger.Alert(format, v...)
}

func Error(format string, v ...interface{}) {
	logger.Error(format, v...)
}

func Warn(format string, v ...interface{}) {
	logger.Warning(format, v...)
}

func Notice(format string, v ...interface{}) {
	logger.Notice(format, v...)
}

func Info(format string, v ...interface{}) {
	logger.Informational(format, v...)
}

func Debug(format string, v ...interface{}) {
	logger.Debug(format, v...)
}
