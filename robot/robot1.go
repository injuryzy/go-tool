package robot

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type CRobot struct {
	Robot
}

func (C *CRobot) RobotMessage(log *logrus.Entry) string {
	str1 := `{
	"msg_type": "post",
	"content": {
		"post": {
			"zh_cn": {
				"title": "%s",
				"content": [
					[{
							"tag": "text",
							"text": "error message: %s \n"
						},
						{
							"tag": "at",
							"user_id": "%s"
						}]
					]
				}
			}
		}
	}`
	return fmt.Sprintf(str1, log.Message, log.Caller.File, C.User)
}
