package logger

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
)

// Объявление глобальных логгеров
var (
	Info  *log.Logger
	Error *log.Logger
	Warn  *log.Logger
	Debug *log.Logger
)

const (
	LogInfo       = "logs/info.log"
	LogError      = "logs/error.log"
	LogWarning    = "logs/warning.log"
	LogDebug      = "logs/debug.log"
	LogMaxSize    = 25
	LogMaxBackups = 5
	LogMaxAge     = 30
	LogCompress   = true
)

func Init() error {
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err = os.Mkdir("logs", 0755)
		if err != nil {
			return err
		}
	}

	// Инициализация логгеров lumberjack
	lumberLogInfo := &lumberjack.Logger{
		Filename:   LogInfo,
		MaxSize:    LogMaxSize, // мегабайты
		MaxBackups: LogMaxBackups,
		MaxAge:     LogMaxAge,   // дни
		Compress:   LogCompress, // отключено по умолчанию
		LocalTime:  true,
	}

	lumberLogError := &lumberjack.Logger{
		Filename:   LogError,
		MaxSize:    LogMaxSize, // мегабайты
		MaxBackups: LogMaxBackups,
		MaxAge:     LogMaxAge,   // дни
		Compress:   LogCompress, // отключено по умолчанию
		LocalTime:  true,
	}

	lumberLogWarn := &lumberjack.Logger{
		Filename:   LogWarning,
		MaxSize:    LogMaxSize, // мегабайты
		MaxBackups: LogMaxBackups,
		MaxAge:     LogMaxAge,   // дни
		Compress:   LogCompress, // отключено по умолчанию
		LocalTime:  true,
	}

	lumberLogDebug := &lumberjack.Logger{
		Filename:   LogDebug,
		MaxSize:    LogMaxSize, // мегабайты
		MaxBackups: LogMaxBackups,
		MaxAge:     LogMaxAge,   // дни
		Compress:   LogCompress, // отключено по умолчанию
		LocalTime:  true,
	}

	// Инициализация глобальных логгеров
	Info = log.New(gin.DefaultWriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(lumberLogError, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(lumberLogWarn, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(lumberLogDebug, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	gin.DefaultWriter = io.MultiWriter(os.Stdout, lumberLogInfo)

	return nil
}
