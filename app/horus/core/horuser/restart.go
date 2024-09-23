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

package horuser

import (
	"context"
	"fmt"
	"github.com/apache/dubbo-kubernetes/app/horus/basic/db"
	"github.com/apache/dubbo-kubernetes/app/horus/core/alert"
	"github.com/gammazero/workerpool"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
	"time"
)

func (h *Horuser) DowntimeRestartManager(ctx context.Context) error {
	go wait.UntilWithContext(ctx, h.RestartOrRepair, time.Duration(h.cc.NodeDownTime.CheckIntervalSecond)*time.Second)
	<-ctx.Done()
	return nil
}

func (h *Horuser) RestartOrRepair(ctx context.Context) {
	nodes, err := db.GetRestartNodeDataInfoDate()
	if err != nil {
		klog.Errorf("RestartOrRepair err:%v", err)
	}
	if len(nodes) == 0 {
		klog.Warningf("Needs to be rebooted or fixed to zero.")
	}
	klog.Infof("GetRestartNodeDataInfoDate.count:%v", len(nodes))
	wp := workerpool.New(10)
	for _, n := range nodes {
		n := n
		wp.Submit(func() {
			h.TryRestart(n)
		})
	}
}

func (h *Horuser) TryRestart(node db.NodeDataInfo) {
	msg := fmt.Sprintf("node restart or force restart.node:%v date:%v cluster:%v", node.NodeName, node.FirstDate, node.ClusterName)
	alert.DingTalkSend(h.cc.DingTalk, msg)
	node.RestartMarker()
}