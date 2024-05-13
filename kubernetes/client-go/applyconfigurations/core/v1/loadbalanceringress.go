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
	v1 "github.com/fslqd/ck-kube/kubernetes/api/core/v1"
)

// LoadBalancerIngressApplyConfiguration represents an declarative configuration of the LoadBalancerIngress type for use
// with apply.
type LoadBalancerIngressApplyConfiguration struct {
	IP       *string                        `json:"ip,omitempty"`
	Hostname *string                        `json:"hostname,omitempty"`
	IPMode   *v1.LoadBalancerIPMode         `json:"ipMode,omitempty"`
	Ports    []PortStatusApplyConfiguration `json:"ports,omitempty"`
}

// LoadBalancerIngressApplyConfiguration constructs an declarative configuration of the LoadBalancerIngress type for use with
// apply.
func LoadBalancerIngress() *LoadBalancerIngressApplyConfiguration {
	return &LoadBalancerIngressApplyConfiguration{}
}

// WithIP sets the IP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IP field is set to the value of the last call.
func (b *LoadBalancerIngressApplyConfiguration) WithIP(value string) *LoadBalancerIngressApplyConfiguration {
	b.IP = &value
	return b
}

// WithHostname sets the Hostname field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Hostname field is set to the value of the last call.
func (b *LoadBalancerIngressApplyConfiguration) WithHostname(value string) *LoadBalancerIngressApplyConfiguration {
	b.Hostname = &value
	return b
}

// WithIPMode sets the IPMode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IPMode field is set to the value of the last call.
func (b *LoadBalancerIngressApplyConfiguration) WithIPMode(value v1.LoadBalancerIPMode) *LoadBalancerIngressApplyConfiguration {
	b.IPMode = &value
	return b
}

// WithPorts adds the given value to the Ports field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Ports field.
func (b *LoadBalancerIngressApplyConfiguration) WithPorts(values ...*PortStatusApplyConfiguration) *LoadBalancerIngressApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPorts")
		}
		b.Ports = append(b.Ports, *values[i])
	}
	return b
}
