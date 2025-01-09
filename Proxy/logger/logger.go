package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log zap.Logger

func CreateLogger() {
	consoleEncoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:      "time",
		LevelKey:     "level",
		MessageKey:   "msg",
		CallerKey:    "caller",
		EncodeLevel:  zapcore.CapitalColorLevelEncoder, // Colorful output
		EncodeTime:   zapcore.ISO8601TimeEncoder,       // ISO8601 format
		EncodeCaller: zapcore.ShortCallerEncoder,
	})

	// Create JSON encoder for file output
	fileEncoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:      "time",
		LevelKey:     "level",
		MessageKey:   "msg",
		CallerKey:    "caller",
		EncodeLevel:  zapcore.LowercaseLevelEncoder, // Lowercase for JSON
		EncodeTime:   zapcore.ISO8601TimeEncoder,    // ISO8601 format
		EncodeCaller: zapcore.ShortCallerEncoder,
	})

	// Open the log file for writing
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	fileSyncer := zapcore.AddSync(file)

	// Create core that combines console and file outputs
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
		zapcore.NewCore(fileEncoder, fileSyncer, zapcore.DebugLevel),
	)

	// Create the logger with the combined core
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	defer logger.Sync()

	log = *logger
}

func GetLogger() zap.Logger {
	return log
}
