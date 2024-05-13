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

package v1beta1

import (
	intstr "ck-kube/kubernetes/apimachinery/pkg/util/intstr"
)

// RollingUpdateDaemonSetApplyConfiguration represents an declarative configuration of the RollingUpdateDaemonSet type for use
// with apply.
type RollingUpdateDaemonSetApplyConfiguration struct {
	MaxUnavailable *intstr.IntOrString `json:"maxUnavailable,omitempty"`
	MaxSurge       *intstr.IntOrString `json:"maxSurge,omitempty"`
}

// RollingUpdateDaemonSetApplyConfiguration constructs an declarative configuration of the RollingUpdateDaemonSet type for use with
// apply.
func RollingUpdateDaemonSet() *RollingUpdateDaemonSetApplyConfiguration {
	return &RollingUpdateDaemonSetApplyConfiguration{}
}

// WithMaxUnavailable sets the MaxUnavailable field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaxUnavailable field is set to the value of the last call.
func (b *RollingUpdateDaemonSetApplyConfiguration) WithMaxUnavailable(value intstr.IntOrString) *RollingUpdateDaemonSetApplyConfiguration {
	b.MaxUnavailable = &value
	return b
}

// WithMaxSurge sets the MaxSurge field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaxSurge field is set to the value of the last call.
func (b *RollingUpdateDaemonSetApplyConfiguration) WithMaxSurge(value intstr.IntOrString) *RollingUpdateDaemonSetApplyConfiguration {
	b.MaxSurge = &value
	return b
}
