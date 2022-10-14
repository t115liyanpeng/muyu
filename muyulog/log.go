package muyulog

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

// 日志切割设置
func getLogWriter(p string, s, m int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   p,     // 日志文件位置
		MaxSize:    s,     // 日志文件最大大小(MB)
		MaxBackups: 500,   // 保留旧文件最大数量
		MaxAge:     m,     // 保留旧文件最长天数
		Compress:   false, // 是否压缩旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	// 使用默认的JSON编码
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// InitLogger 初始化Logger
func InitLogger(p string, s, m int) {
	writeSyncer := getLogWriter(p, s, m)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	Log = zap.New(core, zap.AddCaller())
}
