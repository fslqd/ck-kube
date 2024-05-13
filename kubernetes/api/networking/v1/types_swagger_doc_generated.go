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

package v1

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
	"pathType": "pathType determines the interpretation of the path matching. PathType can be one of the following values: * Exact: Matches the URL path exactly. * Prefix: Matches based on a URL path prefix split by '/'. Matching is\n  done on a path element by element basis. A path element refers is the\n  list of labels in the path split by the '/' separator. A request is a\n  match for path p if every p is an element-wise prefix of p of the\n  request path. Note that if the last element of the path is a substring\n  of the last element in request path, it is not a match (e.g. /foo/bar\n  matches /foo/bar/baz, but does not match /foo/barbaz).\n* ImplementationSpecific: Interpretation of the Path matching is up to\n  the IngressClass. Implementations can treat this as a separate PathType\n  or treat it identically to Prefix or Exact path types.\nImplementations are required to support all path types.",
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

var map_IPBlock = map[string]string{
	"":       "IPBlock describes a particular CIDR (Ex. \"192.168.1.0/24\",\"2001:db8::/64\") that is allowed to the pods matched by a NetworkPolicySpec's podSelector. The except entry describes CIDRs that should not be included within this rule.",
	"cidr":   "cidr is a string representing the IPBlock Valid examples are \"192.168.1.0/24\" or \"2001:db8::/64\"",
	"except": "except is a slice of CIDRs that should not be included within an IPBlock Valid examples are \"192.168.1.0/24\" or \"2001:db8::/64\" Except values will be rejected if they are outside the cidr range",
}

