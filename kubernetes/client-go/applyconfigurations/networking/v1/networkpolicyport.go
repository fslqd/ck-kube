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
	intstr "ck-kube/kubernetes/apimachinery/pkg/util/intstr"
)

// NetworkPolicyPortApplyConfiguration represents an declarative configuration of the NetworkPolicyPort type for use
// with apply.
type NetworkPolicyPortApplyConfiguration struct {
	Protocol *v1.Protocol        `json:"protocol,omitempty"`
	Port     *intstr.IntOrString `json:"port,omitempty"`
	EndPort  *int32              `json:"endPort,omitempty"`
}

// NetworkPolicyPortApplyConfiguration constructs an declarative configuration of the NetworkPolicyPort type for use with
// apply.
func NetworkPolicyPort() *NetworkPolicyPortApplyConfiguration {
	return &NetworkPolicyPortApplyConfiguration{}
}

// WithProtocol sets the Protocol field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Protocol field is set to the value of the last call.
func (b *NetworkPolicyPortApplyConfiguration) WithProtocol(value v1.Protocol) *NetworkPolicyPortApplyConfiguration {
	b.Protocol = &value
	return b
}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *NetworkPolicyPortApplyConfiguration) WithPort(value intstr.IntOrString) *NetworkPolicyPortApplyConfiguration {
	b.Port = &value
	return b
}

// WithEndPort sets the EndPort field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EndPort field is set to the value of the last call.
func (b *NetworkPolicyPortApplyConfiguration) WithEndPort(value int32) *NetworkPolicyPortApplyConfiguration {
	b.EndPort = &value
	return b
}