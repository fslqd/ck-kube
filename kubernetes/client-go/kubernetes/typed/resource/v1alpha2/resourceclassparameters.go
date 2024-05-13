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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha2

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1alpha2 "ck-kube/kubernetes/api/resource/v1alpha2"
	v1 "ck-kube/kubernetes/apimachinery/pkg/apis/meta/v1"
	types "ck-kube/kubernetes/apimachinery/pkg/types"
	watch "ck-kube/kubernetes/apimachinery/pkg/watch"
	resourcev1alpha2 "ck-kube/kubernetes/client-go/applyconfigurations/resource/v1alpha2"
	scheme "ck-kube/kubernetes/client-go/kubernetes/scheme"
	rest "ck-kube/kubernetes/client-go/rest"
)

// ResourceClassParametersGetter has a method to return a ResourceClassParametersInterface.
// A group's client should implement this interface.
type ResourceClassParametersGetter interface {
	ResourceClassParameters(namespace string) ResourceClassParametersInterface
}

// ResourceClassParametersInterface has methods to work with ResourceClassParameters resources.
type ResourceClassParametersInterface interface {
	Create(ctx context.Context, resourceClassParameters *v1alpha2.ResourceClassParameters, opts v1.CreateOptions) (*v1alpha2.ResourceClassParameters, error)
	Update(ctx context.Context, resourceClassParameters *v1alpha2.ResourceClassParameters, opts v1.UpdateOptions) (*v1alpha2.ResourceClassParameters, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.ResourceClassParameters, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.ResourceClassParametersList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.ResourceClassParameters, err error)
	Apply(ctx context.Context, resourceClassParameters *resourcev1alpha2.ResourceClassParametersApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha2.ResourceClassParameters, err error)
	ResourceClassParametersExpansion
}

// resourceClassParameters implements ResourceClassParametersInterface
type resourceClassParameters struct {
	client rest.Interface
	ns     string
}

// newResourceClassParameters returns a ResourceClassParameters
func newResourceClassParameters(c *ResourceV1alpha2Client, namespace string) *resourceClassParameters {
	return &resourceClassParameters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the resourceClassParameters, and returns the corresponding resourceClassParameters object, and an error if there is any.
func (c *resourceClassParameters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.ResourceClassParameters, err error) {
	result = &v1alpha2.ResourceClassParameters{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("resourceclassparameters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ResourceClassParameters that match those selectors.
func (c *resourceClassParameters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.ResourceClassParametersList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.ResourceClassParametersList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("resourceclassparameters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested resourceClassParameters.
func (c *resourceClassParameters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("resourceclassparameters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a resourceClassParameters and creates it.  Returns the server's representation of the resourceClassParameters, and an error, if there is any.
func (c *resourceClassParameters) Create(ctx context.Context, resourceClassParameters *v1alpha2.ResourceClassParameters, opts v1.CreateOptions) (result *v1alpha2.ResourceClassParameters, err error) {
	result = &v1alpha2.ResourceClassParameters{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("resourceclassparameters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(resourceClassParameters).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a resourceClassParameters and updates it. Returns the server's representation of the resourceClassParameters, and an error, if there is any.
func (c *resourceClassParameters) Update(ctx context.Context, resourceClassParameters *v1alpha2.ResourceClassParameters, opts v1.UpdateOptions) (result *v1alpha2.ResourceClassParameters, err error) {
	result = &v1alpha2.ResourceClassParameters{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("resourceclassparameters").
		Name(resourceClassParameters.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(resourceClassParameters).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the resourceClassParameters and deletes it. Returns an error if one occurs.
func (c *resourceClassParameters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("resourceclassparameters").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *resourceClassParameters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("resourceclassparameters").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched resourceClassParameters.
func (c *resourceClassParameters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.ResourceClassParameters, err error) {
	result = &v1alpha2.ResourceClassParameters{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("resourceclassparameters").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied resourceClassParameters.
func (c *resourceClassParameters) Apply(ctx context.Context, resourceClassParameters *resourcev1alpha2.ResourceClassParametersApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha2.ResourceClassParameters, err error) {
	if resourceClassParameters == nil {
		return nil, fmt.Errorf("resourceClassParameters provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(resourceClassParameters)
	if err != nil {
		return nil, err
	}
	name := resourceClassParameters.Name
	if name == nil {
		return nil, fmt.Errorf("resourceClassParameters.Name must be provided to Apply")
	}
	result = &v1alpha2.ResourceClassParameters{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("resourceclassparameters").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}