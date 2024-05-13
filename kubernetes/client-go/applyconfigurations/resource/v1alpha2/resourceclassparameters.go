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
	resourcev1alpha2 "ck-kube/kubernetes/api/resource/v1alpha2"
	metav1 "ck-kube/kubernetes/apimachinery/pkg/apis/meta/v1"
	types "ck-kube/kubernetes/apimachinery/pkg/types"
	managedfields "ck-kube/kubernetes/apimachinery/pkg/util/managedfields"
	internal "ck-kube/kubernetes/client-go/applyconfigurations/internal"
	v1 "ck-kube/kubernetes/client-go/applyconfigurations/meta/v1"
)

// ResourceClassParametersApplyConfiguration represents an declarative configuration of the ResourceClassParameters type for use
// with apply.
type ResourceClassParametersApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	GeneratedFrom                    *ResourceClassParametersReferenceApplyConfiguration `json:"generatedFrom,omitempty"`
	VendorParameters                 []VendorParametersApplyConfiguration                `json:"vendorParameters,omitempty"`
	Filters                          []ResourceFilterApplyConfiguration                  `json:"filters,omitempty"`
}

// ResourceClassParameters constructs an declarative configuration of the ResourceClassParameters type for use with
// apply.
func ResourceClassParameters(name, namespace string) *ResourceClassParametersApplyConfiguration {
	b := &ResourceClassParametersApplyConfiguration{}
	b.WithName(name)
	b.WithNamespace(namespace)
	b.WithKind("ResourceClassParameters")
	b.WithAPIVersion("resource.ck-kube/kubernetes/v1alpha2")
	return b
}

// ExtractResourceClassParameters extracts the applied configuration owned by fieldManager from
// resourceClassParameters. If no managedFields are found in resourceClassParameters for fieldManager, a
// ResourceClassParametersApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// resourceClassParameters must be a unmodified ResourceClassParameters API object that was retrieved from the Kubernetes API.
// ExtractResourceClassParameters provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractResourceClassParameters(resourceClassParameters *resourcev1alpha2.ResourceClassParameters, fieldManager string) (*ResourceClassParametersApplyConfiguration, error) {
	return extractResourceClassParameters(resourceClassParameters, fieldManager, "")
}

// ExtractResourceClassParametersStatus is the same as ExtractResourceClassParameters except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractResourceClassParametersStatus(resourceClassParameters *resourcev1alpha2.ResourceClassParameters, fieldManager string) (*ResourceClassParametersApplyConfiguration, error) {
	return extractResourceClassParameters(resourceClassParameters, fieldManager, "status")
}

func extractResourceClassParameters(resourceClassParameters *resourcev1alpha2.ResourceClassParameters, fieldManager string, subresource string) (*ResourceClassParametersApplyConfiguration, error) {
	b := &ResourceClassParametersApplyConfiguration{}
	err := managedfields.ExtractInto(resourceClassParameters, internal.Parser().Type("io.k8s.api.resource.v1alpha2.ResourceClassParameters"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(resourceClassParameters.Name)
	b.WithNamespace(resourceClassParameters.Namespace)

	b.WithKind("ResourceClassParameters")
	b.WithAPIVersion("resource.ck-kube/kubernetes/v1alpha2")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *ResourceClassParametersApplyConfiguration) WithKind(value string) *ResourceClassParametersApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *ResourceClassParametersApplyConfiguration) WithAPIVersion(value string) *ResourceClassParametersApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ResourceClassParametersApplyConfiguration) WithName(value string) *ResourceClassParametersApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *ResourceClassParametersApplyConfiguration) WithGenerateName(value string) *ResourceClassParametersApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *ResourceClassParametersApplyConfiguration) WithNamespace(value string) *ResourceClassParametersApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *ResourceClassParametersApplyConfiguration) WithUID(value types.UID) *ResourceClassParametersApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *ResourceClassParametersApplyConfiguration) WithResourceVersion(value string) *ResourceClassParametersApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *ResourceClassParametersApplyConfiguration) WithGeneration(value int64) *ResourceClassParametersApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *ResourceClassParametersApplyConfiguration) WithCreationTimestamp(value metav1.Time) *ResourceClassParametersApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *ResourceClassParametersApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *ResourceClassParametersApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *ResourceClassParametersApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *ResourceClassParametersApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *ResourceClassParametersApplyConfiguration) WithLabels(entries map[string]string) *ResourceClassParametersApplyConfiguration {
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
func (b *ResourceClassParametersApplyConfiguration) WithAnnotations(entries map[string]string) *ResourceClassParametersApplyConfiguration {
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
func (b *ResourceClassParametersApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *ResourceClassParametersApplyConfiguration {
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
func (b *ResourceClassParametersApplyConfiguration) WithFinalizers(values ...string) *ResourceClassParametersApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *ResourceClassParametersApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithGeneratedFrom sets the GeneratedFrom field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GeneratedFrom field is set to the value of the last call.
func (b *ResourceClassParametersApplyConfiguration) WithGeneratedFrom(value *ResourceClassParametersReferenceApplyConfiguration) *ResourceClassParametersApplyConfiguration {
	b.GeneratedFrom = value
	return b
}

// WithVendorParameters adds the given value to the VendorParameters field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the VendorParameters field.
func (b *ResourceClassParametersApplyConfiguration) WithVendorParameters(values ...*VendorParametersApplyConfiguration) *ResourceClassParametersApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithVendorParameters")
		}
		b.VendorParameters = append(b.VendorParameters, *values[i])
	}
	return b
}

// WithFilters adds the given value to the Filters field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Filters field.
func (b *ResourceClassParametersApplyConfiguration) WithFilters(values ...*ResourceFilterApplyConfiguration) *ResourceClassParametersApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithFilters")
		}
		b.Filters = append(b.Filters, *values[i])
	}
	return b
}
