package log

import (
	"fmt"
	"log"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"apple/common/cfg"
)

var (
	logInfo   *zap.SugaredLogger
	logErr    *zap.SugaredLogger
	logAccess *zap.SugaredLogger
	logDebug  *zap.SugaredLogger
	logWarn   *zap.SugaredLogger

	loggerConfig *cfg.LoggerConfig
)

// 判断是否已初始化
func AlreadyInit() bool {
	return logInfo != nil && logErr != nil && logAccess != nil && logDebug != nil
}

// 初始化日志配置
func InitLog(config *cfg.LoggerConfig) {
	if AlreadyInit() {
		return
	}

	loggerConfig = config

	// todo kafka
	if config.EnableKafka {
		//InitKafkaProducer(config.Kafka.NameServer)
	}

	// 创建日志目录
	if err := MakeDir(config.Base.LogPath); err != nil {
		panic(err)
	}

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "ts"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	// 增加几个初始字段
	additionalFields := initAdditionFields(config)

	var core zapcore.Core
	var logLevel zapcore.Level

	switch config.Base.LogLevel {
	case "debug":
		logLevel = zap.DebugLevel
	case "info":
		logLevel = zap.InfoLevel
	case "warn":
		logLevel = zap.WarnLevel
	case "error":
		logLevel = zap.ErrorLevel
	default:
		logLevel = zap.InfoLevel
	}

	//  配置一个 debug
	dw := zapcore.AddSync(&lumberjack.Logger{
		Filename:   strings.Join([]string{config.Base.LogPath, "debug.log"}, "/"),
		MaxSize:    200, // megabytes
		MaxBackups: 7,
		MaxAge:     3, // days
		LocalTime:  false,
	})

	core = zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		dw,
		logLevel,
	)
	dLogger := zap.New(core)
	logDebug = dLogger.Sugar()

	//  配置一个 access
	aw := zapcore.AddSync(&lumberjack.Logger{
		Filename:   strings.Join([]string{config.Base.LogPath, "access.log"}, "/"),
		MaxSize:    200, // megabytes
		MaxBackups: 7,
		MaxAge:     3, // days
		LocalTime:  false,
	})

	core = zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		aw,
		logLevel,
	)
	aLogger := zap.New(core)
	logAccess = aLogger.Sugar()

	// 配置一个 log info
	iw := zapcore.AddSync(&lumberjack.Logger{
		Filename:   strings.Join([]string{config.Base.LogPath, "info.log"}, "/"),
		MaxSize:    200, // megabytes
		MaxBackups: 7,
		MaxAge:     3, // days
		LocalTime:  false,
	})

	core = zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		iw,
		logLevel,
	)

	iLogger := zap.New(core)
	logInfo = iLogger.Sugar()

	// 配置一个 log warn
	ww := zapcore.AddSync(&lumberjack.Logger{
		Filename:   strings.Join([]string{config.Base.LogPath, "warn.log"}, "/"),
		MaxSize:    200, // megabytes
		MaxBackups: 7,
		MaxAge:     3, // days
		LocalTime:  false,
	})

	encoderCfg.CallerKey = "caller"
	encoderCfg.StacktraceKey = "stacktrace"

	// 仅当配置了警告级别，才会往ES上发
	if config.EnableKafka && len(config.Kafka.WarnTopic) > 0 {
		//core = zapcore.NewCore(
		//	zapcore.NewJSONEncoder(encoderCfg),
		//	zapcore.NewMultiWriteSyncer(New(config.Kafka.WarnTopic, config.Kafka.Filter), ww),
		//	logLevel,
		//)
	} else {
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			ww,
			logLevel,
		)
	}

	wLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1+config.Base.CallerSkip), additionalFields,
		zap.AddStacktrace(zap.WarnLevel))
	logWarn = wLogger.Sugar()

	// 配置一个 log error
	ew := zapcore.AddSync(&lumberjack.Logger{
		Filename:   strings.Join([]string{config.Base.LogPath, "error.log"}, "/"),
		MaxSize:    200, // megabytes
		MaxBackups: 7,
		MaxAge:     3, // days
		LocalTime:  false,
	})

	if config.EnableKafka {
		//core = zapcore.NewCore(
		//	zapcore.NewJSONEncoder(encoderCfg),
		//	zapcore.NewMultiWriteSyncer(New(config.Kafka.ErrorTopic, config.Kafka.Filter), ew),
		//	logLevel,
		//)
	} else {
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			ew,
			logLevel,
		)
	}

	eLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1+config.Base.CallerSkip), additionalFields,
		zap.AddStacktrace(zap.ErrorLevel))
	logErr = eLogger.Sugar()
}

// 必须main函数退出前调用
func WaitForLogComplete() {
	if loggerConfig.EnableKafka {
		//CloseKafkaProducer()
	}
}

// 初始化几个额外的字段
func initAdditionFields(config *cfg.LoggerConfig) zap.Option {
	additionFields := zap.Fields(
		zap.String("serviceName", config.Base.ServiceName),
		zap.String("logPath", config.Base.LogPath))
	return additionFields
}

func FatalInTestEnv(args ...interface{}) {
	if logErr != nil {
		logErr.Error(args)
	} else {
		log.Println(args)
	}
}

func FatalfInTestEnv(format string, args ...interface{}) {
	errStr := fmt.Sprintf(format, args)
	FatalInTestEnv(errStr)
}

func Error(args ...interface{}) {
	if logErr != nil {
		logErr.Error(args)
	} else {
		log.Println(args)
	}
}

func Errorf(format string, args ...interface{}) {
	if logErr != nil {
		logErr.Errorf(format, args...)
	} else {
		log.Printf(format, args...)
	}
}

func Info(args ...interface{}) {
	if logInfo != nil {
		logInfo.Info(args)
	} else {
		log.Println(args)
	}
}

func Infof(format string, args ...interface{}) {
	if logInfo != nil {
		logInfo.Infof(format, args...)
	} else {
		log.Printf(format, args...)
	}
}

func Access(args ...interface{}) {
	if logAccess != nil {
		logAccess.Info(args)
	}
}

func Accessf(format string, args ...interface{}) {
	if logAccess != nil {
		logAccess.Infof(format, args...)
	}
}

func Debug(args ...interface{}) {
	if logDebug != nil {
		logDebug.Info(args)
	}
}

func Debugf(format string, args ...interface{}) {
	if logDebug != nil {
		logDebug.Infof(format, args...)
	}
}

func Warn(args ...interface{}) {
	if logWarn != nil {
		logWarn.Warn(args)
	}
}

func Warnf(format string, args ...interface{}) {
	if logWarn != nil {
		logWarn.Warnf(format, args...)
	}
}

func IsFileExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

func MakeDir(f string) error {
	if IsFileExist(f) {
		return nil
	}
	return os.MkdirAll(f, os.ModePerm)
}
