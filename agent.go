/*
 * Licensed to the OpenSkywalking under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package skywalking

import (
	"context"

	"github.com/OpenSkywalking/skywalking-go/propagation"
	"github.com/OpenSkywalking/skywalking-go/trace"
	"github.com/OpenSkywalking/skywalking-go/reporter"
	traceContext "github.com/OpenSkywalking/skywalking-go/context"
)

// In most tracing system, you will know this as tracer
// SkyWalking's agent intends to collector Memory, CPU, process id etc.
// so it is not just a simple tracer.
// Initialize agent by using NewAgent method.
type Agent struct {
	applicationCode string
	contextCreator  traceContext.ContextCreator
	queue           chan trace.TraceSegment
	reporter        reporter.SegmentListener
}

// Initialize agent with given options
func NewAgent(opts ...AgentOptions) (*Agent, error) {
	agent := &Agent{
		applicationCode: "",
	}

	for _, opt := range opts {
		if err := opt(agent); err != nil {
			return nil, err
		}
	}

	return agent, nil
}

// Initialize agent with necessary arguments only
// for easier usage.
func NewAgentWithDefaultOptions(applicationCode string, directServerList ...string) (*Agent, error) {
	return NewAgent(
		WithApplicationCode(applicationCode),
		WithGRPCReporter(directServerList...),
		WithTracingContext(),
	)
}

// Create an entry span for incoming request, for serve side of RPC
func (a *Agent) CreateEntrySpan(ctx context.Context, operationName string, carrier *propagation.ContextCarrier) (context.Context, trace.Span) {
	ctx, swContext := traceContext.GetOrCreateContext(ctx, a.contextCreator)
	span := swContext.CreateEntrySpan(swContext, operationName)
	swContext.Extract(swContext, carrier)
	return ctx, span
}

// Create a local span for local span, no across process related
func (a *Agent) CreateLocalSpan(ctx context.Context, operationName string) (context.Context, trace.Span) {
	ctx, swContext := traceContext.GetOrCreateContext(ctx, a.contextCreator)
	span := swContext.CreateLocalSpan(swContext, operationName)
	return ctx, span
}

// Create an exit span for outgoing request, for client side of RPC
func (a *Agent) CreateExitSpan(ctx context.Context, operationName string) (context.Context, trace.Span, *propagation.ContextCarrier) {
	ctx, swContext := traceContext.GetOrCreateContext(ctx, a.contextCreator)
	span, carrier := swContext.CreateExitSpan(swContext, operationName)
	return ctx, span, carrier
}

// Inject the current status of SWContext into the ContextCarrier for across thread propagation
// Inject func is a part of CreateExitSpan
func (a *Agent) Inject(ctx context.Context) (context.Context, *propagation.ContextCarrier) {
	ctx, swContext := traceContext.GetOrCreateContext(ctx, a.contextCreator)
	return ctx, swContext.Inject(swContext)
}

// Extract the ContextCarrier's info into SWContext for continue the trace from client side
// Extract fun is a part of Create CreateEntrySpan
func (a *Agent) Extract(ctx context.Context, carrier *propagation.ContextCarrier) context.Context {
	ctx, swContext := traceContext.GetOrCreateContext(ctx, a.contextCreator)
	swContext.Extract(swContext, carrier)
	return ctx
}
