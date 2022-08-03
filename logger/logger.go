package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

<<<<<<< HEAD
var Log *zap.Logger
=======
var log *zap.Logger

>>>>>>> 933acab27e67785d8e63b2f07537dd6e22bf744f
func init() {

	config := zap.NewProductionConfig()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""

	config.EncoderConfig = encoderConfig

<<<<<<< HEAD

	var err error 
	Log, err =  config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic (err)
	}

}


func Info(message string, fields ...zapcore.Field){
	Log.Info(message, fields...)
}

func Debug(message string, fields ...zapcore.Field){
	Log.Debug(message, fields...)
}

func Error(message string, fields ...zapcore.Field){
	Log.Error(message, fields...)
}

func Fatal(message string, fields ...zapcore.Field){
	Log.Fatal(message, fields...)
}
=======
	var err error
	// log, err = zap.NewProduction(zap.AddCallerSkip(1))
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zapcore.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zapcore.Field) {
	log.Debug(message, fields...)
}

func Error(message string, fields ...zapcore.Field) {
	log.Error(message, fields...)
}

func Fatal(message string, fields ...zapcore.Field) {
	log.Fatal(message, fields...)
}
>>>>>>> 933acab27e67785d8e63b2f07537dd6e22bf744f
