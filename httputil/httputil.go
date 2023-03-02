package httputil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func HttpGet(url string, args map[string]string, out interface{}) interface{} {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		logrus.Error("http get fail", err)
	}
	q := req.URL.Query()
	for k, v := range args {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())

	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error("ioutill read fail ", err)
	}
	err = json.Unmarshal(all, out)
	if err != nil {
		logrus.Error("json Unmarshal fail ", err)
	}
	defer resp.Body.Close()

	return out
}

func HttpPost(url string, args map[string]string, out interface{}) interface{} {

	marshal, err2 := json.Marshal(args)
	if err2 != nil {
		logrus.Error("wx phone json marsh err ", err2)
	}

	resp, err := http.Post(url, "application/json;charset=utf-8", bytes.NewBuffer(marshal))

	if err != nil {
		logrus.Error("http post fail", err)
	}
	defer resp.Body.Close()

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}
	err = json.Unmarshal(all, out)
	if err != nil {
		logrus.Error("json Unmarshal fail ", err)
	}

	return out
}
