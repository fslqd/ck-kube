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

package fake

import (
	"context"

	v1beta1 "github.com/fslqd/ck-kube/kubernetes/api/authorization/v1beta1"
	v1 "github.com/fslqd/ck-kube/kubernetes/apimachinery/pkg/apis/meta/v1"
	testing "github.com/fslqd/ck-kube/kubernetes/client-go/testing"
)

// FakeSelfSubjectRulesReviews implements SelfSubjectRulesReviewInterface
type FakeSelfSubjectRulesReviews struct {
	Fake *FakeAuthorizationV1beta1
}

var selfsubjectrulesreviewsResource = v1beta1.SchemeGroupVersion.WithResource("selfsubjectrulesreviews")

var selfsubjectrulesreviewsKind = v1beta1.SchemeGroupVersion.WithKind("SelfSubjectRulesReview")

// Create takes the representation of a selfSubjectRulesReview and creates it.  Returns the server's representation of the selfSubjectRulesReview, and an error, if there is any.
func (c *FakeSelfSubjectRulesReviews) Create(ctx context.Context, selfSubjectRulesReview *v1beta1.SelfSubjectRulesReview, opts v1.CreateOptions) (result *v1beta1.SelfSubjectRulesReview, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(selfsubjectrulesreviewsResource, selfSubjectRulesReview), &v1beta1.SelfSubjectRulesReview{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.SelfSubjectRulesReview), err
}
