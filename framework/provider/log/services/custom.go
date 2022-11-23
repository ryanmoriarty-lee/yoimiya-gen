package services

import (
  "github.com/ryanmoriarty-lee/yoimiya-gen/framework"
  "github.com/ryanmoriarty-lee/yoimiya-gen/framework/contract"
  "io"
)

type YoimiyaCustomLog struct {
  YoimiyaLog
}

func NewYoimiyaCustomLog(params ...interface{}) (interface{}, error) {
  c := params[0].(framework.Container)
  level := params[1].(contract.LogLevel)
  ctxFielder := params[2].(contract.CtxFielder)
  formatter := params[3].(contract.Formatter)
  output := params[4].(io.Writer)

  log := &YoimiyaConsoleLog{}

  log.SetLevel(level)
  log.SetCtxFielder(ctxFielder)
  log.SetFormatter(formatter)

  log.SetOutput(output)
  log.c = c
  return log, nil
}
