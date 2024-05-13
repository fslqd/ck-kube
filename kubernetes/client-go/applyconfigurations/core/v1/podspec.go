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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	corev1 "github.com/fslqd/ck-kube/kubernetes/api/core/v1"
)

// PodSpecApplyConfiguration represents an declarative configuration of the PodSpec type for use
// with apply.
type PodSpecApplyConfiguration struct {
	Volumes                       []VolumeApplyConfiguration                   `json:"volumes,omitempty"`
	InitContainers                []ContainerApplyConfiguration                `json:"initContainers,omitempty"`
	Containers                    []ContainerApplyConfiguration                `json:"containers,omitempty"`
	EphemeralContainers           []EphemeralContainerApplyConfiguration       `json:"ephemeralContainers,omitempty"`
	RestartPolicy                 *corev1.RestartPolicy                        `json:"restartPolicy,omitempty"`
	TerminationGracePeriodSeconds *int64                                       `json:"terminationGracePeriodSeconds,omitempty"`
	ActiveDeadlineSeconds         *int64                                       `json:"activeDeadlineSeconds,omitempty"`
	DNSPolicy                     *corev1.DNSPolicy                            `json:"dnsPolicy,omitempty"`
	NodeSelector                  map[string]string                            `json:"nodeSelector,omitempty"`
	ServiceAccountName            *string                                      `json:"serviceAccountName,omitempty"`
	DeprecatedServiceAccount      *string                                      `json:"serviceAccount,omitempty"`
	AutomountServiceAccountToken  *bool                                        `json:"automountServiceAccountToken,omitempty"`
	NodeName                      *string                                      `json:"nodeName,omitempty"`
	HostNetwork                   *bool                                        `json:"hostNetwork,omitempty"`
	HostPID                       *bool                                        `json:"hostPID,omitempty"`
	HostIPC                       *bool                                        `json:"hostIPC,omitempty"`
	ShareProcessNamespace         *bool                                        `json:"shareProcessNamespace,omitempty"`
	SecurityContext               *PodSecurityContextApplyConfiguration        `json:"securityContext,omitempty"`
	ImagePullSecrets              []LocalObjectReferenceApplyConfiguration     `json:"imagePullSecrets,omitempty"`
	Hostname                      *string                                      `json:"hostname,omitempty"`
	Subdomain                     *string                                      `json:"subdomain,omitempty"`
	Affinity                      *AffinityApplyConfiguration                  `json:"affinity,omitempty"`
	SchedulerName                 *string                                      `json:"schedulerName,omitempty"`
	Tolerations                   []TolerationApplyConfiguration               `json:"tolerations,omitempty"`
	HostAliases                   []HostAliasApplyConfiguration                `json:"hostAliases,omitempty"`
	PriorityClassName             *string                                      `json:"priorityClassName,omitempty"`
	Priority                      *int32                                       `json:"priority,omitempty"`
	DNSConfig                     *PodDNSConfigApplyConfiguration              `json:"dnsConfig,omitempty"`
	ReadinessGates                []PodReadinessGateApplyConfiguration         `json:"readinessGates,omitempty"`
	RuntimeClassName              *string                                      `json:"runtimeClassName,omitempty"`
	EnableServiceLinks            *bool                                        `json:"enableServiceLinks,omitempty"`
	PreemptionPolicy              *corev1.PreemptionPolicy                     `json:"preemptionPolicy,omitempty"`
	Overhead                      *corev1.ResourceList                         `json:"overhead,omitempty"`
	TopologySpreadConstraints     []TopologySpreadConstraintApplyConfiguration `json:"topologySpreadConstraints,omitempty"`
	SetHostnameAsFQDN             *bool                                        `json:"setHostnameAsFQDN,omitempty"`
	OS                            *PodOSApplyConfiguration                     `json:"os,omitempty"`
	HostUsers                     *bool                                        `json:"hostUsers,omitempty"`
	SchedulingGates               []PodSchedulingGateApplyConfiguration        `json:"schedulingGates,omitempty"`
	ResourceClaims                []PodResourceClaimApplyConfiguration         `json:"resourceClaims,omitempty"`
}

