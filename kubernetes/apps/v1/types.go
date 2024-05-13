package v1

import metav1 "github.com/fslqd/ck-kube/kubernetes/apis/meta/v1"

type DeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Deployment `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type Deployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              DeploymentSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

type DeploymentSpec struct {
	Replicas            *int32 `json:"replicas,omitempty" protobuf:"varint,1,opt,name=replicas"`
	UpdatedReplicas     int32  `json:"updatedReplicas,omitempty" protobuf:"varint,3,opt,name=updatedReplicas"`
	ReadyReplicas       int32  `json:"readyReplicas,omitempty" protobuf:"varint,7,opt,name=readyReplicas"`
	AvailableReplicas   int32  `json:"availableReplicas,omitempty" protobuf:"varint,4,opt,name=availableReplicas"`
	UnavailableReplicas int32  `json:"unavailableReplicas,omitempty" protobuf:"varint,5,opt,name=unavailableReplicas"`
}

type StatefulSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []StatefulSet `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type StatefulSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              StatefulSetSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}
type StatefulSetSpec struct {
	Replicas *int32 `json:"replicas,omitempty" protobuf:"varint,1,opt,name=replicas"`
}

type DaemonSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []DaemonSet `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type DaemonSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              DaemonSetSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            DaemonSetStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type DaemonSetSpec struct {
	MinReadySeconds      int32  `json:"minReadySeconds,omitempty" protobuf:"varint,4,opt,name=minReadySeconds"`
	RevisionHistoryLimit *int32 `json:"revisionHistoryLimit,omitempty" protobuf:"varint,6,opt,name=revisionHistoryLimit"`
}
type DaemonSetStatus struct {
	CurrentNumberScheduled int32 `json:"currentNumberScheduled" protobuf:"varint,1,opt,name=currentNumberScheduled"`
}
