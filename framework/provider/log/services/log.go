package services

import (
	"context"
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework"
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework/contract"
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework/provider/log/formatter"
	"io"
	pkgLog "log"
	"time"
)

// YoimiyaLog 的通用实例
type YoimiyaLog struct {
	// 五个必要参数
	level      contract.LogLevel   // 日志级别
	formatter  contract.Formatter  // 日志格式化方法
	ctxFielder contract.CtxFielder // ctx获取上下文字段
	output     io.Writer           // 输出
	c          framework.Container // 容器
}

// IsLevelEnable 判断这个级别是否可以打印
func (log *YoimiyaLog) IsLevelEnable(level contract.LogLevel) bool {
	return level <= log.level
}

// logf 为打印日志的核心函数
func (log *YoimiyaLog) logf(level contract.LogLevel, ctx context.Context, msg string, fields map[string]interface{}) error {
	// 先判断日志级别
	if !log.IsLevelEnable(level) {
		return nil
	}

	// 使用ctxFielder 获取context中的信息
	fs := fields
	if log.ctxFielder != nil {
		t := log.ctxFielder(ctx)
		if t != nil {
			for k, v := range t {
				fs[k] = v
			}
		}
	}

	// 如果绑定了trace服务，获取trace信息
	if log.c != nil && log.c.IsBind(contract.TraceKey) {
		tracer := log.c.MustMake(contract.TraceKey).(contract.Trace)
		tc := tracer.GetTrace(ctx)
		if tc != nil {
			maps := tracer.ToMap(tc)
			for k, v := range maps {
				fs[k] = v
			}
		}
	}

	// 将日志信息按照formatter序列化为字符串
	if log.formatter == nil {
		log.formatter = formatter.TextFormatter
	}
	ct, err := log.formatter(level, time.Now(), msg, fs)
	if err != nil {
		return err
	}

	// 如果是panic级别，则使用log进行panic
	if level == contract.PanicLevel {
		pkgLog.Panicln(string(ct))
		return nil
	}

	// 通过output进行输出
	log.output.Write(ct)
	log.output.Write([]byte("\r\n"))
	return nil
}

// SetOutput 设置output
func (log *YoimiyaLog) SetOutput(output io.Writer) {
	log.output = output
}

// Panic 输出panic的日志信息
func (log *YoimiyaLog) Panic(ctx context.Context, msg string, fields map[string]interface{}) {
	log.logf(contract.PanicLevel, ctx, msg, fields)
}

// Fatal will add fatal record which contains msg and fields
func (log *YoimiyaLog) Fatal(ctx context.Context, msg string, fields map[string]interface{}) {
	log.logf(contract.FatalLevel, ctx, msg, fields)
}

// Error will add error record which contains msg and fields
func (log *YoimiyaLog) Error(ctx context.Context, msg string, fields map[string]interface{}) {
	log.logf(contract.ErrorLevel, ctx, msg, fields)
}

// Warn will add warn record which contains msg and fields
func (log *YoimiyaLog) Warn(ctx context.Context, msg string, fields map[string]interface{}) {
	log.logf(contract.WarnLevel, ctx, msg, fields)
}

// Info 会打印出普通的日志信息
func (log *YoimiyaLog) Info(ctx context.Context, msg string, fields map[string]interface{}) {
	log.logf(contract.InfoLevel, ctx, msg, fields)
}

// Debug will add debug record which contains msg and fields
func (log *YoimiyaLog) Debug(ctx context.Context, msg string, fields map[string]interface{}) {
	log.logf(contract.DebugLevel, ctx, msg, fields)
}

// Trace will add trace info which contains msg and fields
func (log *YoimiyaLog) Trace(ctx context.Context, msg string, fields map[string]interface{}) {
	log.logf(contract.TraceLevel, ctx, msg, fields)
}

// SetLevel set log level, and higher level will be recorded
func (log *YoimiyaLog) SetLevel(level contract.LogLevel) {
	log.level = level
}

// SetCxtFielder will get fields from context
func (log *YoimiyaLog) SetCtxFielder(handler contract.CtxFielder) {
	log.ctxFielder = handler
}

// SetFormatter will set formatter handler will covert data to string for recording
func (log *YoimiyaLog) SetFormatter(formatter contract.Formatter) {
	log.formatter = formatter
}
