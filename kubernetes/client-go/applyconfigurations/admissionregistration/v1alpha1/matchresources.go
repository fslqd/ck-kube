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

package v1alpha1

import (
	admissionregistrationv1alpha1 "github.com/fslqd/ck-kube/kubernetes/api/admissionregistration/v1alpha1"
	v1 "github.com/fslqd/ck-kube/kubernetes/client-go/applyconfigurations/meta/v1"
)

// MatchResourcesApplyConfiguration represents an declarative configuration of the MatchResources type for use
// with apply.
type MatchResourcesApplyConfiguration struct {
	NamespaceSelector    *v1.LabelSelectorApplyConfiguration            `json:"namespaceSelector,omitempty"`
	ObjectSelector       *v1.LabelSelectorApplyConfiguration            `json:"objectSelector,omitempty"`
	ResourceRules        []NamedRuleWithOperationsApplyConfiguration    `json:"resourceRules,omitempty"`
	ExcludeResourceRules []NamedRuleWithOperationsApplyConfiguration    `json:"excludeResourceRules,omitempty"`
	MatchPolicy          *admissionregistrationv1alpha1.MatchPolicyType `json:"matchPolicy,omitempty"`
}

// MatchResourcesApplyConfiguration constructs an declarative configuration of the MatchResources type for use with
// apply.
func MatchResources() *MatchResourcesApplyConfiguration {
	return &MatchResourcesApplyConfiguration{}
}

// WithNamespaceSelector sets the NamespaceSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NamespaceSelector field is set to the value of the last call.
func (b *MatchResourcesApplyConfiguration) WithNamespaceSelector(value *v1.LabelSelectorApplyConfiguration) *MatchResourcesApplyConfiguration {
	b.NamespaceSelector = value
	return b
}

// WithObjectSelector sets the ObjectSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObjectSelector field is set to the value of the last call.
func (b *MatchResourcesApplyConfiguration) WithObjectSelector(value *v1.LabelSelectorApplyConfiguration) *MatchResourcesApplyConfiguration {
	b.ObjectSelector = value
	return b
}

// WithResourceRules adds the given value to the ResourceRules field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ResourceRules field.
func (b *MatchResourcesApplyConfiguration) WithResourceRules(values ...*NamedRuleWithOperationsApplyConfiguration) *MatchResourcesApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithResourceRules")
		}
		b.ResourceRules = append(b.ResourceRules, *values[i])
	}
	return b
}

// WithExcludeResourceRules adds the given value to the ExcludeResourceRules field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ExcludeResourceRules field.
func (b *MatchResourcesApplyConfiguration) WithExcludeResourceRules(values ...*NamedRuleWithOperationsApplyConfiguration) *MatchResourcesApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithExcludeResourceRules")
		}
		b.ExcludeResourceRules = append(b.ExcludeResourceRules, *values[i])
	}
	return b
}

// WithMatchPolicy sets the MatchPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MatchPolicy field is set to the value of the last call.
func (b *MatchResourcesApplyConfiguration) WithMatchPolicy(value admissionregistrationv1alpha1.MatchPolicyType) *MatchResourcesApplyConfiguration {
	b.MatchPolicy = &value
	return b
}
