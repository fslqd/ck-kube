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
var map_HTTPIngressPath = map[string]string{
	"":         "HTTPIngressPath associates a path with a backend. Incoming urls matching the path are forwarded to the backend.",
	"path":     "path is matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional \"path\" part of a URL as defined by RFC 3986. Paths must begin with a '/' and must be present when using PathType with value \"Exact\" or \"Prefix\".",
	"pathType": "pathType determines the interpretation of the path matching. PathType can be one of the following values: * Exact: Matches the URL path exactly. * Prefix: Matches based on a URL path prefix split by '/'. Matching is\n  done on a path element by element basis. A path element refers is the\n  list of labels in the path split by the '/' separator. A request is a\n  match for path p if every p is an element-wise prefix of p of the\n  request path. Note that if the last element of the path is a substring\n  of the last element in request path, it is not a match (e.g. /foo/bar\n  matches /foo/bar/baz, but does not match /foo/barbaz).\n* ImplementationSpecific: Interpretation of the Path matching is up to\n  the IngressClass. Implementations can treat this as a separate PathType\n  or treat it identically to Prefix or Exact path types.\nImplementations are required to support all path types. Defaults to ImplementationSpecific.",
	"backend":  "backend defines the referenced service endpoint to which the traffic will be forwarded to.",
}

func (HTTPIngressPath) SwaggerDoc() map[string]string {
	return map_HTTPIngressPath
}

var map_HTTPIngressRuleValue = map[string]string{
	"":      "HTTPIngressRuleValue is a list of http selectors pointing to backends. In the example: http://<host>/<path>?<searchpart> -> backend where where parts of the url correspond to RFC 3986, this resource will be used to match against everything after the last '/' and before the first '?' or '#'.",
	"paths": "paths is a collection of paths that map requests to backends.",
}

func (HTTPIngressRuleValue) SwaggerDoc() map[string]string {
	return map_HTTPIngressRuleValue
}

var map_Ingress = map[string]string{
	"":         "Ingress is a collection of rules that allow inbound connections to reach the endpoints defined by a backend. An Ingress can be configured to give services externally-reachable urls, load balance traffic, terminate SSL, offer name based virtual hosting etc.",
	"metadata": "Standard object's metadata. More info: https://git.ck-kube/kubernetes/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "spec is the desired state of the Ingress. More info: https://git.ck-kube/kubernetes/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
	"status":   "status is the current state of the Ingress. More info: https://git.ck-kube/kubernetes/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
}

func (Ingress) SwaggerDoc() map[string]string {
	return map_Ingress
}

var map_IngressBackend = map[string]string{
	"":            "IngressBackend describes all endpoints for a given service and port.",
	"serviceName": "serviceName specifies the name of the referenced service.",
	"servicePort": "servicePort Specifies the port of the referenced service.",
	"resource":    "resource is an ObjectRef to another Kubernetes resource in the namespace of the Ingress object. If resource is specified, serviceName and servicePort must not be specified.",
}

func (IngressBackend) SwaggerDoc() map[string]string {
	return map_IngressBackend
}

var map_IngressClass = map[string]string{
	"":         "IngressClass represents the class of the Ingress, referenced by the Ingress Spec. The `ingressclass.kubernetes.io/is-default-class` annotation can be used to indicate that an IngressClass should be considered default. When a single IngressClass resource has this annotation set to true, new Ingress resources without a class specified will be assigned this default class.",
	"metadata": "Standard object's metadata. More info: https://git.ck-kube/kubernetes/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "spec is the desired state of the IngressClass. More info: https://git.ck-kube/kubernetes/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
}

func (IngressClass) SwaggerDoc() map[string]string {
	return map_IngressClass
}

var map_IngressClassList = map[string]string{
	"":         "IngressClassList is a collection of IngressClasses.",
	"metadata": "Standard list metadata.",
	"items":    "items is the list of IngressClasses.",
}

func (IngressClassList) SwaggerDoc() map[string]string {
	return map_IngressClassList
}

var map_IngressClassParametersReference = map[string]string{
	"":          "IngressClassParametersReference identifies an API object. This can be used to specify a cluster or namespace-scoped resource.",
	"apiGroup":  "apiGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
	"kind":      "kind is the type of resource being referenced.",
	"name":      "name is the name of resource being referenced.",
	"scope":     "scope represents if this refers to a cluster or namespace scoped resource. This may be set to \"Cluster\" (default) or \"Namespace\".",
	"namespace": "namespace is the namespace of the resource being referenced. This field is required when scope is set to \"Namespace\" and must be unset when scope is set to \"Cluster\".",
}

func (IngressClassParametersReference) SwaggerDoc() map[string]string {
	return map_IngressClassParametersReference
}

var map_IngressClassSpec = map[string]string{
	"":           "IngressClassSpec provides information about the class of an Ingress.",
	"controller": "controller refers to the name of the controller that should handle this class. This allows for different \"flavors\" that are controlled by the same controller. For example, you may have different parameters for the same implementing controller. This should be specified as a domain-prefixed path no more than 250 characters in length, e.g. \"acme.io/ingress-controller\". This field is immutable.",
	"parameters": "parameters is a link to a custom resource containing additional configuration for the controller. This is optional if the controller does not require extra parameters.",
}

func (IngressClassSpec) SwaggerDoc() map[string]string {
	return map_IngressClassSpec
}

