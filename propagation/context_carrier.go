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

package propagation

import (
	"container/list"
)

type CarrierItem interface {
	headKey() string
	headValue() string
	setValue(t string)
	isValid() bool
}

// SW3CarrierItem is the only implementation of CarrierItem for now.
// In roadmap, W3C trace context implementation may be added and actived by somehow.
type SW3CarrierItem struct {
}

func (s *SW3CarrierItem) headKey() string {
	return "sw3"
}

func (s *SW3CarrierItem) headValue() string {
	return "";
}

func (s *SW3CarrierItem) setValue(t string) string {
	return "";
}

func (s *SW3CarrierItem) isValid() bool {
	return true;
}

func NewSW3CarrierItem() *SW3CarrierItem {
	item := new(SW3CarrierItem)

	return item
}

// ContextCarrier is a data carrier of tracing context,
// it holds a snapshot for across process propagation.
type ContextCarrier struct {
	items *list.List
}

func (c *ContextCarrier) GetAllItems() *list.List {
	return c.items
}

func NewContextCarrier() (*ContextCarrier) {
	carrier := ContextCarrier{items: list.New()};
	carrier.items.PushBack(NewSW3CarrierItem())
	return &carrier
}
