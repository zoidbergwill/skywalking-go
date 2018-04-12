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
	"container/list"
)

type TracingContext struct {
	finishedSpans list.List
}

// Create an entry span for incoming request, for serve side of RPC
func (t *TracingContext) CreateEntrySpan(parentSpan trace.Span, operationName string) trace.Span {
	span := newTracingSpan(t, operationName)
	span.AsEntry()
	return span
}

// Create a local span for local span, no across process related
func (t *TracingContext) CreateLocalSpan(parentSpan trace.Span, operationName string) trace.Span {
	span := newTracingSpan(t, operationName)
	return span
}

// Create an exit span for outgoing request, for client side of RPC
func (t *TracingContext) CreateExitSpan(parentSpan trace.Span, operationName string, remotePeer string) trace.Span {
	span := newTracingSpan(t, operationName)
	span.AsExit()
	return span
}

func (t *TracingContext) Extract(carrier *propagation.ContextCarrier) {
}

func (t *TracingContext) Inject() *propagation.ContextCarrier {
	carrier := propagation.NewContextCarrier()
	return carrier
}

type TracingContextCreator struct {
}

func (*TracingContextCreator) Create() SWContext {
	tracingContext := &TracingContext{}
	return tracingContext
}

//////////////////////////////////////////////////////
// Implementor of Span in Tracing Context
//////////////////////////////////////////////////////
type TracingSpan struct {
	spanData *TracingSpanCoreData
	owner    *TracingContext
}

func (s *TracingSpan) End() {
	s.spanData.endTime = GetMillis()
}

func (s *TracingSpan) AsEntry() {
	s.spanData.isEntry = true
}

func (s *TracingSpan) AsExit() {
	s.spanData.isExit = true;
}

func newTracingSpan(contextOwner *TracingContext, operationName string) *TracingSpan {
	coreData := &TracingSpanCoreData{}
	span := &TracingSpan{
		spanData: coreData,
		owner:    contextOwner,
	}
	span.spanData.operationName = operationName
	return span;
}

type TracingSpanCoreData struct {
	spanId        int
	parentSpanId  int
	operationName string
	startTime     int64
	endTime       int64
	isEntry       bool
	isExit        bool
	layer         int8
}
