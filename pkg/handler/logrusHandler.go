package handler

import (
	"github.com/sirupsen/logrus"
)

type ChannelHook struct {
	LogChannel chan *logrus.Entry
}

func NewChannelHook(logChannel chan *logrus.Entry) *ChannelHook {
	return &ChannelHook{LogChannel: logChannel}
}

func (hook *ChannelHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *ChannelHook) Fire(entry *logrus.Entry) error {
	hook.LogChannel <- entry
	return nil
}
