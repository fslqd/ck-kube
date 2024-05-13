//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v2

import (
	v1 "github.com/fslqd/ck-kube/kubernetes/apimachinery/pkg/apis/meta/v1"
	runtime "github.com/fslqd/ck-kube/kubernetes/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIGroupDiscovery) DeepCopyInto(out *APIGroupDiscovery) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]APIVersionDiscovery, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIGroupDiscovery.
func (in *APIGroupDiscovery) DeepCopy() *APIGroupDiscovery {
	if in == nil {
		return nil
	}
	out := new(APIGroupDiscovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIGroupDiscovery) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIGroupDiscoveryList) DeepCopyInto(out *APIGroupDiscoveryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]APIGroupDiscovery, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIGroupDiscoveryList.
func (in *APIGroupDiscoveryList) DeepCopy() *APIGroupDiscoveryList {
	if in == nil {
		return nil
	}
	out := new(APIGroupDiscoveryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIGroupDiscoveryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIResourceDiscovery) DeepCopyInto(out *APIResourceDiscovery) {
	*out = *in
	if in.ResponseKind != nil {
		in, out := &in.ResponseKind, &out.ResponseKind
		*out = new(v1.GroupVersionKind)
		**out = **in
	}
	if in.Verbs != nil {
		in, out := &in.Verbs, &out.Verbs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ShortNames != nil {
		in, out := &in.ShortNames, &out.ShortNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Categories != nil {
		in, out := &in.Categories, &out.Categories
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Subresources != nil {
		in, out := &in.Subresources, &out.Subresources
		*out = make([]APISubresourceDiscovery, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIResourceDiscovery.
func (in *APIResourceDiscovery) DeepCopy() *APIResourceDiscovery {
	if in == nil {
		return nil
	}
	out := new(APIResourceDiscovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APISubresourceDiscovery) DeepCopyInto(out *APISubresourceDiscovery) {
	*out = *in
	if in.ResponseKind != nil {
		in, out := &in.ResponseKind, &out.ResponseKind
		*out = new(v1.GroupVersionKind)
		**out = **in
	}
	if in.AcceptedTypes != nil {
		in, out := &in.AcceptedTypes, &out.AcceptedTypes
		*out = make([]v1.GroupVersionKind, len(*in))
		copy(*out, *in)
	}
	if in.Verbs != nil {
		in, out := &in.Verbs, &out.Verbs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APISubresourceDiscovery.
func (in *APISubresourceDiscovery) DeepCopy() *APISubresourceDiscovery {
	if in == nil {
		return nil
	}
	out := new(APISubresourceDiscovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIVersionDiscovery) DeepCopyInto(out *APIVersionDiscovery) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]APIResourceDiscovery, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIVersionDiscovery.
func (in *APIVersionDiscovery) DeepCopy() *APIVersionDiscovery {
	if in == nil {
		return nil
	}
	out := new(APIVersionDiscovery)
	in.DeepCopyInto(out)
	return out
}
