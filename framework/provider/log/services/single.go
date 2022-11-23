package services

import (
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework"
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework/contract"
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework/util"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type YoimiyaSingleLog struct {
	YoimiyaLog

	folder string
	file   string
	fd     *os.File
}

// NewYoimiyaSingleLog params sequence: level, ctxFielder, Formatter, map[string]interface(folder/file)
func NewYoimiyaSingleLog(params ...interface{}) (interface{}, error) {
	c := params[0].(framework.Container)
	level := params[1].(contract.LogLevel)
	ctxFielder := params[2].(contract.CtxFielder)
	formatter := params[3].(contract.Formatter)

	appService := c.MustMake(contract.AppKey).(contract.App)
	configService := c.MustMake(contract.ConfigKey).(contract.Config)

	log := &YoimiyaSingleLog{}
	log.SetLevel(level)
	log.SetCtxFielder(ctxFielder)
	log.SetFormatter(formatter)

	folder := appService.LogFolder()
	if configService.IsExist("log.folder") {
		folder = configService.GetString("log.folder")
	}
	log.folder = folder
	if !util.Exists(folder) {
		os.MkdirAll(folder, os.ModePerm)
	}

	log.file = "yoimiya.log"
	if configService.IsExist("log.file") {
		log.file = configService.GetString("log.file")
	}

	fd, err := os.OpenFile(filepath.Join(log.folder, log.file), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, errors.Wrap(err, "open log file err")
	}

	log.SetOutput(fd)
	log.c = c

	return log, nil
}
