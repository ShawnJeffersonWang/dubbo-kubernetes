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
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"strings"
)

var (
	FeatureInfo = prometheus.NewDesc(
		"horus_feature_info",
		"Indicates the enabled status of specific features in different clusters, with each feature represented by its name, enabled state, and associated cluster name.",
		[]string{
			"feature",
			"enabled",
			"cluster_name",
		},
		nil)

	MultipleInfo = prometheus.NewDesc(
		"horus_multiple_info",
		"Tracks the Prometheus multiple addresses associated with different clusters, providing visibility into the Prometheus endpoints used by each cluster.",
		[]string{
			"cluster_name",
			"prometheus_multiple_address",
		},
		nil)
)

func (h *Horuser) Collect(ch chan<- prometheus.Metric) {
	kFunc := func(m map[string]string) string {
		s := []string{}
		for k := range m {
			s = append(s, k)
		}
		return strings.Join(s, ",")
	}
	info := map[string]string{}
	buttons := map[bool]string{true: "Open", false: "Close"}

	modularKey := fmt.Sprintf("custom modular,%s", buttons[h.cc.CustomModular.Enabled])
	info[modularKey] = kFunc(h.cc.CustomModular.KubeMultiple)
	downtimeKey := fmt.Sprintf("node downtime,%s", buttons[h.cc.NodeDownTime.Enabled])
	info[downtimeKey] = kFunc(h.cc.NodeDownTime.KubeMultiple)

	for k, clusterName := range info {
		s := strings.Split(k, ",")
		feature, enabled := s[0], s[1]
		p := prometheus.MustNewConstMetric(FeatureInfo,
			prometheus.GaugeValue, 1,
			feature,
			enabled,
			clusterName,
		)
		ch <- p
	}
	for clusterName, address := range h.cc.PromMultiple {
		p := prometheus.MustNewConstMetric(MultipleInfo,
			prometheus.GaugeValue, 1,
			clusterName,
			address,
		)
		ch <- p
	}
}

func (h *Horuser) Describe(ch chan<- *prometheus.Desc) {
	ch <- FeatureInfo
	ch <- MultipleInfo
}