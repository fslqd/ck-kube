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

package v1alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1alpha1 "ck-kube/kubernetes/api/rbac/v1alpha1"
	v1 "ck-kube/kubernetes/apimachinery/pkg/apis/meta/v1"
	types "ck-kube/kubernetes/apimachinery/pkg/types"
	watch "ck-kube/kubernetes/apimachinery/pkg/watch"
	rbacv1alpha1 "ck-kube/kubernetes/client-go/applyconfigurations/rbac/v1alpha1"
	scheme "ck-kube/kubernetes/client-go/kubernetes/scheme"
	rest "ck-kube/kubernetes/client-go/rest"
)

// ClusterRolesGetter has a method to return a ClusterRoleInterface.
// A group's client should implement this interface.
type ClusterRolesGetter interface {
	ClusterRoles() ClusterRoleInterface
}

// ClusterRoleInterface has methods to work with ClusterRole resources.
type ClusterRoleInterface interface {
	Create(ctx context.Context, clusterRole *v1alpha1.ClusterRole, opts v1.CreateOptions) (*v1alpha1.ClusterRole, error)
	Update(ctx context.Context, clusterRole *v1alpha1.ClusterRole, opts v1.UpdateOptions) (*v1alpha1.ClusterRole, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ClusterRole, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ClusterRoleList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterRole, err error)
	Apply(ctx context.Context, clusterRole *rbacv1alpha1.ClusterRoleApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ClusterRole, err error)
	ClusterRoleExpansion
}

// clusterRoles implements ClusterRoleInterface
type clusterRoles struct {
	client rest.Interface
}

// newClusterRoles returns a ClusterRoles
func newClusterRoles(c *RbacV1alpha1Client) *clusterRoles {
	return &clusterRoles{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterRole, and returns the corresponding clusterRole object, and an error if there is any.
func (c *clusterRoles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterRole, err error) {
	result = &v1alpha1.ClusterRole{}
	err = c.client.Get().
		Resource("clusterroles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterRoles that match those selectors.
func (c *clusterRoles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterRoleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClusterRoleList{}
	err = c.client.Get().
		Resource("clusterroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterRoles.
func (c *clusterRoles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clusterroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterRole and creates it.  Returns the server's representation of the clusterRole, and an error, if there is any.
func (c *clusterRoles) Create(ctx context.Context, clusterRole *v1alpha1.ClusterRole, opts v1.CreateOptions) (result *v1alpha1.ClusterRole, err error) {
	result = &v1alpha1.ClusterRole{}
	err = c.client.Post().
		Resource("clusterroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterRole).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterRole and updates it. Returns the server's representation of the clusterRole, and an error, if there is any.
func (c *clusterRoles) Update(ctx context.Context, clusterRole *v1alpha1.ClusterRole, opts v1.UpdateOptions) (result *v1alpha1.ClusterRole, err error) {
	result = &v1alpha1.ClusterRole{}
	err = c.client.Put().
		Resource("clusterroles").
		Name(clusterRole.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterRole).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterRole and deletes it. Returns an error if one occurs.
func (c *clusterRoles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusterroles").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterRoles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clusterroles").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterRole.
func (c *clusterRoles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterRole, err error) {
	result = &v1alpha1.ClusterRole{}
	err = c.client.Patch(pt).
		Resource("clusterroles").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied clusterRole.
func (c *clusterRoles) Apply(ctx context.Context, clusterRole *rbacv1alpha1.ClusterRoleApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ClusterRole, err error) {
	if clusterRole == nil {
		return nil, fmt.Errorf("clusterRole provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(clusterRole)
	if err != nil {
		return nil, err
	}
	name := clusterRole.Name
	if name == nil {
		return nil, fmt.Errorf("clusterRole.Name must be provided to Apply")
	}
	result = &v1alpha1.ClusterRole{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("clusterroles").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
