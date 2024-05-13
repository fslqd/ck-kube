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
	v1 "github.com/fslqd/ck-kube/kubernetes/client-go/applyconfigurations/core/v1"
)

// EndpointApplyConfiguration represents an declarative configuration of the Endpoint type for use
// with apply.
type EndpointApplyConfiguration struct {
	Addresses  []string                              `json:"addresses,omitempty"`
	Conditions *EndpointConditionsApplyConfiguration `json:"conditions,omitempty"`
	Hostname   *string                               `json:"hostname,omitempty"`
	TargetRef  *v1.ObjectReferenceApplyConfiguration `json:"targetRef,omitempty"`
	Topology   map[string]string                     `json:"topology,omitempty"`
	NodeName   *string                               `json:"nodeName,omitempty"`
	Hints      *EndpointHintsApplyConfiguration      `json:"hints,omitempty"`
}

// EndpointApplyConfiguration constructs an declarative configuration of the Endpoint type for use with
// apply.
func Endpoint() *EndpointApplyConfiguration {
	return &EndpointApplyConfiguration{}
}

// WithAddresses adds the given value to the Addresses field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Addresses field.
func (b *EndpointApplyConfiguration) WithAddresses(values ...string) *EndpointApplyConfiguration {
	for i := range values {
		b.Addresses = append(b.Addresses, values[i])
	}
	return b
}

// WithConditions sets the Conditions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Conditions field is set to the value of the last call.
func (b *EndpointApplyConfiguration) WithConditions(value *EndpointConditionsApplyConfiguration) *EndpointApplyConfiguration {
	b.Conditions = value
	return b
}

// WithHostname sets the Hostname field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Hostname field is set to the value of the last call.
func (b *EndpointApplyConfiguration) WithHostname(value string) *EndpointApplyConfiguration {
	b.Hostname = &value
	return b
}

// WithTargetRef sets the TargetRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TargetRef field is set to the value of the last call.
func (b *EndpointApplyConfiguration) WithTargetRef(value *v1.ObjectReferenceApplyConfiguration) *EndpointApplyConfiguration {
	b.TargetRef = value
	return b
}

// WithTopology puts the entries into the Topology field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Topology field,
// overwriting an existing map entries in Topology field with the same key.
func (b *EndpointApplyConfiguration) WithTopology(entries map[string]string) *EndpointApplyConfiguration {
	if b.Topology == nil && len(entries) > 0 {
		b.Topology = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Topology[k] = v
	}
	return b
}

// WithNodeName sets the NodeName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeName field is set to the value of the last call.
func (b *EndpointApplyConfiguration) WithNodeName(value string) *EndpointApplyConfiguration {
	b.NodeName = &value
	return b
}

// WithHints sets the Hints field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Hints field is set to the value of the last call.
func (b *EndpointApplyConfiguration) WithHints(value *EndpointHintsApplyConfiguration) *EndpointApplyConfiguration {
	b.Hints = value
	return b
}
