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
	"context"
	"github.com/OpenSkywalking/skywalking-go/propagation"
	"github.com/OpenSkywalking/skywalking-go/trace"
	"errors"
)

type ctxKey struct{}
type parentSpan struct{}

var ContextKeyHolder = ctxKey{}
var ParentSpanKey = parentSpan{}

type SWContext interface {
	CreateEntrySpan(ctx context.Context, operationName string, carrier *propagation.ContextCarrier) (trace.Span, context.Context)
	CreateLocalSpan(ctx context.Context, operationName string) (trace.Span, context.Context)
	CreateExitSpan(ctx context.Context, operationName string) (trace.Span, context.Context, *propagation.ContextCarrier)
}

// Create or get the existed SkyWalking context from go context.
func GetOrCreateContext(ctx context.Context, creator ContextCreator) (context.Context, SWContext) {
	if swCtx, ok := ctx.Value(ContextKeyHolder).(SWContext); ok {
		return ctx, swCtx
	} else {
		newContext, _ := createNewContext(ctx, creator)
		return newContext, swCtx;
	}
}

func prepareNextContext(ctx context.Context, parentSpan trace.Span) (context.Context, error) {
	if _, ok := ctx.Value(ctxKey{}).(SWContext); ok {
		return context.WithValue(ctx, ParentSpanKey, parentSpan), nil
	} else {
		return ctx, errors.New("prepareNextContext can only be called inside an existed context")
	}
}

func getParenSpan(ctx context.Context) trace.Span {
	if span, ok := ctx.Value(ParentSpanKey).(trace.Span); ok {
		return span
	} else {
		return nil
	}
}

func createNewContext(ctx context.Context, creator ContextCreator) (context.Context, SWContext) {
	swCtx := creator.Create()
	context.WithValue(ctx, ContextKeyHolder, swCtx)
	return ctx, swCtx;
}

type ContextCreator interface {
	Create() SWContext
}
