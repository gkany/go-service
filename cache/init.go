package cache

import (
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/liyue201/go-logger"
	"go-service/common"
	"go-service/models"
	"os"
)

func InitConfig(filePath string) error {
	err := configor.Load(&Cfg, filePath)
	//logger.Debugf("[Init] config: %#v\n", Cfg)
	return err
}

func InitLog(config models.LogConfig) error {
	if config.File != "" && config.OutputFile {
		logPath := common.AbsPath(config.Path)
		logFile := config.File
		if !common.PathExist(logPath) {
			err := os.MkdirAll(logPath, os.ModePerm)
			if err != nil {
				fmt.Println("logger init error, make dir failed. err=", err)
				return err
			}
		}
		err := logger.InitFileLog(logPath, logFile, config.Level, true, false)
		if err != nil {
			fmt.Println("logger init error, InitFileLog failed. err=", err)
			return err
		}
	}
	logger.InitStdOutput(config.OutputConsole, config.Level, true)

	fmt.Println("logger init done")
	return nil
}

