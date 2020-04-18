package logutil

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"sync"
)

var once sync.Once

var logger *zap.Logger

// GetDefaultLogger 获取缺省logger
func GetDefaultLogger() *zap.Logger {
	return initLogger(zapcore.DebugLevel, "./app.log")
}

// GetLogger 获取logger实例
func GetLogger(level zapcore.Level, logFilePath string) *zap.Logger {
	return initLogger(level, logFilePath)
}

// initLogger 初始化logger对象
// 用once控制，只初始化一次
func initLogger(level zapcore.Level, logFilePath string) *zap.Logger {
	once.Do(func() {
		hook := lumberjack.Logger{
			Filename:   logFilePath, // 日志文件路径
			MaxSize:    128,         // 每个日志文件保存的最大尺寸 单位：M
			MaxBackups: 30,          // 日志文件最多保存多少个备份
			MaxAge:     7,           // 文件最多保存多少天
			Compress:   true,        // 是否压缩
		}

		encoderConfig := zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "linenum",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
			EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
			EncodeDuration: zapcore.SecondsDurationEncoder, //
			EncodeCaller:   zapcore.ShortCallerEncoder,     // 全路径编码器
			EncodeName:     zapcore.FullNameEncoder,
		}

		// 设置日志级别
		atomicLevel := zap.NewAtomicLevel()
		atomicLevel.SetLevel(level)

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
			atomicLevel, // 日志级别
		)

		// 开启开发模式，堆栈跟踪
		caller := zap.AddCaller()
		// 开启文件及行号
		development := zap.Development()
		// 构造日志
		logger = zap.New(core, caller, development)
	})

	return logger
}
