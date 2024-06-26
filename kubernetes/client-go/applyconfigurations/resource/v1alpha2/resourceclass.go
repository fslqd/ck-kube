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
	corev1 "github.com/fslqd/ck-kube/kubernetes/client-go/applyconfigurations/core/v1"
	internal "github.com/fslqd/ck-kube/kubernetes/client-go/applyconfigurations/internal"
	v1 "github.com/fslqd/ck-kube/kubernetes/client-go/applyconfigurations/meta/v1"
)

// ResourceClassApplyConfiguration represents an declarative configuration of the ResourceClass type for use
// with apply.
type ResourceClassApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	DriverName                       *string                                             `json:"driverName,omitempty"`
	ParametersRef                    *ResourceClassParametersReferenceApplyConfiguration `json:"parametersRef,omitempty"`
	SuitableNodes                    *corev1.NodeSelectorApplyConfiguration              `json:"suitableNodes,omitempty"`
	StructuredParameters             *bool                                               `json:"structuredParameters,omitempty"`
}

// ResourceClass constructs an declarative configuration of the ResourceClass type for use with
// apply.
func ResourceClass(name string) *ResourceClassApplyConfiguration {
	b := &ResourceClassApplyConfiguration{}
	b.WithName(name)
	b.WithKind("ResourceClass")
	b.WithAPIVersion("resource.github.com/fslqd/ck-kube/kubernetes/v1alpha2")
	return b
}

// ExtractResourceClass extracts the applied configuration owned by fieldManager from
// resourceClass. If no managedFields are found in resourceClass for fieldManager, a
// ResourceClassApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// resourceClass must be a unmodified ResourceClass API object that was retrieved from the Kubernetes API.
// ExtractResourceClass provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractResourceClass(resourceClass *resourcev1alpha2.ResourceClass, fieldManager string) (*ResourceClassApplyConfiguration, error) {
	return extractResourceClass(resourceClass, fieldManager, "")
}

// ExtractResourceClassStatus is the same as ExtractResourceClass except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractResourceClassStatus(resourceClass *resourcev1alpha2.ResourceClass, fieldManager string) (*ResourceClassApplyConfiguration, error) {
	return extractResourceClass(resourceClass, fieldManager, "status")
}

func extractResourceClass(resourceClass *resourcev1alpha2.ResourceClass, fieldManager string, subresource string) (*ResourceClassApplyConfiguration, error) {
	b := &ResourceClassApplyConfiguration{}
	err := managedfields.ExtractInto(resourceClass, internal.Parser().Type("io.k8s.api.resource.v1alpha2.ResourceClass"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(resourceClass.Name)

	b.WithKind("ResourceClass")
	b.WithAPIVersion("resource.github.com/fslqd/ck-kube/kubernetes/v1alpha2")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *ResourceClassApplyConfiguration) WithKind(value string) *ResourceClassApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *ResourceClassApplyConfiguration) WithAPIVersion(value string) *ResourceClassApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ResourceClassApplyConfiguration) WithName(value string) *ResourceClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *ResourceClassApplyConfiguration) WithGenerateName(value string) *ResourceClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *ResourceClassApplyConfiguration) WithNamespace(value string) *ResourceClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *ResourceClassApplyConfiguration) WithUID(value types.UID) *ResourceClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *ResourceClassApplyConfiguration) WithResourceVersion(value string) *ResourceClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *ResourceClassApplyConfiguration) WithGeneration(value int64) *ResourceClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *ResourceClassApplyConfiguration) WithCreationTimestamp(value metav1.Time) *ResourceClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *ResourceClassApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *ResourceClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *ResourceClassApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *ResourceClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *ResourceClassApplyConfiguration) WithLabels(entries map[string]string) *ResourceClassApplyConfiguration {
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
func (b *ResourceClassApplyConfiguration) WithAnnotations(entries map[string]string) *ResourceClassApplyConfiguration {
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
func (b *ResourceClassApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *ResourceClassApplyConfiguration {
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
func (b *ResourceClassApplyConfiguration) WithFinalizers(values ...string) *ResourceClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *ResourceClassApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithDriverName sets the DriverName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DriverName field is set to the value of the last call.
func (b *ResourceClassApplyConfiguration) WithDriverName(value string) *ResourceClassApplyConfiguration {
	b.DriverName = &value
	return b
}

// WithParametersRef sets the ParametersRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ParametersRef field is set to the value of the last call.
func (b *ResourceClassApplyConfiguration) WithParametersRef(value *ResourceClassParametersReferenceApplyConfiguration) *ResourceClassApplyConfiguration {
	b.ParametersRef = value
	return b
}

// WithSuitableNodes sets the SuitableNodes field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SuitableNodes field is set to the value of the last call.
func (b *ResourceClassApplyConfiguration) WithSuitableNodes(value *corev1.NodeSelectorApplyConfiguration) *ResourceClassApplyConfiguration {
	b.SuitableNodes = value
	return b
}

// WithStructuredParameters sets the StructuredParameters field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StructuredParameters field is set to the value of the last call.
func (b *ResourceClassApplyConfiguration) WithStructuredParameters(value bool) *ResourceClassApplyConfiguration {
	b.StructuredParameters = &value
	return b
}