func (IPBlock) SwaggerDoc() map[string]string {
	return map_IPBlock
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
	"":         "IngressBackend describes all endpoints for a given service and port.",
	"service":  "service references a service as a backend. This is a mutually exclusive setting with \"Resource\".",
	"resource": "resource is an ObjectRef to another Kubernetes resource in the namespace of the Ingress object. If resource is specified, a service.Name and service.Port must not be specified. This is a mutually exclusive setting with \"Service\".",
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
	"":        "IngressLoadBalancerStatus represents the status of a load-balancer.",
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
	"host": "host is the fully qualified domain name of a network host, as defined by RFC 3986. Note the following deviations from the \"host\" part of the URI as defined in RFC 3986: 1. IPs are not allowed. Currently an IngressRuleValue can only apply to\n   the IP in the Spec of the parent Ingress.\n2. The `:` delimiter is not respected because ports are not allowed.\n\t  Currently the port of an Ingress is implicitly :80 for http and\n\t  :443 for https.\nBoth these may change in the future. Incoming requests are matched against the host before the IngressRuleValue. If the host is unspecified, the Ingress routes all traffic based on the specified IngressRuleValue.\n\nhost can be \"precise\" which is a domain name without the terminating dot of a network host (e.g. \"foo.bar.com\") or \"wildcard\", which is a domain name prefixed with a single wildcard label (e.g. \"*.foo.com\"). The wildcard character '*' must appear by itself as the first DNS label and matches only a single label. You cannot have a wildcard label by itself (e.g. Host == \"*\"). Requests will be matched against the Host field in the following way: 1. If host is precise, the request matches this rule if the http host header is equal to Host. 2. If host is a wildcard, then the request matches this rule if the http host header is to equal to the suffix (removing the first label) of the wildcard rule.",
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

var map_IngressServiceBackend = map[string]string{
	"":     "IngressServiceBackend references a Kubernetes Service as a Backend.",
	"name": "name is the referenced service. The service must exist in the same namespace as the Ingress object.",
	"port": "port of the referenced service. A port name or port number is required for a IngressServiceBackend.",
}

func (IngressServiceBackend) SwaggerDoc() map[string]string {
	return map_IngressServiceBackend
}

var map_IngressSpec = map[string]string{
	"":                 "IngressSpec describes the Ingress the user wishes to exist.",
	"ingressClassName": "ingressClassName is the name of an IngressClass cluster resource. Ingress controller implementations use this field to know whether they should be serving this Ingress resource, by a transitive connection (controller -> IngressClass -> Ingress resource). Although the `kubernetes.io/ingress.class` annotation (simple constant name) was never formally defined, it was widely supported by Ingress controllers to create a direct binding between Ingress controller and Ingress resources. Newly created Ingress resources should prefer using the field. However, even though the annotation is officially deprecated, for backwards compatibility reasons, ingress controllers should still honor that annotation if present.",
	"defaultBackend":   "defaultBackend is the backend that should handle requests that don't match any rule. If Rules are not specified, DefaultBackend must be specified. If DefaultBackend is not set, the handling of requests that do not match any of the rules will be up to the Ingress controller.",
	"tls":              "tls represents the TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.",
	"rules":            "rules is a list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend.",
}

func (IngressSpec) SwaggerDoc() map[string]string {
	return map_IngressSpec
}

var map_IngressStatus = map[string]string{
	"":             "IngressStatus describe the current state of the Ingress.",
	"loadBalancer": "loadBalancer contains the current status of the load-balancer.",
}

func (IngressStatus) SwaggerDoc() map[string]string {
	return map_IngressStatus
}

var map_IngressTLS = map[string]string{
	"":           "IngressTLS describes the transport layer security associated with an ingress.",
	"hosts":      "hosts is a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",
	"secretName": "secretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the \"Host\" header field used by an IngressRule, the SNI host is used for termination and value of the \"Host\" header is used for routing.",
}

func (IngressTLS) SwaggerDoc() map[string]string {
	return map_IngressTLS
}

var map_NetworkPolicy = map[string]string{
	"":         "NetworkPolicy describes what network traffic is allowed for a set of Pods",
	"metadata": "Standard object's metadata. More info: https://git.ck-kube/kubernetes/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "spec represents the specification of the desired behavior for this NetworkPolicy.",
}

func (NetworkPolicy) SwaggerDoc() map[string]string {
	return map_NetworkPolicy
}

var map_NetworkPolicyEgressRule = map[string]string{
	"":      "NetworkPolicyEgressRule describes a particular set of traffic that is allowed out of pods matched by a NetworkPolicySpec's podSelector. The traffic must match both ports and to. This type is beta-level in 1.8",
	"ports": "ports is a list of destination ports for outgoing traffic. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list.",
	"to":    "to is a list of destinations for outgoing traffic of pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all destinations (traffic not restricted by destination). If this field is present and contains at least one item, this rule allows traffic only if the traffic matches at least one item in the to list.",
}

func (NetworkPolicyEgressRule) SwaggerDoc() map[string]string {
	return map_NetworkPolicyEgressRule
}

var map_NetworkPolicyIngressRule = map[string]string{
	"":      "NetworkPolicyIngressRule describes a particular set of traffic that is allowed to the pods matched by a NetworkPolicySpec's podSelector. The traffic must match both ports and from.",
	"ports": "ports is a list of ports which should be made accessible on the pods selected for this rule. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list.",
	"from":  "from is a list of sources which should be able to access the pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all sources (traffic not restricted by source). If this field is present and contains at least one item, this rule allows traffic only if the traffic matches at least one item in the from list.",
}

func (NetworkPolicyIngressRule) SwaggerDoc() map[string]string {
	return map_NetworkPolicyIngressRule
}

var map_NetworkPolicyList = map[string]string{
	"":         "NetworkPolicyList is a list of NetworkPolicy objects.",
	"metadata": "Standard list metadata. More info: https://git.ck-kube/kubernetes/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"items":    "items is a list of schema objects.",
}

func (NetworkPolicyList) SwaggerDoc() map[string]string {
	return map_NetworkPolicyList
}

var map_NetworkPolicyPeer = map[string]string{
	"":                  "NetworkPolicyPeer describes a peer to allow traffic to/from. Only certain combinations of fields are allowed",
	"podSelector":       "podSelector is a label selector which selects pods. This field follows standard label selector semantics; if present but empty, it selects all pods.\n\nIf namespaceSelector is also set, then the NetworkPolicyPeer as a whole selects the pods matching podSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects the pods matching podSelector in the policy's own namespace.",
	"namespaceSelector": "namespaceSelector selects namespaces using cluster-scoped labels. This field follows standard label selector semantics; if present but empty, it selects all namespaces.\n\nIf podSelector is also set, then the NetworkPolicyPeer as a whole selects the pods matching podSelector in the namespaces selected by namespaceSelector. Otherwise it selects all pods in the namespaces selected by namespaceSelector.",
	"ipBlock":           "ipBlock defines policy on a particular IPBlock. If this field is set then neither of the other fields can be.",
}

func (NetworkPolicyPeer) SwaggerDoc() map[string]string {
	return map_NetworkPolicyPeer
}

var map_NetworkPolicyPort = map[string]string{
	"":         "NetworkPolicyPort describes a port to allow traffic on",
	"protocol": "protocol represents the protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.",
	"port":     "port represents the port on the given protocol. This can either be a numerical or named port on a pod. If this field is not provided, this matches all port names and numbers. If present, only traffic on the specified protocol AND port will be matched.",
	"endPort":  "endPort indicates that the range of ports from port to endPort if set, inclusive, should be allowed by the policy. This field cannot be defined if the port field is not defined or if the port field is defined as a named (string) port. The endPort must be equal or greater than port.",
}

func (NetworkPolicyPort) SwaggerDoc() map[string]string {
	return map_NetworkPolicyPort
}

var map_NetworkPolicySpec = map[string]string{
	"":            "NetworkPolicySpec provides the specification of a NetworkPolicy",
	"podSelector": "podSelector selects the pods to which this NetworkPolicy object applies. The array of ingress rules is applied to any pods selected by this field. Multiple network policies can select the same set of pods. In this case, the ingress rules for each are combined additively. This field is NOT optional and follows standard label selector semantics. An empty podSelector matches all pods in this namespace.",
	"ingress":     "ingress is a list of ingress rules to be applied to the selected pods. Traffic is allowed to a pod if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic source is the pod's local node, OR if the traffic matches at least one ingress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this field is empty then this NetworkPolicy does not allow any traffic (and serves solely to ensure that the pods it selects are isolated by default)",
	"egress":      "egress is a list of egress rules to be applied to the selected pods. Outgoing traffic is allowed if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic matches at least one egress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this field is empty then this NetworkPolicy limits all outgoing traffic (and serves solely to ensure that the pods it selects are isolated by default). This field is beta-level in 1.8",
	"policyTypes": "policyTypes is a list of rule types that the NetworkPolicy relates to. Valid options are [\"Ingress\"], [\"Egress\"], or [\"Ingress\", \"Egress\"]. If this field is not specified, it will default based on the existence of ingress or egress rules; policies that contain an egress section are assumed to affect egress, and all policies (whether or not they contain an ingress section) are assumed to affect ingress. If you want to write an egress-only policy, you must explicitly specify policyTypes [ \"Egress\" ]. Likewise, if you want to write a policy that specifies that no egress is allowed, you must specify a policyTypes value that include \"Egress\" (since such a policy would not include an egress section and would otherwise default to just [ \"Ingress\" ]). This field is beta-level in 1.8",
}

func (NetworkPolicySpec) SwaggerDoc() map[string]string {
	return map_NetworkPolicySpec
}

var map_ServiceBackendPort = map[string]string{
	"":       "ServiceBackendPort is the service port being referenced.",
	"name":   "name is the name of the port on the Service. This is a mutually exclusive setting with \"Number\".",
	"number": "number is the numerical port number (e.g. 80) on the Service. This is a mutually exclusive setting with \"Name\".",
}

func (ServiceBackendPort) SwaggerDoc() map[string]string {
	return map_ServiceBackendPort
}

// AUTO-GENERATED FUNCTIONS END HERE