// PodSpecApplyConfiguration constructs an declarative configuration of the PodSpec type for use with
// apply.
func PodSpec() *PodSpecApplyConfiguration {
	return &PodSpecApplyConfiguration{}
}

// WithVolumes adds the given value to the Volumes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Volumes field.
func (b *PodSpecApplyConfiguration) WithVolumes(values ...*VolumeApplyConfiguration) *PodSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithVolumes")
		}
		b.Volumes = append(b.Volumes, *values[i])
	}
	return b
}

// WithInitContainers adds the given value to the InitContainers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the InitContainers field.
func (b *PodSpecApplyConfiguration) WithInitContainers(values ...*ContainerApplyConfiguration) *PodSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithInitContainers")
		}
		b.InitContainers = append(b.InitContainers, *values[i])
	}
	return b
}

// WithContainers adds the given value to the Containers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Containers field.
func (b *PodSpecApplyConfiguration) WithContainers(values ...*ContainerApplyConfiguration) *PodSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithContainers")
		}
		b.Containers = append(b.Containers, *values[i])
	}
	return b
}

// WithEphemeralContainers adds the given value to the EphemeralContainers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the EphemeralContainers field.
func (b *PodSpecApplyConfiguration) WithEphemeralContainers(values ...*EphemeralContainerApplyConfiguration) *PodSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithEphemeralContainers")
		}
		b.EphemeralContainers = append(b.EphemeralContainers, *values[i])
	}
	return b
}

// WithRestartPolicy sets the RestartPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RestartPolicy field is set to the value of the last call.
func (b *PodSpecApplyConfiguration) WithRestartPolicy(value corev1.RestartPolicy) *PodSpecApplyConfiguration {
	b.RestartPolicy = &value
	return b
}

// WithTerminationGracePeriodSeconds sets the TerminationGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TerminationGracePeriodSeconds field is set to the value of the last call.
func (b *PodSpecApplyConfiguration) WithTerminationGracePeriodSeconds(value int64) *PodSpecApplyConfiguration {
	b.TerminationGracePeriodSeconds = &value
	return b
}

// WithActiveDeadlineSeconds sets the ActiveDeadlineSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ActiveDeadlineSeconds field is set to the value of the last call.
func (b *PodSpecApplyConfiguration) WithActiveDeadlineSeconds(value int64) *PodSpecApplyConfiguration {
	b.ActiveDeadlineSeconds = &value
	return b
}

// WithDNSPolicy sets the DNSPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DNSPolicy field is set to the value of the last call.
func (b *PodSpecApplyConfiguration) WithDNSPolicy(value corev1.DNSPolicy) *PodSpecApplyConfiguration {
	b.DNSPolicy = &value
	return b
}

// WithNodeSelector puts the entries into the NodeSelector field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the NodeSelector field,
// overwriting an existing map entries in NodeSelector field with the same key.
func (b *PodSpecApplyConfiguration) WithNodeSelector(entries map[string]string) *PodSpecApplyConfiguration {
	if b.NodeSelector == nil && len(entries) > 0 {
		b.NodeSelector = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.NodeSelector[k] = v
	}
	return b
}

// WithServiceAccountName sets the ServiceAccountName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServiceAccountName field is set to the value of the last call.
func (b *PodSpecApplyConfiguration) WithServiceAccountName(value string) *PodSpecApplyConfiguration {
	b.ServiceAccountName = &value
	return b
}

// WithDeprecatedServiceAccount sets the DeprecatedServiceAccount field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeprecatedServiceAccount field is set to the value of the last call.
func (b *PodSpecApplyConfiguration) WithDeprecatedServiceAccount(value string) *PodSpecApplyConfiguration {
	b.DeprecatedServiceAccount = &value
	return b
}

