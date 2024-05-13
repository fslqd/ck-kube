package v1

import metav1 "ck-kube/kubernetes/apis/meta/v1"

type Lease struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              LeaseSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

type LeaseSpec struct {
	HolderIdentity       *string           `json:"holderIdentity,omitempty" protobuf:"bytes,1,opt,name=holderIdentity"`
	LeaseDurationSeconds *int32            `json:"leaseDurationSeconds,omitempty" protobuf:"varint,2,opt,name=leaseDurationSeconds"`
	AcquireTime          *metav1.MicroTime `json:"acquireTime,omitempty" protobuf:"bytes,3,opt,name=acquireTime"`
	RenewTime            *metav1.MicroTime `json:"renewTime,omitempty" protobuf:"bytes,4,opt,name=renewTime"`
	LeaseTransitions     *int32            `json:"leaseTransitions,omitempty" protobuf:"varint,5,opt,name=leaseTransitions"`
}
