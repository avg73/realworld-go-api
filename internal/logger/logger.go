package logger

import "github.com/sirupsen/logrus"

func ConfigureLogger(l *logrus.Logger, lvl string) error {
	level, err := logrus.ParseLevel(lvl)
	if err != nil {
		return err
	}
	l.SetLevel(level)
	return nil
}
