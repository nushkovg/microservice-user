// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "user": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/JormungandrK/user-microservice/design
// --out=$(GOPATH)/src/github.com/JormungandrK/user-microservice
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// SwaggerController is the controller interface for the Swagger actions.
type SwaggerController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl SwaggerController) {
	initService(service)
	var h goa.Handler

	h = ctrl.FileHandler("/swagger-ui/*filepath", "swagger-ui/dist")
	service.Mux.Handle("GET", "/swagger-ui/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger-ui/dist", "route", "GET /swagger-ui/*filepath")

	h = ctrl.FileHandler("/swagger.json", "swagger/swagger.json")
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger/swagger.json", "route", "GET /swagger.json")

	h = ctrl.FileHandler("/swagger-ui/", "swagger-ui/dist/index.html")
	service.Mux.Handle("GET", "/swagger-ui/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger-ui/dist/index.html", "route", "GET /swagger-ui/")
}

// UserController is the controller interface for the User actions.
type UserController interface {
	goa.Muxer
	Create(*CreateUserContext) error
	Find(*FindUserContext) error
	Get(*GetUserContext) error
	GetMe(*GetMeUserContext) error
	Update(*UpdateUserContext) error
}

// MountUserController "mounts" a User resource controller on the given service.
func MountUserController(service *goa.Service, ctrl UserController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateUserContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UserPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/users/", ctrl.MuxHandler("create", h, unmarshalCreateUserPayload))
	service.LogInfo("mount", "ctrl", "User", "action", "Create", "route", "POST /users/")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewFindUserContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Find(rctx)
	}
	service.Mux.Handle("POST", "/users/find", ctrl.MuxHandler("find", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "Find", "route", "POST /users/find")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetUserContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Get(rctx)
	}
	service.Mux.Handle("GET", "/users/:userId", ctrl.MuxHandler("get", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "Get", "route", "GET /users/:userId")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetMeUserContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.GetMe(rctx)
	}
	service.Mux.Handle("GET", "/users/me", ctrl.MuxHandler("getMe", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "GetMe", "route", "GET /users/me")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateUserContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UserPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	service.Mux.Handle("PUT", "/users/:userId", ctrl.MuxHandler("update", h, unmarshalUpdateUserPayload))
	service.LogInfo("mount", "ctrl", "User", "action", "Update", "route", "PUT /users/:userId")
}

// unmarshalCreateUserPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateUserPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &userPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	payload.Finalize()
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateUserPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateUserPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &userPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	payload.Finalize()
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}
