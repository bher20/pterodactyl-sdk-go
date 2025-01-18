package utils

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type PlainFormatter struct {
}

func (f *PlainFormatter) Format(entry *log.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("%s\n", entry.Message)), nil
}
func SetLogging(logLevel string) {
	switch logLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
		log.SetFormatter(&log.TextFormatter{})
	case "trace":
		log.SetLevel(log.TraceLevel)
		log.SetFormatter(&log.TextFormatter{
			DisableLevelTruncation: true,
		})
	}
}
