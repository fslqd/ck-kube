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
	apicorev1 "github.com/fslqd/ck-kube/kubernetes/api/core/v1"
	metav1 "github.com/fslqd/ck-kube/kubernetes/apimachinery/pkg/apis/meta/v1"
	types "github.com/fslqd/ck-kube/kubernetes/apimachinery/pkg/types"
	managedfields "github.com/fslqd/ck-kube/kubernetes/apimachinery/pkg/util/managedfields"
	internal "github.com/fslqd/ck-kube/kubernetes/client-go/applyconfigurations/internal"
	v1 "github.com/fslqd/ck-kube/kubernetes/client-go/applyconfigurations/meta/v1"
)

// NamespaceApplyConfiguration represents an declarative configuration of the Namespace type for use
// with apply.
type NamespaceApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                             *NamespaceSpecApplyConfiguration   `json:"spec,omitempty"`
	Status                           *NamespaceStatusApplyConfiguration `json:"status,omitempty"`
}

// Namespace constructs an declarative configuration of the Namespace type for use with
// apply.
func Namespace(name string) *NamespaceApplyConfiguration {
	b := &NamespaceApplyConfiguration{}
	b.WithName(name)
	b.WithKind("Namespace")
	b.WithAPIVersion("v1")
	return b
}

// ExtractNamespace extracts the applied configuration owned by fieldManager from
// namespace. If no managedFields are found in namespace for fieldManager, a
// NamespaceApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// namespace must be a unmodified Namespace API object that was retrieved from the Kubernetes API.
// ExtractNamespace provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractNamespace(namespace *apicorev1.Namespace, fieldManager string) (*NamespaceApplyConfiguration, error) {
	return extractNamespace(namespace, fieldManager, "")
}

// ExtractNamespaceStatus is the same as ExtractNamespace except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractNamespaceStatus(namespace *apicorev1.Namespace, fieldManager string) (*NamespaceApplyConfiguration, error) {
	return extractNamespace(namespace, fieldManager, "status")
}

func extractNamespace(namespace *apicorev1.Namespace, fieldManager string, subresource string) (*NamespaceApplyConfiguration, error) {
	b := &NamespaceApplyConfiguration{}
	err := managedfields.ExtractInto(namespace, internal.Parser().Type("io.k8s.api.core.v1.Namespace"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(namespace.Name)

	b.WithKind("Namespace")
	b.WithAPIVersion("v1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *NamespaceApplyConfiguration) WithKind(value string) *NamespaceApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *NamespaceApplyConfiguration) WithAPIVersion(value string) *NamespaceApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *NamespaceApplyConfiguration) WithName(value string) *NamespaceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *NamespaceApplyConfiguration) WithGenerateName(value string) *NamespaceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *NamespaceApplyConfiguration) WithNamespace(value string) *NamespaceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *NamespaceApplyConfiguration) WithUID(value types.UID) *NamespaceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *NamespaceApplyConfiguration) WithResourceVersion(value string) *NamespaceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *NamespaceApplyConfiguration) WithGeneration(value int64) *NamespaceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *NamespaceApplyConfiguration) WithCreationTimestamp(value metav1.Time) *NamespaceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *NamespaceApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *NamespaceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *NamespaceApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *NamespaceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *NamespaceApplyConfiguration) WithLabels(entries map[string]string) *NamespaceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *NamespaceApplyConfiguration) WithAnnotations(entries map[string]string) *NamespaceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *NamespaceApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *NamespaceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *NamespaceApplyConfiguration) WithFinalizers(values ...string) *NamespaceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *NamespaceApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *NamespaceApplyConfiguration) WithSpec(value *NamespaceSpecApplyConfiguration) *NamespaceApplyConfiguration {
	b.Spec = value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *NamespaceApplyConfiguration) WithStatus(value *NamespaceStatusApplyConfiguration) *NamespaceApplyConfiguration {
	b.Status = value
	return b
}
