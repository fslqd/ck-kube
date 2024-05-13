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
	v1alpha1 "github.com/fslqd/ck-kube/kubernetes/api/storage/v1alpha1"
	metav1 "github.com/fslqd/ck-kube/kubernetes/apimachinery/pkg/apis/meta/v1"
	types "github.com/fslqd/ck-kube/kubernetes/apimachinery/pkg/types"
	managedfields "github.com/fslqd/ck-kube/kubernetes/apimachinery/pkg/util/managedfields"
	internal "github.com/fslqd/ck-kube/kubernetes/client-go/applyconfigurations/internal"
	v1 "github.com/fslqd/ck-kube/kubernetes/client-go/applyconfigurations/meta/v1"
)

// VolumeAttributesClassApplyConfiguration represents an declarative configuration of the VolumeAttributesClass type for use
// with apply.
type VolumeAttributesClassApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	DriverName                       *string           `json:"driverName,omitempty"`
	Parameters                       map[string]string `json:"parameters,omitempty"`
}

// VolumeAttributesClass constructs an declarative configuration of the VolumeAttributesClass type for use with
// apply.
func VolumeAttributesClass(name string) *VolumeAttributesClassApplyConfiguration {
	b := &VolumeAttributesClassApplyConfiguration{}
	b.WithName(name)
	b.WithKind("VolumeAttributesClass")
	b.WithAPIVersion("storage.github.com/fslqd/ck-kube/kubernetes/v1alpha1")
	return b
}

// ExtractVolumeAttributesClass extracts the applied configuration owned by fieldManager from
// volumeAttributesClass. If no managedFields are found in volumeAttributesClass for fieldManager, a
// VolumeAttributesClassApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// volumeAttributesClass must be a unmodified VolumeAttributesClass API object that was retrieved from the Kubernetes API.
// ExtractVolumeAttributesClass provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractVolumeAttributesClass(volumeAttributesClass *v1alpha1.VolumeAttributesClass, fieldManager string) (*VolumeAttributesClassApplyConfiguration, error) {
	return extractVolumeAttributesClass(volumeAttributesClass, fieldManager, "")
}

// ExtractVolumeAttributesClassStatus is the same as ExtractVolumeAttributesClass except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractVolumeAttributesClassStatus(volumeAttributesClass *v1alpha1.VolumeAttributesClass, fieldManager string) (*VolumeAttributesClassApplyConfiguration, error) {
	return extractVolumeAttributesClass(volumeAttributesClass, fieldManager, "status")
}

func extractVolumeAttributesClass(volumeAttributesClass *v1alpha1.VolumeAttributesClass, fieldManager string, subresource string) (*VolumeAttributesClassApplyConfiguration, error) {
	b := &VolumeAttributesClassApplyConfiguration{}
	err := managedfields.ExtractInto(volumeAttributesClass, internal.Parser().Type("io.k8s.api.storage.v1alpha1.VolumeAttributesClass"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(volumeAttributesClass.Name)

	b.WithKind("VolumeAttributesClass")
	b.WithAPIVersion("storage.github.com/fslqd/ck-kube/kubernetes/v1alpha1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *VolumeAttributesClassApplyConfiguration) WithKind(value string) *VolumeAttributesClassApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *VolumeAttributesClassApplyConfiguration) WithAPIVersion(value string) *VolumeAttributesClassApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *VolumeAttributesClassApplyConfiguration) WithName(value string) *VolumeAttributesClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *VolumeAttributesClassApplyConfiguration) WithGenerateName(value string) *VolumeAttributesClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *VolumeAttributesClassApplyConfiguration) WithNamespace(value string) *VolumeAttributesClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *VolumeAttributesClassApplyConfiguration) WithUID(value types.UID) *VolumeAttributesClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *VolumeAttributesClassApplyConfiguration) WithResourceVersion(value string) *VolumeAttributesClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *VolumeAttributesClassApplyConfiguration) WithGeneration(value int64) *VolumeAttributesClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *VolumeAttributesClassApplyConfiguration) WithCreationTimestamp(value metav1.Time) *VolumeAttributesClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *VolumeAttributesClassApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *VolumeAttributesClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *VolumeAttributesClassApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *VolumeAttributesClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *VolumeAttributesClassApplyConfiguration) WithLabels(entries map[string]string) *VolumeAttributesClassApplyConfiguration {
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
func (b *VolumeAttributesClassApplyConfiguration) WithAnnotations(entries map[string]string) *VolumeAttributesClassApplyConfiguration {
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
func (b *VolumeAttributesClassApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *VolumeAttributesClassApplyConfiguration {
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
func (b *VolumeAttributesClassApplyConfiguration) WithFinalizers(values ...string) *VolumeAttributesClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *VolumeAttributesClassApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithDriverName sets the DriverName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DriverName field is set to the value of the last call.
func (b *VolumeAttributesClassApplyConfiguration) WithDriverName(value string) *VolumeAttributesClassApplyConfiguration {
	b.DriverName = &value
	return b
}

// WithParameters puts the entries into the Parameters field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Parameters field,
// overwriting an existing map entries in Parameters field with the same key.
func (b *VolumeAttributesClassApplyConfiguration) WithParameters(entries map[string]string) *VolumeAttributesClassApplyConfiguration {
	if b.Parameters == nil && len(entries) > 0 {
		b.Parameters = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Parameters[k] = v
	}
	return b
}