var map_IngressList = map[string]string{
	"":         "IngressList is a collection of Ingress.",
	"metadata": "Standard object's metadata. More info: https://git.ck-kube/kubernetes/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"items":    "items is the list of Ingress.",
}

func (IngressList) SwaggerDoc() map[string]string {
	return map_IngressList
}

var map_IngressLoadBalancerIngress = map[string]string{
	"":         "IngressLoadBalancerIngress represents the status of a load-balancer ingress point.",
	"ip":       "ip is set for load-balancer ingress points that are IP based.",
	"hostname": "hostname is set for load-balancer ingress points that are DNS based.",
	"ports":    "ports provides information about the ports exposed by this LoadBalancer.",
}

func (IngressLoadBalancerIngress) SwaggerDoc() map[string]string {
	return map_IngressLoadBalancerIngress
}

var map_IngressLoadBalancerStatus = map[string]string{
	"":        "LoadBalancerStatus represents the status of a load-balancer.",
	"ingress": "ingress is a list containing ingress points for the load-balancer.",
}

func (IngressLoadBalancerStatus) SwaggerDoc() map[string]string {
	return map_IngressLoadBalancerStatus
}

var map_IngressPortStatus = map[string]string{
	"":         "IngressPortStatus represents the error condition of a service port",
	"port":     "port is the port number of the ingress port.",
	"protocol": "protocol is the protocol of the ingress port. The supported values are: \"TCP\", \"UDP\", \"SCTP\"",
	"error":    "error is to record the problem with the service port The format of the error shall comply with the following rules: - built-in error values shall be specified in this file and those shall use\n  CamelCase names\n- cloud provider specific error values must have names that comply with the\n  format foo.example.com/CamelCase.",
}

func (IngressPortStatus) SwaggerDoc() map[string]string {
	return map_IngressPortStatus
}

var map_IngressRule = map[string]string{
	"":     "IngressRule represents the rules mapping the paths under a specified host to the related backend services. Incoming requests are first evaluated for a host match, then routed to the backend associated with the matching IngressRuleValue.",
	"host": "host is the fully qualified domain name of a network host, as defined by RFC 3986. Note the following deviations from the \"host\" part of the URI as defined in RFC 3986: 1. IPs are not allowed. Currently an IngressRuleValue can only apply to\n   the IP in the Spec of the parent Ingress.\n2. The `:` delimiter is not respected because ports are not allowed.\n\t  Currently the port of an Ingress is implicitly :80 for http and\n\t  :443 for https.\nBoth these may change in the future. Incoming requests are matched against the host before the IngressRuleValue. If the host is unspecified, the Ingress routes all traffic based on the specified IngressRuleValue.\n\nhost can be \"precise\" which is a domain name without the terminating dot of a network host (e.g. \"foo.bar.com\") or \"wildcard\", which is a domain name prefixed with a single wildcard label (e.g. \"*.foo.com\"). The wildcard character '*' must appear by itself as the first DNS label and matches only a single label. You cannot have a wildcard label by itself (e.g. Host == \"*\"). Requests will be matched against the Host field in the following way: 1. If Host is precise, the request matches this rule if the http host header is equal to Host. 2. If Host is a wildcard, then the request matches this rule if the http host header is to equal to the suffix (removing the first label) of the wildcard rule.",
}

func (IngressRule) SwaggerDoc() map[string]string {
	return map_IngressRule
}

var map_IngressRuleValue = map[string]string{
	"": "IngressRuleValue represents a rule to apply against incoming requests. If the rule is satisfied, the request is routed to the specified backend. Currently mixing different types of rules in a single Ingress is disallowed, so exactly one of the following must be set.",
}

func (IngressRuleValue) SwaggerDoc() map[string]string {
	return map_IngressRuleValue
}

var map_IngressSpec = map[string]string{
	"":                 "IngressSpec describes the Ingress the user wishes to exist.",
	"ingressClassName": "ingressClassName is the name of the IngressClass cluster resource. The associated IngressClass defines which controller will implement the resource. This replaces the deprecated `kubernetes.io/ingress.class` annotation. For backwards compatibility, when that annotation is set, it must be given precedence over this field. The controller may emit a warning if the field and annotation have different values. Implementations of this API should ignore Ingresses without a class specified. An IngressClass resource may be marked as default, which can be used to set a default value for this field. For more information, refer to the IngressClass documentation.",
	"backend":          "backend is the default backend capable of servicing requests that don't match any rule. At least one of 'backend' or 'rules' must be specified. This field is optional to allow the loadbalancer controller or defaulting logic to specify a global default.",
	"tls":              "tls represents the TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.",
	"rules":            "rules is a list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend.",
}

func (IngressSpec) SwaggerDoc() map[string]string {
	return map_IngressSpec
}

var map_IngressStatus = map[string]string{
	"":             "IngressStatus describes the current state of the Ingress.",
	"loadBalancer": "loadBalancer contains the current status of the load-balancer.",
}

func (IngressStatus) SwaggerDoc() map[string]string {
	return map_IngressStatus
}

var map_IngressTLS = map[string]string{
	"":           "IngressTLS describes the transport layer security associated with an Ingress.",
	"hosts":      "hosts is a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",
	"secretName": "secretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the \"Host\" header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.",
}

func (IngressTLS) SwaggerDoc() map[string]string {
	return map_IngressTLS
}

// AUTO-GENERATED FUNCTIONS END HERE