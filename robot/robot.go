package robot

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

type Roboter interface {
	RobotMessage(log *logrus.Entry) string
	SendMessage(str string)
}
type Robot struct {
	//robot  webhook
	WebHook string
	User    string
}

// NewRobot  generate robot by webhook
func NewRobot(webhook string, user string) Robot {
	if user == "" {
		user = "all"
	}
	return Robot{WebHook: webhook, User: user}
}

// SendMessage robot send message
func (r *Robot) SendMessage(str string) {
	request, err := http.NewRequest("POST", r.WebHook, strings.NewReader(str))
	request.Header = map[string][]string{
		"Content-Type": {"application/json"},
	}
	client := http.DefaultClient
	do, err := client.Do(request)
	if err != nil {
		fmt.Println("sendMessage=", err)
	}
	defer do.Body.Close()
	all, err := ioutil.ReadAll(do.Body)
	if err != nil {
		fmt.Printf("reponse=%s , body = %s ", err, string(all))
	}
}

//RobotMessage rlog convert robot message
func (r *Robot) RobotMessage(log *logrus.Entry) string {
	str1 := `{
	"msg_type": "post",
	"content": {
		"post": {
			"zh_cn": {
				"title": "%s",
				"content": [
					[{
							"tag": "text",
							"text": "error message: %s \nerror line : %d \n"
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
	return fmt.Sprintf(str1, log.Message, log.Caller.File, log.Caller.Line, r.User)
}
