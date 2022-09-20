package rlog

import (
	"github.com/injuryzy/go-tool/robot"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLog(t *testing.T) {
	log := NewLog()
	l := NewLog()

	newRobot := robot.NewRobot("https://open.feishu.cn/open-apis/bot/v2/hook/136aee8f-4d5a-4ea5-82a4-e6cdb1295224", "")
	newRobot1 := robot.CRobot{robot.Robot{
		WebHook: "https://open.feishu.cn/open-apis/bot/v2/hook/136aee8f-4d5a-4ea5-82a4-e6cdb1295224",
		User:    "4321",
	}}

	log.AddHook(NewRobotHook(&newRobot, logrus.ErrorLevel))
	l.AddHook(NewRobotHook(&newRobot1, logrus.ErrorLevel))

	log.Error(34214)
	l.Error("faeq")
}
