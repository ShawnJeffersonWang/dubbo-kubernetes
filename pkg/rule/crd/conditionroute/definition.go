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

package conditionroute

type Policy struct {
	Name string `json:"name,omitempty"`

	Spec *PolicySpec `json:"spec"`
}

func (p *Policy) CopyToClient() *PolicyToClient {
	toClient := &PolicyToClient{
		Name: p.Name,
	}

	if p.Spec != nil {
		toClient.Spec = p.Spec.CopyToClient()
	}

	return toClient
}

type PolicySpec struct {
	Priority      int      `json:"priority" yaml:"priority,omitempty"`
	Enabled       bool     `json:"enabled" yaml:"enabled"`
	Force         bool     `json:"force" yaml:"force"`
	Runtime       bool     `json:"runtime" yaml:"runtime"`
	Key           string   `json:"key" yaml:"key"`
	Scope         string   `json:"scope" yaml:"scope"`
	Conditions    []string `json:"conditions" yaml:"conditions"`
	ConfigVersion string   `json:"configVersion" yaml:"configVersion"`
}

func (p *PolicySpec) CopyToClient() *PolicySpecToClient {
	toClient := &PolicySpecToClient{
		Priority:      p.Priority,
		Enabled:       p.Enabled,
		Force:         p.Force,
		Runtime:       p.Runtime,
		Key:           p.Key,
		Scope:         p.Scope,
		Conditions:    p.Conditions,
		ConfigVersion: p.ConfigVersion,
	}
	return toClient
}

// To Client Rule

type PolicyToClient struct {
	Name string `json:"name,omitempty"`

	Spec *PolicySpecToClient `json:"spec"`
}

type PolicySpecToClient struct {
	Priority      int      `json:"priority" yaml:"priority,omitempty"`
	Enabled       bool     `json:"enabled" yaml:"enabled"`
	Force         bool     `json:"force" yaml:"force"`
	Runtime       bool     `json:"runtime" yaml:"runtime"`
	Key           string   `json:"key" yaml:"key"`
	Scope         string   `json:"scope" yaml:"scope"`
	Conditions    []string `json:"conditions" yaml:"conditions"`
	ConfigVersion string   `json:"configVersion" yaml:"configVersion"`
}
