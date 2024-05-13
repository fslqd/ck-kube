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

// Code generated by client-gen. DO NOT EDIT.

package scheme

import (
	admissionregistrationv1 "github.com/fslqd/ck-kube/kubernetes/api/admissionregistration/v1"
	admissionregistrationv1alpha1 "github.com/fslqd/ck-kube/kubernetes/api/admissionregistration/v1alpha1"
	admissionregistrationv1beta1 "github.com/fslqd/ck-kube/kubernetes/api/admissionregistration/v1beta1"
	internalv1alpha1 "github.com/fslqd/ck-kube/kubernetes/api/apiserverinternal/v1alpha1"
	appsv1 "github.com/fslqd/ck-kube/kubernetes/api/apps/v1"
	appsv1beta1 "github.com/fslqd/ck-kube/kubernetes/api/apps/v1beta1"
	appsv1beta2 "github.com/fslqd/ck-kube/kubernetes/api/apps/v1beta2"
	authenticationv1 "github.com/fslqd/ck-kube/kubernetes/api/authentication/v1"
	authenticationv1alpha1 "github.com/fslqd/ck-kube/kubernetes/api/authentication/v1alpha1"
	authenticationv1beta1 "github.com/fslqd/ck-kube/kubernetes/api/authentication/v1beta1"
	authorizationv1 "github.com/fslqd/ck-kube/kubernetes/api/authorization/v1"
	authorizationv1beta1 "github.com/fslqd/ck-kube/kubernetes/api/authorization/v1beta1"
	autoscalingv1 "github.com/fslqd/ck-kube/kubernetes/api/autoscaling/v1"
	autoscalingv2 "github.com/fslqd/ck-kube/kubernetes/api/autoscaling/v2"
	autoscalingv2beta1 "github.com/fslqd/ck-kube/kubernetes/api/autoscaling/v2beta1"
	autoscalingv2beta2 "github.com/fslqd/ck-kube/kubernetes/api/autoscaling/v2beta2"
	batchv1 "github.com/fslqd/ck-kube/kubernetes/api/batch/v1"
	batchv1beta1 "github.com/fslqd/ck-kube/kubernetes/api/batch/v1beta1"
	certificatesv1 "github.com/fslqd/ck-kube/kubernetes/api/certificates/v1"
	certificatesv1alpha1 "github.com/fslqd/ck-kube/kubernetes/api/certificates/v1alpha1"
	certificatesv1beta1 "github.com/fslqd/ck-kube/kubernetes/api/certificates/v1beta1"
	coordinationv1 "github.com/fslqd/ck-kube/kubernetes/api/coordination/v1"
	coordinationv1beta1 "github.com/fslqd/ck-kube/kubernetes/api/coordination/v1beta1"
	corev1 "github.com/fslqd/ck-kube/kubernetes/api/core/v1"
	discoveryv1 "github.com/fslqd/ck-kube/kubernetes/api/discovery/v1"
	discoveryv1beta1 "github.com/fslqd/ck-kube/kubernetes/api/discovery/v1beta1"
	eventsv1 "github.com/fslqd/ck-kube/kubernetes/api/events/v1"
	eventsv1beta1 "github.com/fslqd/ck-kube/kubernetes/api/events/v1beta1"
	extensionsv1beta1 "github.com/fslqd/ck-kube/kubernetes/api/extensions/v1beta1"
	flowcontrolv1 "github.com/fslqd/ck-kube/kubernetes/api/flowcontrol/v1"
	flowcontrolv1beta1 "github.com/fslqd/ck-kube/kubernetes/api/flowcontrol/v1beta1"
	flowcontrolv1beta2 "github.com/fslqd/ck-kube/kubernetes/api/flowcontrol/v1beta2"
	flowcontrolv1beta3 "github.com/fslqd/ck-kube/kubernetes/api/flowcontrol/v1beta3"
	networkingv1 "github.com/fslqd/ck-kube/kubernetes/api/networking/v1"
	networkingv1alpha1 "github.com/fslqd/ck-kube/kubernetes/api/networking/v1alpha1"
	networkingv1beta1 "github.com/fslqd/ck-kube/kubernetes/api/networking/v1beta1"
	nodev1 "github.com/fslqd/ck-kube/kubernetes/api/node/v1"
	nodev1alpha1 "github.com/fslqd/ck-kube/kubernetes/api/node/v1alpha1"
	nodev1beta1 "github.com/fslqd/ck-kube/kubernetes/api/node/v1beta1"
	policyv1 "github.com/fslqd/ck-kube/kubernetes/api/policy/v1"
	policyv1beta1 "github.com/fslqd/ck-kube/kubernetes/api/policy/v1beta1"
	rbacv1 "github.com/fslqd/ck-kube/kubernetes/api/rbac/v1"
	rbacv1alpha1 "github.com/fslqd/ck-kube/kubernetes/api/rbac/v1alpha1"
	rbacv1beta1 "github.com/fslqd/ck-kube/kubernetes/api/rbac/v1beta1"
	resourcev1alpha2 "github.com/fslqd/ck-kube/kubernetes/api/resource/v1alpha2"
	schedulingv1 "github.com/fslqd/ck-kube/kubernetes/api/scheduling/v1"
	schedulingv1alpha1 "github.com/fslqd/ck-kube/kubernetes/api/scheduling/v1alpha1"
	schedulingv1beta1 "github.com/fslqd/ck-kube/kubernetes/api/scheduling/v1beta1"
	storagev1 "github.com/fslqd/ck-kube/kubernetes/api/storage/v1"
	storagev1alpha1 "github.com/fslqd/ck-kube/kubernetes/api/storage/v1alpha1"
	storagev1beta1 "github.com/fslqd/ck-kube/kubernetes/api/storage/v1beta1"
	storagemigrationv1alpha1 "github.com/fslqd/ck-kube/kubernetes/api/storagemigration/v1alpha1"
	v1 "github.com/fslqd/ck-kube/kubernetes/apimachinery/pkg/apis/meta/v1"
	runtime "github.com/fslqd/ck-kube/kubernetes/apimachinery/pkg/runtime"
	schema "github.com/fslqd/ck-kube/kubernetes/apimachinery/pkg/runtime/schema"
	serializer "github.com/fslqd/ck-kube/kubernetes/apimachinery/pkg/runtime/serializer"
	utilruntime "github.com/fslqd/ck-kube/kubernetes/apimachinery/pkg/util/runtime"
)

