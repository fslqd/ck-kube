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

package v1beta1

import (
	"context"

	certificates "github.com/fslqd/ck-kube/kubernetes/api/certificates/v1beta1"
	metav1 "github.com/fslqd/ck-kube/kubernetes/apimachinery/pkg/apis/meta/v1"
	scheme "github.com/fslqd/ck-kube/kubernetes/client-go/kubernetes/scheme"
)

type CertificateSigningRequestExpansion interface {
	UpdateApproval(ctx context.Context, certificateSigningRequest *certificates.CertificateSigningRequest, opts metav1.UpdateOptions) (result *certificates.CertificateSigningRequest, err error)
}

func (c *certificateSigningRequests) UpdateApproval(ctx context.Context, certificateSigningRequest *certificates.CertificateSigningRequest, opts metav1.UpdateOptions) (result *certificates.CertificateSigningRequest, err error) {
	result = &certificates.CertificateSigningRequest{}
	err = c.client.Put().
		Resource("certificatesigningrequests").
		Name(certificateSigningRequest.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(certificateSigningRequest).
		SubResource("approval").
		Do(ctx).
		Into(result)
	return
}
