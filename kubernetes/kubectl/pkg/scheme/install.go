/*
Copyright 2017 The Kubernetes Authors.

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

package scheme

import (
	admissionv1 "ck-kube/kubernetes/api/admission/v1"
	admissionv1beta1 "ck-kube/kubernetes/api/admission/v1beta1"
	admissionregistrationv1 "ck-kube/kubernetes/api/admissionregistration/v1"
	admissionregistrationv1beta1 "ck-kube/kubernetes/api/admissionregistration/v1beta1"
	appsv1 "ck-kube/kubernetes/api/apps/v1"
	appsv1beta1 "ck-kube/kubernetes/api/apps/v1beta1"
	appsv1beta2 "ck-kube/kubernetes/api/apps/v1beta2"
	authenticationv1 "ck-kube/kubernetes/api/authentication/v1"
	authenticationv1beta1 "ck-kube/kubernetes/api/authentication/v1beta1"
	authorizationv1 "ck-kube/kubernetes/api/authorization/v1"
	authorizationv1beta1 "ck-kube/kubernetes/api/authorization/v1beta1"
	autoscalingv1 "ck-kube/kubernetes/api/autoscaling/v1"
	autoscalingv2 "ck-kube/kubernetes/api/autoscaling/v2"
	batchv1 "ck-kube/kubernetes/api/batch/v1"
	batchv1beta1 "ck-kube/kubernetes/api/batch/v1beta1"
	certificatesv1 "ck-kube/kubernetes/api/certificates/v1"
	certificatesv1beta1 "ck-kube/kubernetes/api/certificates/v1beta1"
	corev1 "ck-kube/kubernetes/api/core/v1"
	extensionsv1beta1 "ck-kube/kubernetes/api/extensions/v1beta1"
	imagepolicyv1alpha1 "ck-kube/kubernetes/api/imagepolicy/v1alpha1"
	networkingv1 "ck-kube/kubernetes/api/networking/v1"
	policyv1 "ck-kube/kubernetes/api/policy/v1"
	policyv1beta1 "ck-kube/kubernetes/api/policy/v1beta1"
	rbacv1 "ck-kube/kubernetes/api/rbac/v1"
	rbacv1alpha1 "ck-kube/kubernetes/api/rbac/v1alpha1"
	rbacv1beta1 "ck-kube/kubernetes/api/rbac/v1beta1"
	schedulingv1alpha1 "ck-kube/kubernetes/api/scheduling/v1alpha1"
	storagev1 "ck-kube/kubernetes/api/storage/v1"
	storagev1alpha1 "ck-kube/kubernetes/api/storage/v1alpha1"
	storagev1beta1 "ck-kube/kubernetes/api/storage/v1beta1"
	metav1 "ck-kube/kubernetes/apimachinery/pkg/apis/meta/v1"
	metav1beta1 "ck-kube/kubernetes/apimachinery/pkg/apis/meta/v1beta1"
	"ck-kube/kubernetes/apimachinery/pkg/runtime/schema"
	utilruntime "ck-kube/kubernetes/apimachinery/pkg/util/runtime"
	"ck-kube/kubernetes/client-go/kubernetes/scheme"
)

// Register all groups in the kubectl's registry, but no componentconfig group since it's not in ck-kube/kubernetes/api
// The code in this file mostly duplicate the install under ck-kube/kubernetes/kubernetes/pkg/api and ck-kube/kubernetes/kubernetes/pkg/apis,
// but does NOT register the internal types.
func init() {
	// Register external types for Scheme
	metav1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(metav1beta1.AddMetaToScheme(Scheme))
	utilruntime.Must(metav1.AddMetaToScheme(Scheme))
	utilruntime.Must(scheme.AddToScheme(Scheme))

	utilruntime.Must(Scheme.SetVersionPriority(corev1.SchemeGroupVersion))
	utilruntime.Must(Scheme.SetVersionPriority(admissionv1beta1.SchemeGroupVersion, admissionv1.SchemeGroupVersion))
	utilruntime.Must(Scheme.SetVersionPriority(admissionregistrationv1beta1.SchemeGroupVersion, admissionregistrationv1.SchemeGroupVersion))
	utilruntime.Must(Scheme.SetVersionPriority(appsv1beta1.SchemeGroupVersion, appsv1beta2.SchemeGroupVersion, appsv1.SchemeGroupVersion))
	utilruntime.Must(Scheme.SetVersionPriority(authenticationv1.SchemeGroupVersion, authenticationv1beta1.SchemeGroupVersion))
	utilruntime.Must(Scheme.SetVersionPriority(authorizationv1.SchemeGroupVersion, authorizationv1beta1.SchemeGroupVersion))
	utilruntime.Must(Scheme.SetVersionPriority(autoscalingv1.SchemeGroupVersion, autoscalingv2.SchemeGroupVersion))
	utilruntime.Must(Scheme.SetVersionPriority(batchv1.SchemeGroupVersion, batchv1beta1.SchemeGroupVersion))
	utilruntime.Must(Scheme.SetVersionPriority(certificatesv1.SchemeGroupVersion, certificatesv1beta1.SchemeGroupVersion))
	utilruntime.Must(Scheme.SetVersionPriority(extensionsv1beta1.SchemeGroupVersion))
	utilruntime.Must(Scheme.SetVersionPriority(imagepolicyv1alpha1.SchemeGroupVersion))
	utilruntime.Must(Scheme.SetVersionPriority(networkingv1.SchemeGroupVersion))
	utilruntime.Must(Scheme.SetVersionPriority(policyv1beta1.SchemeGroupVersion, policyv1.SchemeGroupVersion))
	utilruntime.Must(Scheme.SetVersionPriority(rbacv1.SchemeGroupVersion, rbacv1beta1.SchemeGroupVersion, rbacv1alpha1.SchemeGroupVersion))
	utilruntime.Must(Scheme.SetVersionPriority(schedulingv1alpha1.SchemeGroupVersion))
	utilruntime.Must(Scheme.SetVersionPriority(storagev1.SchemeGroupVersion, storagev1beta1.SchemeGroupVersion, storagev1alpha1.SchemeGroupVersion))
}
