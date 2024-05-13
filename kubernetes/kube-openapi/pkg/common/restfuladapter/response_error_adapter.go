package restfuladapter

import (
	"ck-kube/kubernetes/kube-openapi/pkg/common"
	"github.com/emicklei/go-restful/v3"
)

var _ common.StatusCodeResponse = &ResponseErrorAdapter{}

// ResponseErrorAdapter adapts a restful.ResponseError to common.StatusCodeResponse.
type ResponseErrorAdapter struct {
	Err *restful.ResponseError
}

func (r *ResponseErrorAdapter) Message() string {
	return r.Err.Message
}

func (r *ResponseErrorAdapter) Model() interface{} {
	return r.Err.Model
}

func (r *ResponseErrorAdapter) Code() int {
	return r.Err.Code
}
