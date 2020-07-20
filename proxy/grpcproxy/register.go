// Copyright 2017 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package grpcproxy

import (
	"github.com/ozonru/etcd/clientv3"

	"go.uber.org/zap"
)

// allow maximum 1 retry per second
const registerRetryRate = 1

// Register is deprecated! Do not use it.
// Now it doesn't return a closer channel.
//
//
// Registers itself as a grpc-proxy server by writing prefixed-key
// with session of specified TTL (in seconds). The returned channel is closed
// when the client's context is canceled.
func Register(lg *zap.Logger, c *clientv3.Client, prefix string, addr string, ttl int) {}
