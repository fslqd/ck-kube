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
	v1 "ck-kube/kubernetes/client-go/applyconfigurations/core/v1"
)

// IngressBackendApplyConfiguration represents an declarative configuration of the IngressBackend type for use
// with apply.
type IngressBackendApplyConfiguration struct {
	ServiceName *string                                         `json:"serviceName,omitempty"`
	ServicePort *intstr.IntOrString                             `json:"servicePort,omitempty"`
	Resource    *v1.TypedLocalObjectReferenceApplyConfiguration `json:"resource,omitempty"`
}

// IngressBackendApplyConfiguration constructs an declarative configuration of the IngressBackend type for use with
// apply.
func IngressBackend() *IngressBackendApplyConfiguration {
	return &IngressBackendApplyConfiguration{}
}

// WithServiceName sets the ServiceName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServiceName field is set to the value of the last call.
func (b *IngressBackendApplyConfiguration) WithServiceName(value string) *IngressBackendApplyConfiguration {
	b.ServiceName = &value
	return b
}

// WithServicePort sets the ServicePort field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServicePort field is set to the value of the last call.
func (b *IngressBackendApplyConfiguration) WithServicePort(value intstr.IntOrString) *IngressBackendApplyConfiguration {
	b.ServicePort = &value
	return b
}

// WithResource sets the Resource field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resource field is set to the value of the last call.
func (b *IngressBackendApplyConfiguration) WithResource(value *v1.TypedLocalObjectReferenceApplyConfiguration) *IngressBackendApplyConfiguration {
	b.Resource = value
	return b
}
