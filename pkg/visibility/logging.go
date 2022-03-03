package visibility

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

func SetLoggingLevel(level string) {
	level = strings.ToLower(level)
	log.WithFields(log.Fields{"level": level}).Trace("setting logging level")
	switch level {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn", "warning":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	default:
		log.WithFields(log.Fields{"level": level}).Warn("Invalid log level, not setting")
	}
}
