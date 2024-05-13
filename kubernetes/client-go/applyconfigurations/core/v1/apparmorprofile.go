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
)

// AppArmorProfileApplyConfiguration represents an declarative configuration of the AppArmorProfile type for use
// with apply.
type AppArmorProfileApplyConfiguration struct {
	Type             *v1.AppArmorProfileType `json:"type,omitempty"`
	LocalhostProfile *string                 `json:"localhostProfile,omitempty"`
}

// AppArmorProfileApplyConfiguration constructs an declarative configuration of the AppArmorProfile type for use with
// apply.
func AppArmorProfile() *AppArmorProfileApplyConfiguration {
	return &AppArmorProfileApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *AppArmorProfileApplyConfiguration) WithType(value v1.AppArmorProfileType) *AppArmorProfileApplyConfiguration {
	b.Type = &value
	return b
}

// WithLocalhostProfile sets the LocalhostProfile field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LocalhostProfile field is set to the value of the last call.
func (b *AppArmorProfileApplyConfiguration) WithLocalhostProfile(value string) *AppArmorProfileApplyConfiguration {
	b.LocalhostProfile = &value
	return b
}
