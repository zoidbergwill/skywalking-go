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
	"strconv"
	"errors"
	"skywalking-go/trace"
)

// AgentOptions used for initialize agent
type AgentOptions func(a *Agent) error

func WithDirectGRPCIpPort(hostname string, port int) AgentOptions {
	return func(a *Agent) error {
		if hostname == "" {
			return errors.New("hostname must not empty.")
		} else if port == 0 || port > 65535 {
			return errors.New("Network port should be between 0(exclude) and 65535(include).")
		}
		return WithDirectGRPCAddress(hostname + ":" + strconv.Itoa(port))(a)
	}
}

func WithDirectGRPCAddress(address string) AgentOptions {
	return func(a *Agent) error {
		if address == "" {
			return errors.New("address must not empty.")
		}
		a.directServerList = append(a.directServerList, address)
		return nil
	}
}

func WithApplicationCode(applicationCode string) AgentOptions {
	return func(a *Agent) error {
		if applicationCode == "" {
			return errors.New("applicationCode must not empty.")
		}
		a.applicationCode = applicationCode
		return nil
	}
}

func WithChannelSize(bufferSize int) AgentOptions {
	return func(a *Agent) error {
		if (bufferSize < 1) {
			return errors.New("BufferSize must be positive int.")
		}
		a.queue = make(chan trace.TraceSegment, bufferSize)
		return nil
	}
}
