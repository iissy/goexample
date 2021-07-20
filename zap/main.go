package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// logpath 日志文件路径
// loglevel 日志级别
func initLogger(logpath string, loglevel string) *zap.Logger {

	hook := lumberjack.Logger{
		Filename:   logpath, // 日志文件路径
		MaxSize:    1,       // megabytes
		MaxBackups: 30,      // 最多保留300个备份
		MaxAge:     7,       // days
		Compress:   false,   // 是否压缩 disabled by default
	}

	// 设置日志级别,debug可以打印出info,debug,warn；info级别可以打印warn，info；warn只能打印warn
	// debug->info->warn->error
	var level zapcore.Level
	switch loglevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	// 彩色显示，纯控制台输出可以彩色输出
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	// 全路径地址，通常不需要
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder

	core := zapcore.NewCore(
		//获取编码器,NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook), zapcore.AddSync(os.Stdout)),
		level,
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(2))
	return logger
}

type Test struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	t := &Test{
		Name: "xiaoming",
		Age:  12,
	}
	data, err := json.Marshal(t)
	if err != nil {
		fmt.Println("marshal is failed,err: ", err)
	}

	// 历史记录日志名字为：all-2018-11-15T07-45-51.763.log，服务重新启动，日志会追加，不会删除
	logger := initLogger("./iissy.log", "debug")
	logger.Info(fmt.Sprint("test log ", 1), zap.Int("line", 47))
	logger.Debug(fmt.Sprint("debug log ", 1), zap.ByteString("level", data))
	logger.Info(fmt.Sprint("Info log ", 1), zap.String("level", `{"a":"4","b":"5"}`))
	logger.Warn(fmt.Sprint("Info log ", 1), zap.String("level", `{"a":"7","b":"8"}`))
	logger.Error(fmt.Sprint("error log ", 1), zap.String("level", "i am error."))
	logger.Error(fmt.Sprint("error log ", 1), zap.Any("code", 48), zap.Any("message", "error in dir."))
}
