// Copyright 2019-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package integration

import (
	"github.com/onosproject/onos-config/pkg/utils"
	"strings"
)

// MakeProtoTarget returns a GNMI proto path for a given target
func MakeProtoTarget(target string, path string) string {
	var protoBuilder strings.Builder

	protoBuilder.WriteString("<target: '")
	protoBuilder.WriteString(target)
	protoBuilder.WriteString("', ")

	pathElements := utils.SplitPath(path)
	for _, pathElement := range pathElements {
		protoBuilder.WriteString("elem: <name: '")
		protoBuilder.WriteString(pathElement)
		protoBuilder.WriteString("'>")
	}
	protoBuilder.WriteString(">")
	return protoBuilder.String()
}

// MakeProtoPath returns a path: element for a given target and path
func MakeProtoPath(target string, path string) string {
	var protoBuilder strings.Builder

	protoBuilder.WriteString("path: ")
	gnmiPath := MakeProtoTarget(target, path)
	protoBuilder.WriteString(gnmiPath)
	return protoBuilder.String()
}

// MakeProtoUpdatePath returns an update: element for a target, path, and new value
func MakeProtoUpdatePath(target string, path string, value string) string {
	var protoBuilder strings.Builder

	protoBuilder.WriteString("update: <")
	protoBuilder.WriteString(MakeProtoPath(target, path))
	protoBuilder.WriteString(" val: <string_val: '")
	protoBuilder.WriteString(value)
	protoBuilder.WriteString("'>")
	protoBuilder.WriteString(">")
	return protoBuilder.String()
}

// MakeProtoDeletePath returns a delete: element for a given target and path
func MakeProtoDeletePath(target string, path string) string {
	var protoBuilder strings.Builder

	protoBuilder.WriteString("delete: ")
	protoBuilder.WriteString(MakeProtoTarget(target, path))
	return protoBuilder.String()
}
