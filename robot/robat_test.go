package robot

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

var str = "https://open.feishu.cn/open-apis/bot/v2/hook/136aee8f-4d5a-4ea5-82a4-e6cdb1295224"

func TestRobot(t *testing.T) {

	str1 := `{
	"msg_type": "post",
	"content": {
		"post": {
			"zh_cn": {
				"title": "项目更新通知",
				"content": [
					[{
							"tag": "text",
							"text": "项目有更新: "
						},
						{
							"tag": "a",
							"text": "请查看",
							"href": "http://www.example.com/"
						},
						{
							"tag": "at",
							"user_id": "all"
						}
					]
				]
			}
		}
	}
}`
	request, err := http.NewRequest("POST", str, strings.NewReader(str1))
	if err != nil {
		t.Log(err)
	}
	request.Header = map[string][]string{
		"Content-Type": {"application/json"},
	}

	client := http.DefaultClient
	do, err := client.Do(request)
	if err != nil {
		t.Log(err)
	}
	defer do.Body.Close()
	all, err := ioutil.ReadAll(do.Body)
	if err != nil {
		t.Log(err)
	}
	fmt.Println(string(all))
}

func TestName(t *testing.T) {

}
