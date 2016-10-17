//************************************************************************//
// API "adder": Application Contexts
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/kkeuning/cug/examples/adder2/design
// --out=$(GOPATH)/src/github.com/kkeuning/cug/examples/adder2
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"strconv"
)

// AddOperandsContext provides the operands add action context.
type AddOperandsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Left  int
	Right int
}

// NewAddOperandsContext parses the incoming request URL and body, performs validations and creates the
// context used by the operands controller add action.
func NewAddOperandsContext(ctx context.Context, service *goa.Service) (*AddOperandsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := AddOperandsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramLeft := req.Params["left"]
	if len(paramLeft) > 0 {
		rawLeft := paramLeft[0]
		if left, err2 := strconv.Atoi(rawLeft); err2 == nil {
			rctx.Left = left
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("left", rawLeft, "integer"))
		}
	}
	paramRight := req.Params["right"]
	if len(paramRight) > 0 {
		rawRight := paramRight[0]
		if right, err2 := strconv.Atoi(rawRight); err2 == nil {
			rctx.Right = right
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("right", rawRight, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *AddOperandsContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "plain/text")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// MultiplyOperandsContext provides the operands multiply action context.
type MultiplyOperandsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Left  int
	Right int
}

// NewMultiplyOperandsContext parses the incoming request URL and body, performs validations and creates the
// context used by the operands controller multiply action.
func NewMultiplyOperandsContext(ctx context.Context, service *goa.Service) (*MultiplyOperandsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := MultiplyOperandsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramLeft := req.Params["left"]
	if len(paramLeft) > 0 {
		rawLeft := paramLeft[0]
		if left, err2 := strconv.Atoi(rawLeft); err2 == nil {
			rctx.Left = left
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("left", rawLeft, "integer"))
		}
	}
	paramRight := req.Params["right"]
	if len(paramRight) > 0 {
		rawRight := paramRight[0]
		if right, err2 := strconv.Atoi(rawRight); err2 == nil {
			rctx.Right = right
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("right", rawRight, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *MultiplyOperandsContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "plain/text")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// SubtractOperandsContext provides the operands subtract action context.
type SubtractOperandsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Left  int
	Right int
}

// NewSubtractOperandsContext parses the incoming request URL and body, performs validations and creates the
// context used by the operands controller subtract action.
func NewSubtractOperandsContext(ctx context.Context, service *goa.Service) (*SubtractOperandsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := SubtractOperandsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramLeft := req.Params["left"]
	if len(paramLeft) > 0 {
		rawLeft := paramLeft[0]
		if left, err2 := strconv.Atoi(rawLeft); err2 == nil {
			rctx.Left = left
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("left", rawLeft, "integer"))
		}
	}
	paramRight := req.Params["right"]
	if len(paramRight) > 0 {
		rawRight := paramRight[0]
		if right, err2 := strconv.Atoi(rawRight); err2 == nil {
			rctx.Right = right
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("right", rawRight, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *SubtractOperandsContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "plain/text")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}
