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

package v2

import (
	v1 "ck-kube/kubernetes/apimachinery/pkg/apis/meta/v1"
)

// HorizontalPodAutoscalerStatusApplyConfiguration represents an declarative configuration of the HorizontalPodAutoscalerStatus type for use
// with apply.
type HorizontalPodAutoscalerStatusApplyConfiguration struct {
	ObservedGeneration *int64                                               `json:"observedGeneration,omitempty"`
	LastScaleTime      *v1.Time                                             `json:"lastScaleTime,omitempty"`
	CurrentReplicas    *int32                                               `json:"currentReplicas,omitempty"`
	DesiredReplicas    *int32                                               `json:"desiredReplicas,omitempty"`
	CurrentMetrics     []MetricStatusApplyConfiguration                     `json:"currentMetrics,omitempty"`
	Conditions         []HorizontalPodAutoscalerConditionApplyConfiguration `json:"conditions,omitempty"`
}

// HorizontalPodAutoscalerStatusApplyConfiguration constructs an declarative configuration of the HorizontalPodAutoscalerStatus type for use with
// apply.
func HorizontalPodAutoscalerStatus() *HorizontalPodAutoscalerStatusApplyConfiguration {
	return &HorizontalPodAutoscalerStatusApplyConfiguration{}
}

// WithObservedGeneration sets the ObservedGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedGeneration field is set to the value of the last call.
func (b *HorizontalPodAutoscalerStatusApplyConfiguration) WithObservedGeneration(value int64) *HorizontalPodAutoscalerStatusApplyConfiguration {
	b.ObservedGeneration = &value
	return b
}

// WithLastScaleTime sets the LastScaleTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastScaleTime field is set to the value of the last call.
func (b *HorizontalPodAutoscalerStatusApplyConfiguration) WithLastScaleTime(value v1.Time) *HorizontalPodAutoscalerStatusApplyConfiguration {
	b.LastScaleTime = &value
	return b
}

// WithCurrentReplicas sets the CurrentReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CurrentReplicas field is set to the value of the last call.
func (b *HorizontalPodAutoscalerStatusApplyConfiguration) WithCurrentReplicas(value int32) *HorizontalPodAutoscalerStatusApplyConfiguration {
	b.CurrentReplicas = &value
	return b
}

// WithDesiredReplicas sets the DesiredReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DesiredReplicas field is set to the value of the last call.
func (b *HorizontalPodAutoscalerStatusApplyConfiguration) WithDesiredReplicas(value int32) *HorizontalPodAutoscalerStatusApplyConfiguration {
	b.DesiredReplicas = &value
	return b
}

// WithCurrentMetrics adds the given value to the CurrentMetrics field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the CurrentMetrics field.
func (b *HorizontalPodAutoscalerStatusApplyConfiguration) WithCurrentMetrics(values ...*MetricStatusApplyConfiguration) *HorizontalPodAutoscalerStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithCurrentMetrics")
		}
		b.CurrentMetrics = append(b.CurrentMetrics, *values[i])
	}
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *HorizontalPodAutoscalerStatusApplyConfiguration) WithConditions(values ...*HorizontalPodAutoscalerConditionApplyConfiguration) *HorizontalPodAutoscalerStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}
