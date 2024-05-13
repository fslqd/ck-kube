/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "ck-kube/kubernetes/api/core/v1"
)

// PortStatusApplyConfiguration represents an declarative configuration of the PortStatus type for use
// with apply.
type PortStatusApplyConfiguration struct {
	Port     *int32       `json:"port,omitempty"`
	Protocol *v1.Protocol `json:"protocol,omitempty"`
	Error    *string      `json:"error,omitempty"`
}

// PortStatusApplyConfiguration constructs an declarative configuration of the PortStatus type for use with
// apply.
func PortStatus() *PortStatusApplyConfiguration {
	return &PortStatusApplyConfiguration{}
}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *PortStatusApplyConfiguration) WithPort(value int32) *PortStatusApplyConfiguration {
	b.Port = &value
	return b
}

// WithProtocol sets the Protocol field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Protocol field is set to the value of the last call.
func (b *PortStatusApplyConfiguration) WithProtocol(value v1.Protocol) *PortStatusApplyConfiguration {
	b.Protocol = &value
	return b
}

// WithError sets the Error field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Error field is set to the value of the last call.
func (b *PortStatusApplyConfiguration) WithError(value string) *PortStatusApplyConfiguration {
	b.Error = &value
	return b
}