// WithAutomountServiceAccountToken sets the AutomountServiceAccountToken field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AutomountServiceAccountToken field is set to the value of the last call.
func (b *PodSpecApplyConfiguration) WithAutomountServiceAccountToken(value bool) *PodSpecApplyConfiguration {
	b.AutomountServiceAccountToken = &value
	return b
}

// WithNodeName sets the NodeName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeName field is set to the value of the last call.
func (b *PodSpecApplyConfiguration) WithNodeName(value string) *PodSpecApplyConfiguration {
	b.NodeName = &value
	return b
}

// WithHostNetwork sets the HostNetwork field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HostNetwork field is set to the value of the last call.
func (b *PodSpecApplyConfiguration) WithHostNetwork(value bool) *PodSpecApplyConfiguration {
	b.HostNetwork = &value
	return b
}

// WithHostPID sets the HostPID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HostPID field is set to the value of the last call.
func (b *PodSpecApplyConfiguration) WithHostPID(value bool) *PodSpecApplyConfiguration {
	b.HostPID = &value
	return b
}

// WithHostIPC sets the HostIPC field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HostIPC field is set to the value of the last call.
func (b *PodSpecApplyConfiguration) WithHostIPC(value bool) *PodSpecApplyConfiguration {
	b.HostIPC = &value
	return b
}

// WithShareProcessNamespace sets the ShareProcessNamespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ShareProcessNamespace field is set to the value of the last call.
func (b *PodSpecApplyConfiguration) WithShareProcessNamespace(value bool) *PodSpecApplyConfiguration {
	b.ShareProcessNamespace = &value
	return b
}

// WithSecurityContext sets the SecurityContext field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecurityContext field is set to the value of the last call.
func (b *PodSpecApplyConfiguration) WithSecurityContext(value *PodSecurityContextApplyConfiguration) *PodSpecApplyConfiguration {
	b.SecurityContext = value
	return b
}

// WithImagePullSecrets adds the given value to the ImagePullSecrets field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ImagePullSecrets field.
func (b *PodSpecApplyConfiguration) WithImagePullSecrets(values ...*LocalObjectReferenceApplyConfiguration) *PodSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithImagePullSecrets")
		}
		b.ImagePullSecrets = append(b.ImagePullSecrets, *values[i])
	}
	return b
}

// WithHostname sets the Hostname field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Hostname field is set to the value of the last call.
func (b *PodSpecApplyConfiguration) WithHostname(value string) *PodSpecApplyConfiguration {
	b.Hostname = &value
	return b
}

// WithSubdomain sets the Subdomain field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Subdomain field is set to the value of the last call.
func (b *PodSpecApplyConfiguration) WithSubdomain(value string) *PodSpecApplyConfiguration {
	b.Subdomain = &value
	return b
}

// WithAffinity sets the Affinity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Affinity field is set to the value of the last call.
func (b *PodSpecApplyConfiguration) WithAffinity(value *AffinityApplyConfiguration) *PodSpecApplyConfiguration {
	b.Affinity = value
	return b
}

// WithSchedulerName sets the SchedulerName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SchedulerName field is set to the value of the last call.
func (b *PodSpecApplyConfiguration) WithSchedulerName(value string) *PodSpecApplyConfiguration {
	b.SchedulerName = &value
	return b
}

// WithTolerations adds the given value to the Tolerations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tolerations field.
func (b *PodSpecApplyConfiguration) WithTolerations(values ...*TolerationApplyConfiguration) *PodSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTolerations")
		}
		b.Tolerations = append(b.Tolerations, *values[i])
	}
	return b
}

// WithHostAliases adds the given value to the HostAliases field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the HostAliases field.
func (b *PodSpecApplyConfiguration) WithHostAliases(values ...*HostAliasApplyConfiguration) *PodSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithHostAliases")
		}
		b.HostAliases = append(b.HostAliases, *values[i])
	}
	return b
}

// WithPriorityClassName sets the PriorityClassName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PriorityClassName field is set to the value of the last call.
func (b *PodSpecApplyConfiguration) WithPriorityClassName(value string) *PodSpecApplyConfiguration {
	b.PriorityClassName = &value
	return b
}

