// Code generated by protoc-gen-orion. DO NOT EDIT.
// source: foo/foo.proto

package foo

import (
	orion "github.com/carousell/Orion/orion"
)

// If you see error please update your orion-protoc-gen by running 'go get -u github.com/carousell/Orion/protoc-gen-orion'
var _ = orion.ProtoGenVersion1_0

// Encoders

// Handlers

// Decoders

//Streams

// RegisterFooServiceOrionServer registers FooService to Orion server
// Services need to pass either ServiceFactory or ServiceFactoryV2 implementation
func RegisterFooServiceOrionServer(sf interface{}, orionServer orion.Server) error {
	err := orionServer.RegisterService(&_FooService_serviceDesc, sf)
	if err != nil {
		return err
	}

	return nil
}

// DefaultEncoder
func RegisterFooServiceDefaultEncoder(svr orion.Server, encoder orion.Encoder) {
	orion.RegisterDefaultEncoder(svr, "FooService", encoder)
}

// DefaultDecoder
func RegisterFooServiceDefaultDecoder(svr orion.Server, decoder orion.Decoder) {
	orion.RegisterDefaultDecoder(svr, "FooService", decoder)
}
