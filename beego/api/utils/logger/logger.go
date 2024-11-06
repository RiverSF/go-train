package logger

import (
	"api/utils/common"
	"errors"
	"fmt"
	"reflect"
	"runtime/debug"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

var (
	logDebug    *logs.BeeLogger
	logInfo     *logs.BeeLogger
	logNotice   *logs.BeeLogger
	logWarn     *logs.BeeLogger
	logError    *logs.BeeLogger
	logCritical *logs.BeeLogger

	logConfig         string
	logDir            = beego.AppConfig.DefaultString("log_dir", "logs")
	logLevel          = beego.AppConfig.DefaultInt("log_level", logs.LevelInfo)
	logRotateUnit     = beego.AppConfig.DefaultString("log_rotate_unit", "daily")
	logRotateMaxValue = beego.AppConfig.DefaultInt("log_rotate_max", 7)
	logRotateMaxKey   = ""
)

func Init() error {
	if logRotateUnit == "daily" {
		logRotateMaxKey = "maxdays"
	} else if logRotateUnit == "hourly" {
		logRotateMaxKey = "maxhours"
	} else {
		return errors.New("fail to init logger: logRotateUnit=" + logRotateUnit + ", err=illegal logRotateUnit")
	}

	var channelLens int64 = 65536
	logDebug = logs.NewLogger(channelLens)
	logInfo = logs.NewLogger(channelLens)
	logNotice = logs.NewLogger(channelLens)
	logWarn = logs.NewLogger(channelLens)
	logError = logs.NewLogger(channelLens)
	logCritical = logs.NewLogger(channelLens)

	logInit(logDebug, "debug")
	logInit(logInfo, "info")
	logInit(logNotice, "notice")
	logInit(logWarn, "warn")
	logInit(logError, "error")
	logInit(logCritical, "critical")

	fmt.Println("init logger: logConfig=%v", logConfig)
	return nil
}

func logInit(log *logs.BeeLogger, logFileName string) error {
	log.SetLogFuncCallDepth(3)
	log.EnableFuncCallDepth(true)

	if logLevel >= logs.LevelDebug {
		if err := log.SetLogger(logs.AdapterConsole, ""); err != nil {
			return errors.New("fail to init logger: err=" + err.Error())
		}
	}

	logConfig = fmt.Sprintf(`{"filename":"logs/%v/%v.log","level":%v,"maxlines":0,"maxsize":0,"%v":true,"%v":%v,"color":true}`,
		logDir, logFileName, logLevel, logRotateUnit, logRotateMaxKey, logRotateMaxValue)
	if err := log.SetLogger(logs.AdapterFile, logConfig); err != nil {
		return errors.New("fail to init logger: err=" + err.Error())
	}

	return nil
}

func Debug(f interface{}, v ...interface{}) {
	if logLevel >= logs.LevelDebug {
		logDebug.Debug(formatLog(f, v...))
	}
}

func Info(f interface{}, v ...interface{}) {
	if logLevel >= logs.LevelInfo {
		logInfo.Info(formatLog(f, v...))
	}
}

func Notice(f interface{}, v ...interface{}) {
	logNotice.Notice(formatLog(f, v...))
}

func Warn(f interface{}, v ...interface{}) {
	logWarn.Warn(formatLog(f, v...))
}

func Error(f interface{}, v ...interface{}) {
	logError.Error(formatLog(f, v...))
}

func Critical(f interface{}, v ...interface{}) {
	logCritical.Critical(formatLog(f, v...))
	logCritical.Critical(GetDebugStack())
}

func formatLog(f interface{}, v ...interface{}) string {
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
		if len(v) == 0 {
			return msg
		}
		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
			//format string
		} else {
			//do not contain format char
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}

	for key, value := range v {
		if reflect.TypeOf(value).Kind() == reflect.Ptr {
			v[key] = common.String(value)
		}
	}

	return fmt.Sprintf(msg, v...)
}

// 当前堆栈错误
func GetDebugStack() string {
	return string(debug.Stack())
}
