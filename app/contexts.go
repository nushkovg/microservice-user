// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "user": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/JormungandrK/microservice-user/design
// --out=$(GOPATH)/src/github.com/JormungandrK/microservice-user
// --version=v1.3.0

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// CreateUserContext provides the user create action context.
type CreateUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateUserPayload
}

// NewCreateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller create action.
func NewCreateUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreateUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreateUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateUserContext) Created(r *Users) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.user+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 201, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreateUserContext) BadRequest(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CreateUserContext) InternalServerError(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// FindUserContext provides the user find action context.
type FindUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *Credentials
}

// NewFindUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller find action.
func NewFindUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*FindUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := FindUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *FindUserContext) OK(r *Users) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.user+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *FindUserContext) BadRequest(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *FindUserContext) NotFound(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *FindUserContext) InternalServerError(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// FindByEmailUserContext provides the user findByEmail action context.
type FindByEmailUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *EmailPayload
}

// NewFindByEmailUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller findByEmail action.
func NewFindByEmailUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*FindByEmailUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := FindByEmailUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *FindByEmailUserContext) OK(r *Users) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.user+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *FindByEmailUserContext) NotFound(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *FindByEmailUserContext) InternalServerError(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// GetUserContext provides the user get action context.
type GetUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	UserID string
}

// NewGetUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller get action.
func NewGetUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*GetUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := GetUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramUserID := req.Params["userId"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		rctx.UserID = rawUserID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetUserContext) OK(r *Users) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.user+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *GetUserContext) BadRequest(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *GetUserContext) NotFound(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *GetUserContext) InternalServerError(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// GetMeUserContext provides the user getMe action context.
type GetMeUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewGetMeUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller getMe action.
func NewGetMeUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*GetMeUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := GetMeUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetMeUserContext) OK(r *Users) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.user+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *GetMeUserContext) BadRequest(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *GetMeUserContext) NotFound(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *GetMeUserContext) InternalServerError(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// ResetVerificationTokenUserContext provides the user resetVerificationToken action context.
type ResetVerificationTokenUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *EmailPayload
}

// NewResetVerificationTokenUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller resetVerificationToken action.
func NewResetVerificationTokenUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*ResetVerificationTokenUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ResetVerificationTokenUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ResetVerificationTokenUserContext) OK(r *ResetToken) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "resettokenmedia")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ResetVerificationTokenUserContext) BadRequest(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ResetVerificationTokenUserContext) NotFound(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ResetVerificationTokenUserContext) InternalServerError(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// UpdateUserContext provides the user update action context.
type UpdateUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	UserID  string
	Payload *UpdateUserPayload
}

// NewUpdateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller update action.
func NewUpdateUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*UpdateUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := UpdateUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramUserID := req.Params["userId"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		rctx.UserID = rawUserID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *UpdateUserContext) OK(r *Users) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.user+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *UpdateUserContext) BadRequest(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateUserContext) NotFound(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *UpdateUserContext) InternalServerError(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// VerifyUserContext provides the user verify action context.
type VerifyUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Token *string
}

// NewVerifyUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller verify action.
func NewVerifyUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*VerifyUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := VerifyUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramToken := req.Params["token"]
	if len(paramToken) > 0 {
		rawToken := paramToken[0]
		rctx.Token = &rawToken
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *VerifyUserContext) OK(resp []byte) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "plain/text")
	}
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *VerifyUserContext) BadRequest(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *VerifyUserContext) NotFound(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *VerifyUserContext) InternalServerError(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}
