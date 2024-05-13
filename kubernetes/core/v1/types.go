package v1

import (
	"fmt"
	//"ck-kube/kubernetes/api/resource"
	metav1 "ck-kube/kubernetes/apis/meta/v1"
	"strconv"
)

const (
	ResourcePods ResourceName = "pods"
)

const (
	ContainerRestartPolicyAlways ContainerRestartPolicy = "Always"
)

const (
	ResourceCPU              ResourceName = "cpu"
	ResourceMemory           ResourceName = "memory"
	ResourceStorage          ResourceName = "storage"
	ResourceEphemeralStorage ResourceName = "ephemeral-storage"
)

const (
	ResourceHugePagesPrefix = "hugepages-"
)

const (
	PodResizeStatusInfeasible PodResizeStatus = "Infeasible"
)

type NodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Node `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type Node struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              NodeSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            NodeStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type ReplicaSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []ReplicaSet `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type ReplicaSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              ReplicaSetSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            ReplicaSetStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type ReplicaSetSpec struct {
	Replicas *int32 `json:"replicas,omitempty" protobuf:"varint,1,opt,name=replicas"`
}

type ReplicaSetStatus struct {
	Replicas int32 `json:"replicas" protobuf:"varint,1,opt,name=replicas"`
}

type Deployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              DeploymentSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            DeploymentStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type DeploymentStatus struct {
	ObservedGeneration  int64                 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
	Replicas            int32                 `json:"replicas,omitempty" protobuf:"varint,2,opt,name=replicas"`
	UpdatedReplicas     int32                 `json:"updatedReplicas,omitempty" protobuf:"varint,3,opt,name=updatedReplicas"`
	ReadyReplicas       int32                 `json:"readyReplicas,omitempty" protobuf:"varint,7,opt,name=readyReplicas"`
	AvailableReplicas   int32                 `json:"availableReplicas,omitempty" protobuf:"varint,4,opt,name=availableReplicas"`
	UnavailableReplicas int32                 `json:"unavailableReplicas,omitempty" protobuf:"varint,5,opt,name=unavailableReplicas"`
	Conditions          []DeploymentCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,6,rep,name=conditions"`
}

type DeploymentConditionType string
type DeploymentCondition struct {
	Type   DeploymentConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=DeploymentConditionType"`
	Status ConditionStatus         `json:"status" protobuf:"bytes,2,opt,name=status,casttype=k8s.io/api/core/v1.ConditionStatus"`
	Reason string                  `json:"reason,omitempty" protobuf:"bytes,4,opt,name=reason"`
}

