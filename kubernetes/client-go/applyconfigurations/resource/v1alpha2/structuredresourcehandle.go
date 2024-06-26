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

package v1alpha2

import (
	runtime "github.com/fslqd/ck-kube/kubernetes/apimachinery/pkg/runtime"
)

// StructuredResourceHandleApplyConfiguration represents an declarative configuration of the StructuredResourceHandle type for use
// with apply.
type StructuredResourceHandleApplyConfiguration struct {
	VendorClassParameters *runtime.RawExtension                      `json:"vendorClassParameters,omitempty"`
	VendorClaimParameters *runtime.RawExtension                      `json:"vendorClaimParameters,omitempty"`
	NodeName              *string                                    `json:"nodeName,omitempty"`
	Results               []DriverAllocationResultApplyConfiguration `json:"results,omitempty"`
}

// StructuredResourceHandleApplyConfiguration constructs an declarative configuration of the StructuredResourceHandle type for use with
// apply.
func StructuredResourceHandle() *StructuredResourceHandleApplyConfiguration {
	return &StructuredResourceHandleApplyConfiguration{}
}

// WithVendorClassParameters sets the VendorClassParameters field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VendorClassParameters field is set to the value of the last call.
func (b *StructuredResourceHandleApplyConfiguration) WithVendorClassParameters(value runtime.RawExtension) *StructuredResourceHandleApplyConfiguration {
	b.VendorClassParameters = &value
	return b
}

// WithVendorClaimParameters sets the VendorClaimParameters field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VendorClaimParameters field is set to the value of the last call.
func (b *StructuredResourceHandleApplyConfiguration) WithVendorClaimParameters(value runtime.RawExtension) *StructuredResourceHandleApplyConfiguration {
	b.VendorClaimParameters = &value
	return b
}

// WithNodeName sets the NodeName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeName field is set to the value of the last call.
func (b *StructuredResourceHandleApplyConfiguration) WithNodeName(value string) *StructuredResourceHandleApplyConfiguration {
	b.NodeName = &value
	return b
}

// WithResults adds the given value to the Results field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Results field.
func (b *StructuredResourceHandleApplyConfiguration) WithResults(values ...*DriverAllocationResultApplyConfiguration) *StructuredResourceHandleApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithResults")
		}
		b.Results = append(b.Results, *values[i])
	}
	return b
}
