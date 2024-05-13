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
	v1 "github.com/fslqd/ck-kube/kubernetes/api/core/v1"
	resource "github.com/fslqd/ck-kube/kubernetes/apimachinery/pkg/api/resource"
)

// ResourceMetricStatusApplyConfiguration represents an declarative configuration of the ResourceMetricStatus type for use
// with apply.
type ResourceMetricStatusApplyConfiguration struct {
	Name                      *v1.ResourceName   `json:"name,omitempty"`
	CurrentAverageUtilization *int32             `json:"currentAverageUtilization,omitempty"`
	CurrentAverageValue       *resource.Quantity `json:"currentAverageValue,omitempty"`
}

// ResourceMetricStatusApplyConfiguration constructs an declarative configuration of the ResourceMetricStatus type for use with
// apply.
func ResourceMetricStatus() *ResourceMetricStatusApplyConfiguration {
	return &ResourceMetricStatusApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ResourceMetricStatusApplyConfiguration) WithName(value v1.ResourceName) *ResourceMetricStatusApplyConfiguration {
	b.Name = &value
	return b
}

// WithCurrentAverageUtilization sets the CurrentAverageUtilization field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CurrentAverageUtilization field is set to the value of the last call.
func (b *ResourceMetricStatusApplyConfiguration) WithCurrentAverageUtilization(value int32) *ResourceMetricStatusApplyConfiguration {
	b.CurrentAverageUtilization = &value
	return b
}

// WithCurrentAverageValue sets the CurrentAverageValue field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CurrentAverageValue field is set to the value of the last call.
func (b *ResourceMetricStatusApplyConfiguration) WithCurrentAverageValue(value resource.Quantity) *ResourceMetricStatusApplyConfiguration {
	b.CurrentAverageValue = &value
	return b
}