// WithPriority sets the Priority field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Priority field is set to the value of the last call.
func (b *PodSpecApplyConfiguration) WithPriority(value int32) *PodSpecApplyConfiguration {
	b.Priority = &value
	return b
}

// WithDNSConfig sets the DNSConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DNSConfig field is set to the value of the last call.
func (b *PodSpecApplyConfiguration) WithDNSConfig(value *PodDNSConfigApplyConfiguration) *PodSpecApplyConfiguration {
	b.DNSConfig = value
	return b
}

// WithReadinessGates adds the given value to the ReadinessGates field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ReadinessGates field.
func (b *PodSpecApplyConfiguration) WithReadinessGates(values ...*PodReadinessGateApplyConfiguration) *PodSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithReadinessGates")
		}
		b.ReadinessGates = append(b.ReadinessGates, *values[i])
	}
	return b
}

// WithRuntimeClassName sets the RuntimeClassName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RuntimeClassName field is set to the value of the last call.
func (b *PodSpecApplyConfiguration) WithRuntimeClassName(value string) *PodSpecApplyConfiguration {
	b.RuntimeClassName = &value
	return b
}

// WithEnableServiceLinks sets the EnableServiceLinks field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnableServiceLinks field is set to the value of the last call.
func (b *PodSpecApplyConfiguration) WithEnableServiceLinks(value bool) *PodSpecApplyConfiguration {
	b.EnableServiceLinks = &value
	return b
}

// WithPreemptionPolicy sets the PreemptionPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PreemptionPolicy field is set to the value of the last call.
func (b *PodSpecApplyConfiguration) WithPreemptionPolicy(value corev1.PreemptionPolicy) *PodSpecApplyConfiguration {
	b.PreemptionPolicy = &value
	return b
}

// WithOverhead sets the Overhead field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Overhead field is set to the value of the last call.
func (b *PodSpecApplyConfiguration) WithOverhead(value corev1.ResourceList) *PodSpecApplyConfiguration {
	b.Overhead = &value
	return b
}

// WithTopologySpreadConstraints adds the given value to the TopologySpreadConstraints field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the TopologySpreadConstraints field.
func (b *PodSpecApplyConfiguration) WithTopologySpreadConstraints(values ...*TopologySpreadConstraintApplyConfiguration) *PodSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTopologySpreadConstraints")
		}
		b.TopologySpreadConstraints = append(b.TopologySpreadConstraints, *values[i])
	}
	return b
}

// WithSetHostnameAsFQDN sets the SetHostnameAsFQDN field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SetHostnameAsFQDN field is set to the value of the last call.
func (b *PodSpecApplyConfiguration) WithSetHostnameAsFQDN(value bool) *PodSpecApplyConfiguration {
	b.SetHostnameAsFQDN = &value
	return b
}

// WithOS sets the OS field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OS field is set to the value of the last call.
func (b *PodSpecApplyConfiguration) WithOS(value *PodOSApplyConfiguration) *PodSpecApplyConfiguration {
	b.OS = value
	return b
}

// WithHostUsers sets the HostUsers field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HostUsers field is set to the value of the last call.
func (b *PodSpecApplyConfiguration) WithHostUsers(value bool) *PodSpecApplyConfiguration {
	b.HostUsers = &value
	return b
}

// WithSchedulingGates adds the given value to the SchedulingGates field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the SchedulingGates field.
func (b *PodSpecApplyConfiguration) WithSchedulingGates(values ...*PodSchedulingGateApplyConfiguration) *PodSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithSchedulingGates")
		}
		b.SchedulingGates = append(b.SchedulingGates, *values[i])
	}
	return b
}

// WithResourceClaims adds the given value to the ResourceClaims field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ResourceClaims field.
func (b *PodSpecApplyConfiguration) WithResourceClaims(values ...*PodResourceClaimApplyConfiguration) *PodSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithResourceClaims")
		}
		b.ResourceClaims = append(b.ResourceClaims, *values[i])
	}
	return b
}
