业务开发工程中，我们会打印许多日志，如果我们能够将这些错误信息通过机器人发送到群中，那么我们将会很快速的发现问题，所以在这里开发了一个工具

1. 我这里采用的是logrus 这个日志包，为了方便通知每一个开发者，这里可以自己new一个 log对象出来
1. 通知群消息，我这里采用的是日志钩子，直接通过这个钩子将这个消息发送出去
1. 这里的机器人采用的是接口方式，这样我们可以实现自定义的模板

![](https://cdn.nlark.com/yuque/0/2022/jpeg/1535149/1663665802140-5a62f41d-34bf-4f13-9767-7ba853e7e059.jpeg)
<a name="hiDKq"></a>
#### 机器人
```go
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

```
这个机器人是可以自定义的，我们可以用飞书或者是企业微信
<a name="md0ru"></a>
#### 实现logrus的钩子 
如果我们不只是给机器人发，可以在下面钩子事件中继续添加 ，也可以做成责任链模式
```go
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

```
<a name="NrBI1"></a>
#### 添加robothook
```go
func TestLog(t *testing.T) {
	log := NewLog()
	l := NewLog()

	newRobot := robot.NewRobot("", "")
	newRobot1 := robot.CRobot{robot.Robot{
		WebHook: "",
		User:    "",
	}}

	log.AddHook(NewRobotHook(&newRobot, logrus.ErrorLevel))
	l.AddHook(NewRobotHook(&newRobot1, logrus.ErrorLevel))

	log.Error(34214)
	l.Error("faeq")
}

```
在user里面填上需要@的人的id 就可以了 ，具体可以去参考 官方文档 <br />**提示**：[https://github.com/injuryzy/go-tool.git](https://github.com/injuryzy/go-tool.git) 这个源码地址可以去参观一下 ，欢迎大家提issues,如果有问题大家也可以留言一下
