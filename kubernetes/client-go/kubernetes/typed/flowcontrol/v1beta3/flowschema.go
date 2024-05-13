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

package v1beta3

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1beta3 "github.com/fslqd/ck-kube/kubernetes/api/flowcontrol/v1beta3"
	v1 "github.com/fslqd/ck-kube/kubernetes/apimachinery/pkg/apis/meta/v1"
	types "github.com/fslqd/ck-kube/kubernetes/apimachinery/pkg/types"
	watch "github.com/fslqd/ck-kube/kubernetes/apimachinery/pkg/watch"
	flowcontrolv1beta3 "github.com/fslqd/ck-kube/kubernetes/client-go/applyconfigurations/flowcontrol/v1beta3"
	scheme "github.com/fslqd/ck-kube/kubernetes/client-go/kubernetes/scheme"
	rest "github.com/fslqd/ck-kube/kubernetes/client-go/rest"
)

// FlowSchemasGetter has a method to return a FlowSchemaInterface.
// A group's client should implement this interface.
type FlowSchemasGetter interface {
	FlowSchemas() FlowSchemaInterface
}

// FlowSchemaInterface has methods to work with FlowSchema resources.
type FlowSchemaInterface interface {
	Create(ctx context.Context, flowSchema *v1beta3.FlowSchema, opts v1.CreateOptions) (*v1beta3.FlowSchema, error)
	Update(ctx context.Context, flowSchema *v1beta3.FlowSchema, opts v1.UpdateOptions) (*v1beta3.FlowSchema, error)
	UpdateStatus(ctx context.Context, flowSchema *v1beta3.FlowSchema, opts v1.UpdateOptions) (*v1beta3.FlowSchema, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta3.FlowSchema, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta3.FlowSchemaList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta3.FlowSchema, err error)
	Apply(ctx context.Context, flowSchema *flowcontrolv1beta3.FlowSchemaApplyConfiguration, opts v1.ApplyOptions) (result *v1beta3.FlowSchema, err error)
	ApplyStatus(ctx context.Context, flowSchema *flowcontrolv1beta3.FlowSchemaApplyConfiguration, opts v1.ApplyOptions) (result *v1beta3.FlowSchema, err error)
	FlowSchemaExpansion
}

// flowSchemas implements FlowSchemaInterface
type flowSchemas struct {
	client rest.Interface
}

// newFlowSchemas returns a FlowSchemas
func newFlowSchemas(c *FlowcontrolV1beta3Client) *flowSchemas {
	return &flowSchemas{
		client: c.RESTClient(),
	}
}

// Get takes name of the flowSchema, and returns the corresponding flowSchema object, and an error if there is any.
func (c *flowSchemas) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta3.FlowSchema, err error) {
	result = &v1beta3.FlowSchema{}
	err = c.client.Get().
		Resource("flowschemas").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FlowSchemas that match those selectors.
func (c *flowSchemas) List(ctx context.Context, opts v1.ListOptions) (result *v1beta3.FlowSchemaList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta3.FlowSchemaList{}
	err = c.client.Get().
		Resource("flowschemas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested flowSchemas.
func (c *flowSchemas) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("flowschemas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a flowSchema and creates it.  Returns the server's representation of the flowSchema, and an error, if there is any.
func (c *flowSchemas) Create(ctx context.Context, flowSchema *v1beta3.FlowSchema, opts v1.CreateOptions) (result *v1beta3.FlowSchema, err error) {
	result = &v1beta3.FlowSchema{}
	err = c.client.Post().
		Resource("flowschemas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(flowSchema).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a flowSchema and updates it. Returns the server's representation of the flowSchema, and an error, if there is any.
func (c *flowSchemas) Update(ctx context.Context, flowSchema *v1beta3.FlowSchema, opts v1.UpdateOptions) (result *v1beta3.FlowSchema, err error) {
	result = &v1beta3.FlowSchema{}
	err = c.client.Put().
		Resource("flowschemas").
		Name(flowSchema.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(flowSchema).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *flowSchemas) UpdateStatus(ctx context.Context, flowSchema *v1beta3.FlowSchema, opts v1.UpdateOptions) (result *v1beta3.FlowSchema, err error) {
	result = &v1beta3.FlowSchema{}
	err = c.client.Put().
		Resource("flowschemas").
		Name(flowSchema.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(flowSchema).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the flowSchema and deletes it. Returns an error if one occurs.
func (c *flowSchemas) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("flowschemas").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *flowSchemas) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("flowschemas").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched flowSchema.
func (c *flowSchemas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta3.FlowSchema, err error) {
	result = &v1beta3.FlowSchema{}
	err = c.client.Patch(pt).
		Resource("flowschemas").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied flowSchema.
func (c *flowSchemas) Apply(ctx context.Context, flowSchema *flowcontrolv1beta3.FlowSchemaApplyConfiguration, opts v1.ApplyOptions) (result *v1beta3.FlowSchema, err error) {
	if flowSchema == nil {
		return nil, fmt.Errorf("flowSchema provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(flowSchema)
	if err != nil {
		return nil, err
	}
	name := flowSchema.Name
	if name == nil {
		return nil, fmt.Errorf("flowSchema.Name must be provided to Apply")
	}
	result = &v1beta3.FlowSchema{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("flowschemas").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *flowSchemas) ApplyStatus(ctx context.Context, flowSchema *flowcontrolv1beta3.FlowSchemaApplyConfiguration, opts v1.ApplyOptions) (result *v1beta3.FlowSchema, err error) {
	if flowSchema == nil {
		return nil, fmt.Errorf("flowSchema provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(flowSchema)
	if err != nil {
		return nil, err
	}

	name := flowSchema.Name
	if name == nil {
		return nil, fmt.Errorf("flowSchema.Name must be provided to Apply")
	}

	result = &v1beta3.FlowSchema{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("flowschemas").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