type DeploymentSpec struct {
	Replicas        *int32             `json:"replicas,omitempty" protobuf:"varint,1,opt,name=replicas"`
	Strategy        DeploymentStrategy `json:"strategy,omitempty" patchStrategy:"retainKeys" protobuf:"bytes,4,opt,name=strategy"`
	MinReadySeconds int32              `json:"minReadySeconds,omitempty" protobuf:"varint,5,opt,name=minReadySeconds"`
	Template        PodTemplateSpec    `json:"template" protobuf:"bytes,3,opt,name=template"`
}
type PodTemplateSpec struct {
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              PodSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

type DeploymentStrategyType string

type DeploymentStrategy struct {
	Type          DeploymentStrategyType   `json:"type,omitempty" protobuf:"bytes,1,opt,name=type,casttype=DeploymentStrategyType"`
	RollingUpdate *RollingUpdateDeployment `json:"rollingUpdate,omitempty" protobuf:"bytes,2,opt,name=rollingUpdate"`
}

type RollingUpdateDeployment struct {
	MaxUnavailable *IntOrString `json:"maxUnavailable,omitempty" protobuf:"bytes,1,opt,name=maxUnavailable"`
	MaxSurge       *IntOrString `json:"maxSurge,omitempty" protobuf:"bytes,2,opt,name=maxSurge"`
}

type IntOrString struct {
	Type   Type   `protobuf:"varint,1,opt,name=type,casttype=Type"`
	IntVal int32  `protobuf:"varint,2,opt,name=intVal"`
	StrVal string `protobuf:"bytes,3,opt,name=strVal"`
}

const (
	Int    Type = iota // The IntOrString holds an int.
	String             // The IntOrString holds a string.
)

func (intstr *IntOrString) String() string {
	if intstr == nil {
		return "<nil>"
	}
	if intstr.Type == String {
		return intstr.StrVal
	}
	return strconv.Itoa(intstr.IntValue())
}

func (intstr *IntOrString) IntValue() int {
	if intstr.Type == String {
		i, _ := strconv.Atoi(intstr.StrVal)
		return i
	}
	return int(intstr.IntVal)
}

type Type int64

type DaemonSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              DaemonSetSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            DaemonSetStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type DaemonSetStatus struct {
	CurrentNumberScheduled int32 `json:"currentNumberScheduled" protobuf:"varint,1,opt,name=currentNumberScheduled"`
}

type DaemonSetSpec struct {
	Selector *metav1.LabelSelector `json:"selector" protobuf:"bytes,1,opt,name=selector"`
}

type NodeSpec struct {
	PodCIDR            string   `json:"podCIDR,omitempty" protobuf:"bytes,1,opt,name=podCIDR"`
	Taints             []Taint  `json:"taints,omitempty" protobuf:"bytes,5,opt,name=taints"`
	PodCIDRs           []string `json:"podCIDRs,omitempty" protobuf:"bytes,7,opt,name=podCIDRs" patchStrategy:"merge"`
	ProviderID         string   `json:"providerID,omitempty" protobuf:"bytes,3,opt,name=providerID"`
	Unschedulable      bool     `json:"unschedulable,omitempty" protobuf:"varint,4,opt,name=unschedulable"`
	DoNotUseExternalID string   `json:"externalID,omitempty" protobuf:"bytes,2,opt,name=externalID"`
}

type NodeStatus struct {
	//Capacity    ResourceList    `json:"capacity,omitempty" protobuf:"bytes,1,rep,name=capacity,casttype=ResourceList,castkey=ResourceName"`
	//Allocatable ResourceList    `json:"allocatable,omitempty" protobuf:"bytes,2,rep,name=allocatable,casttype=ResourceList,castkey=ResourceName"`
	Addresses  []NodeAddress   `json:"addresses,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,5,rep,name=addresses"`
	NodeInfo   NodeSystemInfo  `json:"nodeInfo,omitempty" protobuf:"bytes,7,opt,name=nodeInfo"`
	Conditions []NodeCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,4,rep,name=conditions"`
}
type NodeCondition struct {
	Type               NodeConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=NodeConditionType"`
	Status             ConditionStatus   `json:"status" protobuf:"bytes,2,opt,name=status,casttype=ConditionStatus"`
	LastHeartbeatTime  metav1.Time       `json:"lastHeartbeatTime,omitempty" protobuf:"bytes,3,opt,name=lastHeartbeatTime"`
	LastTransitionTime metav1.Time       `json:"lastTransitionTime,omitempty" protobuf:"bytes,4,opt,name=lastTransitionTime"`
	Reason             string            `json:"reason,omitempty" protobuf:"bytes,5,opt,name=reason"`
	Message            string            `json:"message,omitempty" protobuf:"bytes,6,opt,name=message"`
}

type NodeConditionType string
type ConditionStatus string

type NodeSystemInfo struct {
	MachineID               string `json:"machineID" protobuf:"bytes,1,opt,name=machineID"`
	SystemUUID              string `json:"systemUUID" protobuf:"bytes,2,opt,name=systemUUID"`
	BootID                  string `json:"bootID" protobuf:"bytes,3,opt,name=bootID"`
	KernelVersion           string `json:"kernelVersion" protobuf:"bytes,4,opt,name=kernelVersion"`
	OSImage                 string `json:"osImage" protobuf:"bytes,5,opt,name=osImage"`
	ContainerRuntimeVersion string `json:"containerRuntimeVersion" protobuf:"bytes,6,opt,name=containerRuntimeVersion"`
	KubeletVersion          string `json:"kubeletVersion" protobuf:"bytes,7,opt,name=kubeletVersion"`
	KubeProxyVersion        string `json:"kubeProxyVersion" protobuf:"bytes,8,opt,name=kubeProxyVersion"`
	OperatingSystem         string `json:"operatingSystem" protobuf:"bytes,9,opt,name=operatingSystem"`
	Architecture            string `json:"architecture" protobuf:"bytes,10,opt,name=architecture"`
}

type NodeAddress struct {
	Type    NodeAddressType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=NodeAddressType"`
	Address string          `json:"address" protobuf:"bytes,2,opt,name=address"`
}

type NodeAddressType string

type PodList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Pod `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type Pod struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              PodSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            PodStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type Container struct {
	Name          string                  `json:"name" protobuf:"bytes,1,opt,name=name"`
	Resources     ResourceRequirements    `json:"resources,omitempty" protobuf:"bytes,8,opt,name=resources"`
	RestartPolicy *ContainerRestartPolicy `json:"restartPolicy,omitempty" protobuf:"bytes,24,opt,name=restartPolicy,casttype=ContainerRestartPolicy"`
}
type ContainerRestartPolicy string

type ResourceRequirements struct {
	//Limits   ResourceList `json:"limits,omitempty" protobuf:"bytes,1,rep,name=limits,casttype=ResourceList,castkey=ResourceName"`
	//Requests ResourceList `json:"requests,omitempty" protobuf:"bytes,2,rep,name=requests,casttype=ResourceList,castkey=ResourceName"`
}

type PodSpec struct {
	NodeName       string      `json:"nodeName,omitempty" protobuf:"bytes,10,opt,name=nodeName"`
	Containers     []Container `json:"containers" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,2,rep,name=containers"`
	InitContainers []Container `json:"initContainers,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,20,rep,name=initContainers"`
	//Overhead           ResourceList `json:"overhead,omitempty" protobuf:"bytes,32,opt,name=overhead"`
	ServiceAccountName string   `json:"serviceAccountName,omitempty" protobuf:"bytes,8,opt,name=serviceAccountName"`
	Volumes            []Volume `json:"volumes,omitempty" patchStrategy:"merge,retainKeys" patchMergeKey:"name" protobuf:"bytes,1,rep,name=volumes"`
}

