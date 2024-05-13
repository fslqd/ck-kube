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
	v1beta1 "ck-kube/kubernetes/client-go/kubernetes/typed/flowcontrol/v1beta1"
	rest "ck-kube/kubernetes/client-go/rest"
	testing "ck-kube/kubernetes/client-go/testing"
)

type FakeFlowcontrolV1beta1 struct {
	*testing.Fake
}

func (c *FakeFlowcontrolV1beta1) FlowSchemas() v1beta1.FlowSchemaInterface {
	return &FakeFlowSchemas{c}
}

func (c *FakeFlowcontrolV1beta1) PriorityLevelConfigurations() v1beta1.PriorityLevelConfigurationInterface {
	return &FakePriorityLevelConfigurations{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeFlowcontrolV1beta1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}