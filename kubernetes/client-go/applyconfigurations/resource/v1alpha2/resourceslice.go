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
	resourcev1alpha2 "github.com/fslqd/ck-kube/kubernetes/api/resource/v1alpha2"
	metav1 "github.com/fslqd/ck-kube/kubernetes/apimachinery/pkg/apis/meta/v1"
	types "github.com/fslqd/ck-kube/kubernetes/apimachinery/pkg/types"
	managedfields "github.com/fslqd/ck-kube/kubernetes/apimachinery/pkg/util/managedfields"
	internal "github.com/fslqd/ck-kube/kubernetes/client-go/applyconfigurations/internal"
	v1 "github.com/fslqd/ck-kube/kubernetes/client-go/applyconfigurations/meta/v1"
)

// ResourceSliceApplyConfiguration represents an declarative configuration of the ResourceSlice type for use
// with apply.
type ResourceSliceApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	NodeName                         *string `json:"nodeName,omitempty"`
	DriverName                       *string `json:"driverName,omitempty"`
	ResourceModelApplyConfiguration  `json:",inline"`
}

// ResourceSlice constructs an declarative configuration of the ResourceSlice type for use with
// apply.
func ResourceSlice(name string) *ResourceSliceApplyConfiguration {
	b := &ResourceSliceApplyConfiguration{}
	b.WithName(name)
	b.WithKind("ResourceSlice")
	b.WithAPIVersion("resource.github.com/fslqd/ck-kube/kubernetes/v1alpha2")
	return b
}

// ExtractResourceSlice extracts the applied configuration owned by fieldManager from
// resourceSlice. If no managedFields are found in resourceSlice for fieldManager, a
// ResourceSliceApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// resourceSlice must be a unmodified ResourceSlice API object that was retrieved from the Kubernetes API.
// ExtractResourceSlice provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractResourceSlice(resourceSlice *resourcev1alpha2.ResourceSlice, fieldManager string) (*ResourceSliceApplyConfiguration, error) {
	return extractResourceSlice(resourceSlice, fieldManager, "")
}

// ExtractResourceSliceStatus is the same as ExtractResourceSlice except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractResourceSliceStatus(resourceSlice *resourcev1alpha2.ResourceSlice, fieldManager string) (*ResourceSliceApplyConfiguration, error) {
	return extractResourceSlice(resourceSlice, fieldManager, "status")
}

func extractResourceSlice(resourceSlice *resourcev1alpha2.ResourceSlice, fieldManager string, subresource string) (*ResourceSliceApplyConfiguration, error) {
	b := &ResourceSliceApplyConfiguration{}
	err := managedfields.ExtractInto(resourceSlice, internal.Parser().Type("io.k8s.api.resource.v1alpha2.ResourceSlice"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(resourceSlice.Name)

	b.WithKind("ResourceSlice")
	b.WithAPIVersion("resource.github.com/fslqd/ck-kube/kubernetes/v1alpha2")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *ResourceSliceApplyConfiguration) WithKind(value string) *ResourceSliceApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *ResourceSliceApplyConfiguration) WithAPIVersion(value string) *ResourceSliceApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ResourceSliceApplyConfiguration) WithName(value string) *ResourceSliceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *ResourceSliceApplyConfiguration) WithGenerateName(value string) *ResourceSliceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *ResourceSliceApplyConfiguration) WithNamespace(value string) *ResourceSliceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *ResourceSliceApplyConfiguration) WithUID(value types.UID) *ResourceSliceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *ResourceSliceApplyConfiguration) WithResourceVersion(value string) *ResourceSliceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *ResourceSliceApplyConfiguration) WithGeneration(value int64) *ResourceSliceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *ResourceSliceApplyConfiguration) WithCreationTimestamp(value metav1.Time) *ResourceSliceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *ResourceSliceApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *ResourceSliceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *ResourceSliceApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *ResourceSliceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *ResourceSliceApplyConfiguration) WithLabels(entries map[string]string) *ResourceSliceApplyConfiguration {
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
func (b *ResourceSliceApplyConfiguration) WithAnnotations(entries map[string]string) *ResourceSliceApplyConfiguration {
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
func (b *ResourceSliceApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *ResourceSliceApplyConfiguration {
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
func (b *ResourceSliceApplyConfiguration) WithFinalizers(values ...string) *ResourceSliceApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *ResourceSliceApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithNodeName sets the NodeName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeName field is set to the value of the last call.
func (b *ResourceSliceApplyConfiguration) WithNodeName(value string) *ResourceSliceApplyConfiguration {
	b.NodeName = &value
	return b
}

// WithDriverName sets the DriverName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DriverName field is set to the value of the last call.
func (b *ResourceSliceApplyConfiguration) WithDriverName(value string) *ResourceSliceApplyConfiguration {
	b.DriverName = &value
	return b
}

// WithNamedResources sets the NamedResources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NamedResources field is set to the value of the last call.
func (b *ResourceSliceApplyConfiguration) WithNamedResources(value *NamedResourcesResourcesApplyConfiguration) *ResourceSliceApplyConfiguration {
	b.NamedResources = value
	return b
}
