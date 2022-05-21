// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package debugger

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	clouddebuggerpb "google.golang.org/genproto/googleapis/devtools/clouddebugger/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newController2ClientHook clientHook

// Controller2CallOptions contains the retry settings for each method of Controller2Client.
type Controller2CallOptions struct {
	RegisterDebuggee       []gax.CallOption
	ListActiveBreakpoints  []gax.CallOption
	UpdateActiveBreakpoint []gax.CallOption
}

func defaultController2GRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("clouddebugger.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("clouddebugger.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://clouddebugger.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultController2CallOptions() *Controller2CallOptions {
	return &Controller2CallOptions{
		RegisterDebuggee: []gax.CallOption{},
		ListActiveBreakpoints: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateActiveBreakpoint: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalController2Client is an interface that defines the methods availaible from Stackdriver Debugger API.
type internalController2Client interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	RegisterDebuggee(context.Context, *clouddebuggerpb.RegisterDebuggeeRequest, ...gax.CallOption) (*clouddebuggerpb.RegisterDebuggeeResponse, error)
	ListActiveBreakpoints(context.Context, *clouddebuggerpb.ListActiveBreakpointsRequest, ...gax.CallOption) (*clouddebuggerpb.ListActiveBreakpointsResponse, error)
	UpdateActiveBreakpoint(context.Context, *clouddebuggerpb.UpdateActiveBreakpointRequest, ...gax.CallOption) (*clouddebuggerpb.UpdateActiveBreakpointResponse, error)
}

// Controller2Client is a client for interacting with Stackdriver Debugger API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The Controller service provides the API for orchestrating a collection of
// debugger agents to perform debugging tasks. These agents are each attached
// to a process of an application which may include one or more replicas.
//
// The debugger agents register with the Controller to identify the application
// being debugged, the Debuggee. All agents that register with the same data,
// represent the same Debuggee, and are assigned the same debuggee_id.
//
// The debugger agents call the Controller to retrieve  the list of active
// Breakpoints. Agents with the same debuggee_id get the same breakpoints
// list. An agent that can fulfill the breakpoint request updates the
// Controller with the breakpoint result. The controller selects the first
// result received and discards the rest of the results.
// Agents that poll again for active breakpoints will no longer have
// the completed breakpoint in the list and should remove that breakpoint from
// their attached process.
//
// The Controller service does not provide a way to retrieve the results of
// a completed breakpoint. This functionality is available using the Debugger
// service.
type Controller2Client struct {
	// The internal transport-dependent client.
	internalClient internalController2Client

	// The call options for this service.
	CallOptions *Controller2CallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Controller2Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Controller2Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *Controller2Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// RegisterDebuggee registers the debuggee with the controller service.
//
// All agents attached to the same application must call this method with
// exactly the same request content to get back the same stable debuggee_id.
// Agents should call this method again whenever google.rpc.Code.NOT_FOUND
// is returned from any controller method.
//
// This protocol allows the controller service to disable debuggees, recover
// from data loss, or change the debuggee_id format. Agents must handle
// debuggee_id value changing upon re-registration.
func (c *Controller2Client) RegisterDebuggee(ctx context.Context, req *clouddebuggerpb.RegisterDebuggeeRequest, opts ...gax.CallOption) (*clouddebuggerpb.RegisterDebuggeeResponse, error) {
	return c.internalClient.RegisterDebuggee(ctx, req, opts...)
}

// ListActiveBreakpoints returns the list of all active breakpoints for the debuggee.
//
// The breakpoint specification (location, condition, and expressions
// fields) is semantically immutable, although the field values may
// change. For example, an agent may update the location line number
// to reflect the actual line where the breakpoint was set, but this
// doesn’t change the breakpoint semantics.
//
// This means that an agent does not need to check if a breakpoint has changed
// when it encounters the same breakpoint on a successive call.
// Moreover, an agent should remember the breakpoints that are completed
// until the controller removes them from the active list to avoid
// setting those breakpoints again.
func (c *Controller2Client) ListActiveBreakpoints(ctx context.Context, req *clouddebuggerpb.ListActiveBreakpointsRequest, opts ...gax.CallOption) (*clouddebuggerpb.ListActiveBreakpointsResponse, error) {
	return c.internalClient.ListActiveBreakpoints(ctx, req, opts...)
}

// UpdateActiveBreakpoint updates the breakpoint state or mutable fields.
// The entire Breakpoint message must be sent back to the controller service.
//
// Updates to active breakpoint fields are only allowed if the new value
// does not change the breakpoint specification. Updates to the location,
// condition and expressions fields should not alter the breakpoint
// semantics. These may only make changes such as canonicalizing a value
// or snapping the location to the correct line of code.
func (c *Controller2Client) UpdateActiveBreakpoint(ctx context.Context, req *clouddebuggerpb.UpdateActiveBreakpointRequest, opts ...gax.CallOption) (*clouddebuggerpb.UpdateActiveBreakpointResponse, error) {
	return c.internalClient.UpdateActiveBreakpoint(ctx, req, opts...)
}

// controller2GRPCClient is a client for interacting with Stackdriver Debugger API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type controller2GRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing Controller2Client
	CallOptions **Controller2CallOptions

	// The gRPC API client.
	controller2Client clouddebuggerpb.Controller2Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewController2Client creates a new controller2 client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The Controller service provides the API for orchestrating a collection of
// debugger agents to perform debugging tasks. These agents are each attached
// to a process of an application which may include one or more replicas.
//
// The debugger agents register with the Controller to identify the application
// being debugged, the Debuggee. All agents that register with the same data,
// represent the same Debuggee, and are assigned the same debuggee_id.
//
// The debugger agents call the Controller to retrieve  the list of active
// Breakpoints. Agents with the same debuggee_id get the same breakpoints
// list. An agent that can fulfill the breakpoint request updates the
// Controller with the breakpoint result. The controller selects the first
// result received and discards the rest of the results.
// Agents that poll again for active breakpoints will no longer have
// the completed breakpoint in the list and should remove that breakpoint from
// their attached process.
//
// The Controller service does not provide a way to retrieve the results of
// a completed breakpoint. This functionality is available using the Debugger
// service.
func NewController2Client(ctx context.Context, opts ...option.ClientOption) (*Controller2Client, error) {
	clientOpts := defaultController2GRPCClientOptions()
	if newController2ClientHook != nil {
		hookOpts, err := newController2ClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := Controller2Client{CallOptions: defaultController2CallOptions()}

	c := &controller2GRPCClient{
		connPool:          connPool,
		disableDeadlines:  disableDeadlines,
		controller2Client: clouddebuggerpb.NewController2Client(connPool),
		CallOptions:       &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *controller2GRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *controller2GRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *controller2GRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *controller2GRPCClient) RegisterDebuggee(ctx context.Context, req *clouddebuggerpb.RegisterDebuggeeRequest, opts ...gax.CallOption) (*clouddebuggerpb.RegisterDebuggeeResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).RegisterDebuggee[0:len((*c.CallOptions).RegisterDebuggee):len((*c.CallOptions).RegisterDebuggee)], opts...)
	var resp *clouddebuggerpb.RegisterDebuggeeResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.controller2Client.RegisterDebuggee(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *controller2GRPCClient) ListActiveBreakpoints(ctx context.Context, req *clouddebuggerpb.ListActiveBreakpointsRequest, opts ...gax.CallOption) (*clouddebuggerpb.ListActiveBreakpointsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "debuggee_id", url.QueryEscape(req.GetDebuggeeId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListActiveBreakpoints[0:len((*c.CallOptions).ListActiveBreakpoints):len((*c.CallOptions).ListActiveBreakpoints)], opts...)
	var resp *clouddebuggerpb.ListActiveBreakpointsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.controller2Client.ListActiveBreakpoints(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *controller2GRPCClient) UpdateActiveBreakpoint(ctx context.Context, req *clouddebuggerpb.UpdateActiveBreakpointRequest, opts ...gax.CallOption) (*clouddebuggerpb.UpdateActiveBreakpointResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "debuggee_id", url.QueryEscape(req.GetDebuggeeId()), "breakpoint.id", url.QueryEscape(req.GetBreakpoint().GetId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateActiveBreakpoint[0:len((*c.CallOptions).UpdateActiveBreakpoint):len((*c.CallOptions).UpdateActiveBreakpoint)], opts...)
	var resp *clouddebuggerpb.UpdateActiveBreakpointResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.controller2Client.UpdateActiveBreakpoint(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
