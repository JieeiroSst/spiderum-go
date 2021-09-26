package log

import (
	"github.com/hhkbp2/go-logging"
	"os"
	"sync"
	"time"
)

type Log struct {
	log logging.Logger
}

var (
	instance *Log
	once sync.Once
	filePath = "./data.log"
	fileMode = os.O_APPEND
	bufferSize = 0
	bufferFlushTime = 30 * time.Second
	inputChanSize = 1
	fileMaxBytes = uint64(100 * 1024 * 1024)
	backupCount = uint32(9)
)

func GetLogInstance() *Log {
	once.Do(func() {
		handler := logging.MustNewRotatingFileHandler(filePath, fileMode, bufferSize, bufferFlushTime, inputChanSize, fileMaxBytes, backupCount)
		format := "%(asctime)s %(levelname)s (%(filename)s:%(lineno)d) %(name)s %(message)s"
		dateFormat := "%Y-%m-%d %H:%M:%S.%3n"
		formatter := logging.NewStandardFormatter(format, dateFormat)
		handler.SetFormatter(formatter)
		logger := logging.GetLogger("a.b.c")
		logger.SetLevel(logging.LevelInfo)
		logger.AddHandler(handler)
		defer logging.Shutdown()

		instance = &Log{log:logger}
	})
	return instance
}

func NewLog() logging.Logger{
	return GetLogInstance().log
}
