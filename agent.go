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

import "context"
import "skywalking-go/propagation"

// In most tracing system, you will know this as tracer
// SkyWalking's agent intends to collector Memory, CPU, process id etc.
// so it is not just a simple tracer.
// Initialize agent by using NewAgent method.
type Agent struct {
}

// Initialize agent with given options
func NewAgent(opts ...AgentOptions) (*Agent, error) {
	return nil, nil
}

// Initialize agent with necessary arguments only
// for easier usage.
func NewAgentWithDefaultOptions(applicationCode string, directServerList ...string) (*Agent, error) {
	return nil, nil
}

// Create an entry span for incoming request, for serve side of RPC
func (a *Agent) CreateEntrySpan(ctx context.Context, operationName string, carrier propagation.ContextCarrier) (Span, context.Context) {
	return nil, ctx;
}

// Create a local span for local span, no across process related
func (a *Agent) CreateLocalSpan(ctx context.Context, operationName string) (Span, context.Context) {
	return nil, ctx;
}

// Create an exit span for outgoing request, for client side of RPC
func (a *Agent) CreateExitSpan(ctx context.Context, operationName string) (Span, context.Context, propagation.ContextCarrier) {
	return nil, ctx, nil;
}

// Inject the current status of Context into the ContextCarrier for across thread propagation
// Inject func is a part of CreateExitSpan
func (a *Agent) Inject(ctx context.Context) propagation.ContextCarrier {
	return nil
}

// Extract the ContextCarrier's info into Context for continue the trace from client side
// Extract fun is a part of Create CreateEntrySpan
func (a *Agent) Extract(ctx context.Context, carrier propagation.ContextCarrier) context.Context {
	return nil;
}
