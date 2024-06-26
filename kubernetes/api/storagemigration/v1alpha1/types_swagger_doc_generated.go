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

package v1alpha1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-codegen.sh

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.
var map_GroupVersionResource = map[string]string{
	"":         "The names of the group, the version, and the resource.",
	"group":    "The name of the group.",
	"version":  "The name of the version.",
	"resource": "The name of the resource.",
}

func (GroupVersionResource) SwaggerDoc() map[string]string {
	return map_GroupVersionResource
}

var map_MigrationCondition = map[string]string{
	"":               "Describes the state of a migration at a certain point.",
	"type":           "Type of the condition.",
	"status":         "Status of the condition, one of True, False, Unknown.",
	"lastUpdateTime": "The last time this condition was updated.",
	"reason":         "The reason for the condition's last transition.",
	"message":        "A human readable message indicating details about the transition.",
}

func (MigrationCondition) SwaggerDoc() map[string]string {
	return map_MigrationCondition
}

var map_StorageVersionMigration = map[string]string{
	"":         "StorageVersionMigration represents a migration of stored data to the latest storage version.",
	"metadata": "Standard object metadata. More info: https://git.github.com/fslqd/ck-kube/kubernetes/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "Specification of the migration.",
	"status":   "Status of the migration.",
}

func (StorageVersionMigration) SwaggerDoc() map[string]string {
	return map_StorageVersionMigration
}

var map_StorageVersionMigrationList = map[string]string{
	"":         "StorageVersionMigrationList is a collection of storage version migrations.",
	"metadata": "Standard list metadata More info: https://git.github.com/fslqd/ck-kube/kubernetes/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"items":    "Items is the list of StorageVersionMigration",
}

func (StorageVersionMigrationList) SwaggerDoc() map[string]string {
	return map_StorageVersionMigrationList
}

var map_StorageVersionMigrationSpec = map[string]string{
	"":              "Spec of the storage version migration.",
	"resource":      "The resource that is being migrated. The migrator sends requests to the endpoint serving the resource. Immutable.",
	"continueToken": "The token used in the list options to get the next chunk of objects to migrate. When the .status.conditions indicates the migration is \"Running\", users can use this token to check the progress of the migration.",
}

func (StorageVersionMigrationSpec) SwaggerDoc() map[string]string {
	return map_StorageVersionMigrationSpec
}

var map_StorageVersionMigrationStatus = map[string]string{
	"":                "Status of the storage version migration.",
	"conditions":      "The latest available observations of the migration's current state.",
	"resourceVersion": "ResourceVersion to compare with the GC cache for performing the migration. This is the current resource version of given group, version and resource when kube-controller-manager first observes this StorageVersionMigration resource.",
}

func (StorageVersionMigrationStatus) SwaggerDoc() map[string]string {
	return map_StorageVersionMigrationStatus
}

// AUTO-GENERATED FUNCTIONS END HERE
