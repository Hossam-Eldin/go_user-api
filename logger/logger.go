package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	//Log : is pointer to zap logger
	Log *zap.Logger
)

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	if Log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}

//Info : this to handle infos
func Info(msg string, tags ...zap.Field) {
	Log.Info(msg, tags...)
	Log.Sync()
}

//Error : this to handle dislaying erros
func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))

	Log.Error(msg, tags...)
	Log.Sync()
}
