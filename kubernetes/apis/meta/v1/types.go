package v1

import (
	"time"
)

type TypeMeta struct {
	Kind       string `json:"kind,omitempty" protobuf:"bytes,1,opt,name=kind"`
	APIVersion string `json:"apiVersion,omitempty" protobuf:"bytes,2,opt,name=apiVersion"`
}

type ListMeta struct {
	SelfLink           string `json:"selfLink,omitempty" protobuf:"bytes,1,opt,name=selfLink"`
	ResourceVersion    string `json:"resourceVersion,omitempty" protobuf:"bytes,2,opt,name=resourceVersion"`
	Continue           string `json:"continue,omitempty" protobuf:"bytes,3,opt,name=continue"`
	RemainingItemCount *int64 `json:"remainingItemCount,omitempty" protobuf:"bytes,4,opt,name=remainingItemCount"`
}
type ObjectMeta struct {
	Name                       string            `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	GenerateName               string            `json:"generateName,omitempty" protobuf:"bytes,2,opt,name=generateName"`
	Namespace                  string            `json:"namespace,omitempty" protobuf:"bytes,3,opt,name=namespace"`
	SelfLink                   string            `json:"selfLink,omitempty" protobuf:"bytes,4,opt,name=selfLink"`
	UID                        string            `json:"uid,omitempty" protobuf:"bytes,5,opt,name=uid"`
	ResourceVersion            string            `json:"resourceVersion,omitempty" protobuf:"bytes,6,opt,name=resourceVersion"`
	Generation                 int64             `json:"generation,omitempty" protobuf:"varint,7,opt,name=generation"`
	DeletionGracePeriodSeconds *int64            `json:"deletionGracePeriodSeconds,omitempty" protobuf:"varint,10,opt,name=deletionGracePeriodSeconds"`
	Labels                     map[string]string `json:"labels,omitempty" protobuf:"bytes,11,rep,name=labels"`
	Annotations                map[string]string `json:"annotations,omitempty" protobuf:"bytes,12,rep,name=annotations"`
	Finalizers                 []string          `json:"finalizers,omitempty" patchStrategy:"merge" protobuf:"bytes,14,rep,name=finalizers"`
	CreationTimestamp          Time              `json:"creationTimestamp,omitempty" protobuf:"bytes,8,opt,name=creationTimestamp"`
}

type Time struct {
	time.Time `protobuf:"-"`
}

type MicroTime struct {
	time.Time `protobuf:"-"`
}

type LabelSelector struct {
	MatchLabels map[string]string `json:"matchLabels,omitempty" protobuf:"bytes,1,rep,name=matchLabels"`
}
