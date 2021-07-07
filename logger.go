package gormlogrus

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type Config struct {
	SlowThreshold             time.Duration
	IgnoreRecordNotFoundError bool
	LogLevel                  logger.LogLevel
	Output                    io.Writer
	modName, hostName         string
}

func New(writer logger.Writer, config Config) logger.Interface {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	})

	logger.SetOutput(config.Output)
	logger.SetLevel(logrus.DebugLevel)
	entry := logger.WithFields(logrus.Fields{
		"module": config.modName,
		"host":   config.hostName,
	})

	l := &gormlogger{
		Writer:   writer,
		Config:   config,
		logEntry: entry,
	}

	return l
}

type gormlogger struct {
	logger.Writer
	Config

	logEntry *logrus.Entry
}

// LogMode log mode
func (l *gormlogger) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

// Info print info
func (l gormlogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		l.logEntry.WithFields(logrus.Fields{
			"eventName": "MySQL_Info",
			"eventId":   uuid.New().String(),
			"event":     logrus.Fields{},
		}).Info(msg)
	}
}

// Warn print warn messages
func (l gormlogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		l.logEntry.WithFields(logrus.Fields{
			"eventName": "MySQL_Warn",
			"eventId":   uuid.New().String(),
			"event":     logrus.Fields{},
		}).Warn(msg)
	}
}

// Error print error messages
func (l gormlogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		l.logEntry.WithFields(logrus.Fields{
			"eventName": "MySQL_Error",
			"eventId":   uuid.New().String(),
			"event":     logrus.Fields{},
		}).Error(msg)
	}
}

// Trace print sql message
func (l gormlogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {

	if l.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= logger.Error && (!errors.Is(err, gorm.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			l.Printf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Printf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case l.LogLevel == logger.Info:
		sql, rows := fc()
		if rows == -1 {
			l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
