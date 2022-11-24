package middleware

import (
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework/contract"
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework/gin"
)

// Trace ...
type Trace struct{}

// NewTrace ...
func NewTrace() *Trace {
	return &Trace{}
}

// Func ...
func (t *Trace) Func() gin.HandlerFunc {
	// 使用函数回调
	return func(c *gin.Context) {
		// 记录开始时间
		tracer := c.MustMake(contract.TraceKey).(contract.Trace)
		traceCtx := tracer.ExtractHTTP(c.Request)

		tracer.WithTrace(c, traceCtx)

		// 使用next执行具体的业务逻辑
		c.Next()
	}
}
