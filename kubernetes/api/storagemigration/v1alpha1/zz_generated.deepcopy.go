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

package v1alpha1

import (
	runtime "ck-kube/kubernetes/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupVersionResource) DeepCopyInto(out *GroupVersionResource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupVersionResource.
func (in *GroupVersionResource) DeepCopy() *GroupVersionResource {
	if in == nil {
		return nil
	}
	out := new(GroupVersionResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigrationCondition) DeepCopyInto(out *MigrationCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigrationCondition.
func (in *MigrationCondition) DeepCopy() *MigrationCondition {
	if in == nil {
		return nil
	}
	out := new(MigrationCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageVersionMigration) DeepCopyInto(out *StorageVersionMigration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageVersionMigration.
func (in *StorageVersionMigration) DeepCopy() *StorageVersionMigration {
	if in == nil {
		return nil
	}
	out := new(StorageVersionMigration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageVersionMigration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageVersionMigrationList) DeepCopyInto(out *StorageVersionMigrationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StorageVersionMigration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageVersionMigrationList.
func (in *StorageVersionMigrationList) DeepCopy() *StorageVersionMigrationList {
	if in == nil {
		return nil
	}
	out := new(StorageVersionMigrationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageVersionMigrationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageVersionMigrationSpec) DeepCopyInto(out *StorageVersionMigrationSpec) {
	*out = *in
	out.Resource = in.Resource
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageVersionMigrationSpec.
func (in *StorageVersionMigrationSpec) DeepCopy() *StorageVersionMigrationSpec {
	if in == nil {
		return nil
	}
	out := new(StorageVersionMigrationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageVersionMigrationStatus) DeepCopyInto(out *StorageVersionMigrationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]MigrationCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageVersionMigrationStatus.
func (in *StorageVersionMigrationStatus) DeepCopy() *StorageVersionMigrationStatus {
	if in == nil {
		return nil
	}
	out := new(StorageVersionMigrationStatus)
	in.DeepCopyInto(out)
	return out
}
