// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package alert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/apache/dubbo-kubernetes/app/horus/basic/config"
	"k8s.io/klog/v2"
	"net/http"
)

const Title = "钉钉机器人"

type content struct {
	Content string `json:"content"`
}

type at struct {
	AtMobiles []string `json:"atMobiles"`
}

type Message struct {
	MessageType string  `json:"messageType"`
	Text        content `json:"text"`
	At          at      `json:"at"`
}

type T struct {
	At struct {
		AtMobiles []string `json:"atMobiles"`
		AtUserIds []string `json:"atUserIds"`
		IsAtAll   bool     `json:"isAtAll"`
	} `json:"at"`
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
	Msgtype string `json:"msgtype"`
}

func DingTalkSend(dk *config.DingTalkConfiguration, msg string) {
	dtm := Message{MessageType: "text"}
	dtm.Text.Content = fmt.Sprintf("%s\n"+
		"【日志详细信息：%s】", Title, msg)
	dtm.At.AtMobiles = dk.AtMobiles
	bs, err := json.Marshal(dtm)
	if err != nil {
		klog.Errorf("dingTalk json marshal err:%v\n msg:%v\n", err, msg)
		return
	}
	res, err := http.Post(dk.WebhookUrl, "application/json", bytes.NewBuffer(bs))
	if err != nil {
		klog.Errorf("push dingTalk err:%v\n msg:%v\n", err, msg)
	}
	if res.StatusCode != 200 && res != nil {
		klog.Errorf("push dingTalk status code err:%v\n code:%v\n msg:%v\n", err, res.StatusCode, msg)
		return
	}
	klog.Infof("push dingTalk success code:%v\n msg:%v\n", res.StatusCode, msg)
}