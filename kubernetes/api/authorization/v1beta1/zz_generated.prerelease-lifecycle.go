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

// Code generated by prerelease-lifecycle-gen. DO NOT EDIT.

package v1beta1

import (
	schema "ck-kube/kubernetes/apimachinery/pkg/runtime/schema"
)

// APILifecycleIntroduced is an autogenerated function, returning the release in which the API struct was introduced as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:introduced" tags in types.go.
func (in *LocalSubjectAccessReview) APILifecycleIntroduced() (major, minor int) {
	return 1, 2
}

// APILifecycleDeprecated is an autogenerated function, returning the release in which the API struct was or will be deprecated as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:deprecated" tags in types.go or  "k8s:prerelease-lifecycle-gen:introduced" plus three minor.
func (in *LocalSubjectAccessReview) APILifecycleDeprecated() (major, minor int) {
	return 1, 19
}

// APILifecycleReplacement is an autogenerated function, returning the group, version, and kind that should be used instead of this deprecated type.
// It is controlled by "k8s:prerelease-lifecycle-gen:replacement=<group>,<version>,<kind>" tags in types.go.
func (in *LocalSubjectAccessReview) APILifecycleReplacement() schema.GroupVersionKind {
	return schema.GroupVersionKind{Group: "authorization.ck-kube/kubernetes", Version: "v1", Kind: "LocalSubjectAccessReview"}
}

// APILifecycleRemoved is an autogenerated function, returning the release in which the API is no longer served as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:removed" tags in types.go or  "k8s:prerelease-lifecycle-gen:deprecated" plus three minor.
func (in *LocalSubjectAccessReview) APILifecycleRemoved() (major, minor int) {
	return 1, 22
}

// APILifecycleIntroduced is an autogenerated function, returning the release in which the API struct was introduced as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:introduced" tags in types.go.
func (in *SelfSubjectAccessReview) APILifecycleIntroduced() (major, minor int) {
	return 1, 2
}

// APILifecycleDeprecated is an autogenerated function, returning the release in which the API struct was or will be deprecated as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:deprecated" tags in types.go or  "k8s:prerelease-lifecycle-gen:introduced" plus three minor.
func (in *SelfSubjectAccessReview) APILifecycleDeprecated() (major, minor int) {
	return 1, 19
}

// APILifecycleReplacement is an autogenerated function, returning the group, version, and kind that should be used instead of this deprecated type.
// It is controlled by "k8s:prerelease-lifecycle-gen:replacement=<group>,<version>,<kind>" tags in types.go.
func (in *SelfSubjectAccessReview) APILifecycleReplacement() schema.GroupVersionKind {
	return schema.GroupVersionKind{Group: "authorization.ck-kube/kubernetes", Version: "v1", Kind: "SelfSubjectAccessReview"}
}

// APILifecycleRemoved is an autogenerated function, returning the release in which the API is no longer served as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:removed" tags in types.go or  "k8s:prerelease-lifecycle-gen:deprecated" plus three minor.
func (in *SelfSubjectAccessReview) APILifecycleRemoved() (major, minor int) {
	return 1, 22
}

// APILifecycleIntroduced is an autogenerated function, returning the release in which the API struct was introduced as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:introduced" tags in types.go.
func (in *SelfSubjectRulesReview) APILifecycleIntroduced() (major, minor int) {
	return 1, 8
}

// APILifecycleDeprecated is an autogenerated function, returning the release in which the API struct was or will be deprecated as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:deprecated" tags in types.go or  "k8s:prerelease-lifecycle-gen:introduced" plus three minor.
func (in *SelfSubjectRulesReview) APILifecycleDeprecated() (major, minor int) {
	return 1, 19
}

// APILifecycleReplacement is an autogenerated function, returning the group, version, and kind that should be used instead of this deprecated type.
// It is controlled by "k8s:prerelease-lifecycle-gen:replacement=<group>,<version>,<kind>" tags in types.go.
func (in *SelfSubjectRulesReview) APILifecycleReplacement() schema.GroupVersionKind {
	return schema.GroupVersionKind{Group: "authorization.ck-kube/kubernetes", Version: "v1", Kind: "SelfSubjectRulesReview"}
}

// APILifecycleRemoved is an autogenerated function, returning the release in which the API is no longer served as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:removed" tags in types.go or  "k8s:prerelease-lifecycle-gen:deprecated" plus three minor.
func (in *SelfSubjectRulesReview) APILifecycleRemoved() (major, minor int) {
	return 1, 22
}

// APILifecycleIntroduced is an autogenerated function, returning the release in which the API struct was introduced as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:introduced" tags in types.go.
func (in *SubjectAccessReview) APILifecycleIntroduced() (major, minor int) {
	return 1, 2
}

// APILifecycleDeprecated is an autogenerated function, returning the release in which the API struct was or will be deprecated as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:deprecated" tags in types.go or  "k8s:prerelease-lifecycle-gen:introduced" plus three minor.
func (in *SubjectAccessReview) APILifecycleDeprecated() (major, minor int) {
	return 1, 19
}

// APILifecycleReplacement is an autogenerated function, returning the group, version, and kind that should be used instead of this deprecated type.
// It is controlled by "k8s:prerelease-lifecycle-gen:replacement=<group>,<version>,<kind>" tags in types.go.
func (in *SubjectAccessReview) APILifecycleReplacement() schema.GroupVersionKind {
	return schema.GroupVersionKind{Group: "authorization.ck-kube/kubernetes", Version: "v1", Kind: "SubjectAccessReview"}
}

// APILifecycleRemoved is an autogenerated function, returning the release in which the API is no longer served as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:removed" tags in types.go or  "k8s:prerelease-lifecycle-gen:deprecated" plus three minor.
func (in *SubjectAccessReview) APILifecycleRemoved() (major, minor int) {
	return 1, 22
}