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

package v2beta1

import (
	v1 "ck-kube/kubernetes/api/core/v1"
	resource "ck-kube/kubernetes/apimachinery/pkg/api/resource"
)

// ContainerResourceMetricSourceApplyConfiguration represents an declarative configuration of the ContainerResourceMetricSource type for use
// with apply.
type ContainerResourceMetricSourceApplyConfiguration struct {
	Name                     *v1.ResourceName   `json:"name,omitempty"`
	TargetAverageUtilization *int32             `json:"targetAverageUtilization,omitempty"`
	TargetAverageValue       *resource.Quantity `json:"targetAverageValue,omitempty"`
	Container                *string            `json:"container,omitempty"`
}

// ContainerResourceMetricSourceApplyConfiguration constructs an declarative configuration of the ContainerResourceMetricSource type for use with
// apply.
func ContainerResourceMetricSource() *ContainerResourceMetricSourceApplyConfiguration {
	return &ContainerResourceMetricSourceApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ContainerResourceMetricSourceApplyConfiguration) WithName(value v1.ResourceName) *ContainerResourceMetricSourceApplyConfiguration {
	b.Name = &value
	return b
}

// WithTargetAverageUtilization sets the TargetAverageUtilization field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TargetAverageUtilization field is set to the value of the last call.
func (b *ContainerResourceMetricSourceApplyConfiguration) WithTargetAverageUtilization(value int32) *ContainerResourceMetricSourceApplyConfiguration {
	b.TargetAverageUtilization = &value
	return b
}

// WithTargetAverageValue sets the TargetAverageValue field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TargetAverageValue field is set to the value of the last call.
func (b *ContainerResourceMetricSourceApplyConfiguration) WithTargetAverageValue(value resource.Quantity) *ContainerResourceMetricSourceApplyConfiguration {
	b.TargetAverageValue = &value
	return b
}

// WithContainer sets the Container field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Container field is set to the value of the last call.
func (b *ContainerResourceMetricSourceApplyConfiguration) WithContainer(value string) *ContainerResourceMetricSourceApplyConfiguration {
	b.Container = &value
	return b
}
