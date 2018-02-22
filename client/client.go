// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "user": Client
//
// Command:
// $ goagen
// --design=github.com/JormungandrK/user-microservice/design
// --out=$(GOPATH)/src/github.com/JormungandrK/user-microservice
// --version=v1.3.1

package client

import (
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
)

// Client is the user service client.
type Client struct {
	*goaclient.Client
	Encoder *goa.HTTPEncoder
	Decoder *goa.HTTPDecoder
}

// New instantiates the client.
func New(c goaclient.Doer) *Client {
	client := &Client{
		Client:  goaclient.New(c),
		Encoder: goa.NewHTTPEncoder(),
		Decoder: goa.NewHTTPDecoder(),
	}

	// Setup encoders and decoders
	client.Encoder.Register(goa.NewJSONEncoder, "application/json")
	client.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	client.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	client.Decoder.Register(goa.NewJSONDecoder, "application/json")
	client.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	client.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	client.Encoder.Register(goa.NewJSONEncoder, "*/*")
	client.Decoder.Register(goa.NewJSONDecoder, "*/*")

	return client
}
