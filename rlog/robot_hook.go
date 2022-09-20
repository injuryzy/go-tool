package rlog

import (
	"github.com/injuryzy/go-tool/robot"
	"github.com/sirupsen/logrus"
)

type RobotHook struct {
	LogLevels []logrus.Level
	Robot     robot.Roboter
}

//NewRobotHook generate hook add robot
func NewRobotHook(robot robot.Roboter, arg ...logrus.Level) *RobotHook {
	r := new(RobotHook)
	r.Robot = robot
	for _, level := range arg {
		r.LogLevels = append(r.LogLevels, level)
	}
	return r
}

func (r *RobotHook) Levels() []logrus.Level {
	return r.LogLevels
}

func (r *RobotHook) Fire(entry *logrus.Entry) error {
	r.Robot.SendMessage(r.Robot.RobotMessage(entry))
	return nil
}
