/*
 *
 *  * Licensed to the OpenSkywalking under one or more
 *  * contributor license agreements.  See the NOTICE file distributed with
 *  * this work for additional information regarding copyright ownership.
 *  * The ASF licenses this file to You under the Apache License, Version 2.0
 *  * (the "License"); you may not use this file except in compliance with
 *  * the License.  You may obtain a copy of the License at
 *  *
 *  *     http://www.apache.org/licenses/LICENSE-2.0
 *  *
 *  * Unless required by applicable law or agreed to in writing, software
 *  * distributed under the License is distributed on an "AS IS" BASIS,
 *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  * See the License for the specific language governing permissions and
 *  * limitations under the License.
 *  *
 *
 */

package context

import (
	"github.com/OpenSkywalking/skywalking-go/trace"
	"github.com/OpenSkywalking/skywalking-go/propagation"
)

type TracingContext struct {
}

// Create an entry span for incoming request, for serve side of RPC
func (t *TracingContext) CreateEntrySpan(swContext SWContext, operationName string) trace.Span {
	return nil
}

// Create a local span for local span, no across process related
func (t *TracingContext) CreateLocalSpan(swContext SWContext, operationName string) trace.Span {
	return nil
}

// Create an exit span for outgoing request, for client side of RPC
func (t *TracingContext) CreateExitSpan(swContext SWContext, operationName string, remotePeer string) (trace.Span, *propagation.ContextCarrier) {
	carrier := propagation.NewContextCarrier()
	return nil, carrier
}

func (t *TracingContext) Extract(swContext SWContext, carrier *propagation.ContextCarrier) {
}

func (t *TracingContext) Inject(swContext SWContext) *propagation.ContextCarrier {
	return nil
}

type TracingContextCreator struct {
}

func (*TracingContextCreator) Create() SWContext {
	tracingContext := &TracingContext{}
	return tracingContext
}