type Volume struct {
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
}

type PodStatus struct {
	ContainerStatuses []ContainerStatus `json:"containerStatuses,omitempty" protobuf:"bytes,8,rep,name=containerStatuses"`
	Resize            PodResizeStatus   `json:"resize,omitempty" protobuf:"bytes,14,opt,name=resize,casttype=PodResizeStatus"`
}
type PodResizeStatus string

type Event struct {
	metav1.TypeMeta     `json:",inline"`
	metav1.ObjectMeta   `json:"metadata" protobuf:"bytes,1,opt,name=metadata"`
	Reason              string           `json:"reason,omitempty" protobuf:"bytes,3,opt,name=reason"`
	Message             string           `json:"message,omitempty" protobuf:"bytes,4,opt,name=message"`
	Source              EventSource      `json:"source,omitempty" protobuf:"bytes,5,opt,name=source"`
	FirstTimestamp      metav1.Time      `json:"firstTimestamp,omitempty" protobuf:"bytes,6,opt,name=firstTimestamp"`
	LastTimestamp       metav1.Time      `json:"lastTimestamp,omitempty" protobuf:"bytes,7,opt,name=lastTimestamp"`
	Count               int32            `json:"count,omitempty" protobuf:"varint,8,opt,name=count"`
	Type                string           `json:"type,omitempty" protobuf:"bytes,9,opt,name=type"`
	EventTime           metav1.MicroTime `json:"eventTime,omitempty" protobuf:"bytes,10,opt,name=eventTime"`
	Series              *EventSeries     `json:"series,omitempty" protobuf:"bytes,11,opt,name=series"`
	ReportingController string           `json:"reportingComponent" protobuf:"bytes,14,opt,name=reportingComponent"`
}

type EventSource struct {
	Component string `json:"component,omitempty" protobuf:"bytes,1,opt,name=component"`
	Host      string `json:"host,omitempty" protobuf:"bytes,2,opt,name=host"`
}

type EventSeries struct {
	Count            int32            `json:"count,omitempty" protobuf:"varint,1,name=count"`
	LastObservedTime metav1.MicroTime `json:"lastObservedTime,omitempty" protobuf:"bytes,2,name=lastObservedTime"`
}

type EventList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Event `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type Taint struct {
	Key    string      `json:"key" protobuf:"bytes,1,opt,name=key"`
	Value  string      `json:"value,omitempty" protobuf:"bytes,2,opt,name=value"`
	Effect TaintEffect `json:"effect" protobuf:"bytes,3,opt,name=effect,casttype=TaintEffect"`
}

func (t *Taint) ToString() string {
	if len(t.Effect) == 0 {
		if len(t.Value) == 0 {
			return fmt.Sprintf("%v", t.Key)
		}
		return fmt.Sprintf("%v=%v:", t.Key, t.Value)
	}
	if len(t.Value) == 0 {
		return fmt.Sprintf("%v:%v", t.Key, t.Effect)
	}
	return fmt.Sprintf("%v=%v:%v", t.Key, t.Value, t.Effect)
}

type TaintEffect string

type ContainerStatus struct {
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
	//AllocatedResources ResourceList `json:"allocatedResources,omitempty" protobuf:"bytes,10,rep,name=allocatedResources,casttype=ResourceList,castkey=ResourceName"`
}
type ResourceName string

func (rn ResourceName) String() string {
	return string(rn)
}
