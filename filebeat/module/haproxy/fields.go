// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package haproxy

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "haproxy", asset.ModuleFieldsPri, Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}
