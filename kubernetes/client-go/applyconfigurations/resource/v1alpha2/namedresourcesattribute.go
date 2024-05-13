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
	resource "ck-kube/kubernetes/apimachinery/pkg/api/resource"
)

// NamedResourcesAttributeApplyConfiguration represents an declarative configuration of the NamedResourcesAttribute type for use
// with apply.
type NamedResourcesAttributeApplyConfiguration struct {
	Name                                           *string `json:"name,omitempty"`
	NamedResourcesAttributeValueApplyConfiguration `json:",inline"`
}

// NamedResourcesAttributeApplyConfiguration constructs an declarative configuration of the NamedResourcesAttribute type for use with
// apply.
func NamedResourcesAttribute() *NamedResourcesAttributeApplyConfiguration {
	return &NamedResourcesAttributeApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *NamedResourcesAttributeApplyConfiguration) WithName(value string) *NamedResourcesAttributeApplyConfiguration {
	b.Name = &value
	return b
}

// WithQuantityValue sets the QuantityValue field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the QuantityValue field is set to the value of the last call.
func (b *NamedResourcesAttributeApplyConfiguration) WithQuantityValue(value resource.Quantity) *NamedResourcesAttributeApplyConfiguration {
	b.QuantityValue = &value
	return b
}

// WithBoolValue sets the BoolValue field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BoolValue field is set to the value of the last call.
func (b *NamedResourcesAttributeApplyConfiguration) WithBoolValue(value bool) *NamedResourcesAttributeApplyConfiguration {
	b.BoolValue = &value
	return b
}

// WithIntValue sets the IntValue field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IntValue field is set to the value of the last call.
func (b *NamedResourcesAttributeApplyConfiguration) WithIntValue(value int64) *NamedResourcesAttributeApplyConfiguration {
	b.IntValue = &value
	return b
}

// WithIntSliceValue sets the IntSliceValue field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IntSliceValue field is set to the value of the last call.
func (b *NamedResourcesAttributeApplyConfiguration) WithIntSliceValue(value *NamedResourcesIntSliceApplyConfiguration) *NamedResourcesAttributeApplyConfiguration {
	b.IntSliceValue = value
	return b
}

// WithStringValue sets the StringValue field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StringValue field is set to the value of the last call.
func (b *NamedResourcesAttributeApplyConfiguration) WithStringValue(value string) *NamedResourcesAttributeApplyConfiguration {
	b.StringValue = &value
	return b
}

// WithStringSliceValue sets the StringSliceValue field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StringSliceValue field is set to the value of the last call.
func (b *NamedResourcesAttributeApplyConfiguration) WithStringSliceValue(value *NamedResourcesStringSliceApplyConfiguration) *NamedResourcesAttributeApplyConfiguration {
	b.StringSliceValue = value
	return b
}

// WithVersionValue sets the VersionValue field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VersionValue field is set to the value of the last call.
func (b *NamedResourcesAttributeApplyConfiguration) WithVersionValue(value string) *NamedResourcesAttributeApplyConfiguration {
	b.VersionValue = &value
	return b
}
