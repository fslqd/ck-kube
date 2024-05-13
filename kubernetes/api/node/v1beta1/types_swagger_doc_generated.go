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

package v1beta1

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
var map_Overhead = map[string]string{
	"":         "Overhead structure represents the resource overhead associated with running a pod.",
	"podFixed": "podFixed represents the fixed resource overhead associated with running a pod.",
}

func (Overhead) SwaggerDoc() map[string]string {
	return map_Overhead
}

var map_RuntimeClass = map[string]string{
	"":           "RuntimeClass defines a class of container runtime supported in the cluster. The RuntimeClass is used to determine which container runtime is used to run all containers in a pod. RuntimeClasses are (currently) manually defined by a user or cluster provisioner, and referenced in the PodSpec. The Kubelet is responsible for resolving the RuntimeClassName reference before running the pod.  For more details, see https://git.ck-kube/kubernetes/enhancements/keps/sig-node/585-runtime-class",
	"metadata":   "More info: https://git.ck-kube/kubernetes/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"handler":    "handler specifies the underlying runtime and configuration that the CRI implementation will use to handle pods of this class. The possible values are specific to the node & CRI configuration.  It is assumed that all handlers are available on every node, and handlers of the same name are equivalent on every node. For example, a handler called \"runc\" might specify that the runc OCI runtime (using native Linux containers) will be used to run the containers in a pod. The handler must be lowercase, conform to the DNS Label (RFC 1123) requirements, and is immutable.",
	"overhead":   "overhead represents the resource overhead associated with running a pod for a given RuntimeClass. For more details, see https://git.ck-kube/kubernetes/enhancements/keps/sig-node/688-pod-overhead/README.md",
	"scheduling": "scheduling holds the scheduling constraints to ensure that pods running with this RuntimeClass are scheduled to nodes that support it. If scheduling is nil, this RuntimeClass is assumed to be supported by all nodes.",
}

func (RuntimeClass) SwaggerDoc() map[string]string {
	return map_RuntimeClass
}

var map_RuntimeClassList = map[string]string{
	"":         "RuntimeClassList is a list of RuntimeClass objects.",
	"metadata": "Standard list metadata. More info: https://git.ck-kube/kubernetes/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"items":    "items is a list of schema objects.",
}

func (RuntimeClassList) SwaggerDoc() map[string]string {
	return map_RuntimeClassList
}

var map_Scheduling = map[string]string{
	"":             "Scheduling specifies the scheduling constraints for nodes supporting a RuntimeClass.",
	"nodeSelector": "nodeSelector lists labels that must be present on nodes that support this RuntimeClass. Pods using this RuntimeClass can only be scheduled to a node matched by this selector. The RuntimeClass nodeSelector is merged with a pod's existing nodeSelector. Any conflicts will cause the pod to be rejected in admission.",
	"tolerations":  "tolerations are appended (excluding duplicates) to pods running with this RuntimeClass during admission, effectively unioning the set of nodes tolerated by the pod and the RuntimeClass.",
}

func (Scheduling) SwaggerDoc() map[string]string {
	return map_Scheduling
}

// AUTO-GENERATED FUNCTIONS END HERE