var Scheme = runtime.NewScheme()
var Codecs = serializer.NewCodecFactory(Scheme)
var ParameterCodec = runtime.NewParameterCodec(Scheme)
var localSchemeBuilder = runtime.SchemeBuilder{
	admissionregistrationv1.AddToScheme,
	admissionregistrationv1alpha1.AddToScheme,
	admissionregistrationv1beta1.AddToScheme,
	internalv1alpha1.AddToScheme,
	appsv1.AddToScheme,
	appsv1beta1.AddToScheme,
	appsv1beta2.AddToScheme,
	authenticationv1.AddToScheme,
	authenticationv1alpha1.AddToScheme,
	authenticationv1beta1.AddToScheme,
	authorizationv1.AddToScheme,
	authorizationv1beta1.AddToScheme,
	autoscalingv1.AddToScheme,
	autoscalingv2.AddToScheme,
	autoscalingv2beta1.AddToScheme,
	autoscalingv2beta2.AddToScheme,
	batchv1.AddToScheme,
	batchv1beta1.AddToScheme,
	certificatesv1.AddToScheme,
	certificatesv1beta1.AddToScheme,
	certificatesv1alpha1.AddToScheme,
	coordinationv1beta1.AddToScheme,
	coordinationv1.AddToScheme,
	corev1.AddToScheme,
	discoveryv1.AddToScheme,
	discoveryv1beta1.AddToScheme,
	eventsv1.AddToScheme,
	eventsv1beta1.AddToScheme,
	extensionsv1beta1.AddToScheme,
	flowcontrolv1.AddToScheme,
	flowcontrolv1beta1.AddToScheme,
	flowcontrolv1beta2.AddToScheme,
	flowcontrolv1beta3.AddToScheme,
	networkingv1.AddToScheme,
	networkingv1alpha1.AddToScheme,
	networkingv1beta1.AddToScheme,
	nodev1.AddToScheme,
	nodev1alpha1.AddToScheme,
	nodev1beta1.AddToScheme,
	policyv1.AddToScheme,
	policyv1beta1.AddToScheme,
	rbacv1.AddToScheme,
	rbacv1beta1.AddToScheme,
	rbacv1alpha1.AddToScheme,
	resourcev1alpha2.AddToScheme,
	schedulingv1alpha1.AddToScheme,
	schedulingv1beta1.AddToScheme,
	schedulingv1.AddToScheme,
	storagev1beta1.AddToScheme,
	storagev1.AddToScheme,
	storagev1alpha1.AddToScheme,
	storagemigrationv1alpha1.AddToScheme,
}

// AddToScheme adds all types of this clientset into the given scheme. This allows composition
// of clientsets, like in:
//
//	import (
//	  "github.com/fslqd/ck-kube/kubernetes/client-go/kubernetes"
//	  clientsetscheme "github.com/fslqd/ck-kube/kubernetes/client-go/kubernetes/scheme"
//	  aggregatorclientsetscheme "github.com/fslqd/ck-kube/kubernetes/kube-aggregator/pkg/client/clientset_generated/clientset/scheme"
//	)
//
//	kclientset, _ := kubernetes.NewForConfig(c)
//	_ = aggregatorclientsetscheme.AddToScheme(clientsetscheme.Scheme)
//
// After this, RawExtensions in Kubernetes types will serialize kube-aggregator types
// correctly.
var AddToScheme = localSchemeBuilder.AddToScheme

func init() {
	v1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(AddToScheme(Scheme))
}
