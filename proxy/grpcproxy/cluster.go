// Copyright 2016 The etcd Authors
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
	"context"
	"errors"
	"time"

	"github.com/ozonru/etcd/v3/clientv3"
	pb "github.com/ozonru/etcd/v3/etcdserver/etcdserverpb"

	"go.uber.org/zap"
)

type clusterProxy struct{}

// NewClusterProxy is deprecated! Do not use it.
//
//
// Takes optional prefix to fetch grpc-proxy member endpoints.
// The returned channel is closed when there is grpc-proxy endpoint registered
// and the client's context is canceled so the 'register' loop returns.
func NewClusterProxy(_ *zap.Logger, _ *clientv3.Client, _ string, _ string) (pb.ClusterServer, <-chan struct{}) {
	cp := &clusterProxy{}

	// this channel always skipped, don't care about it's state
	donec := make(chan struct{})
	go func() {
		// close after enough time
		time.Sleep(time.Millisecond * 100)
		close(donec)
	}()

	return cp, donec
}

func (cp *clusterProxy) MemberAdd(ctx context.Context, r *pb.MemberAddRequest) (*pb.MemberAddResponse, error) {
	return nil, errors.New("not implemented")
}

func (cp *clusterProxy) MemberRemove(ctx context.Context, r *pb.MemberRemoveRequest) (*pb.MemberRemoveResponse, error) {
	return nil, errors.New("not implemented")
}

func (cp *clusterProxy) MemberUpdate(ctx context.Context, r *pb.MemberUpdateRequest) (*pb.MemberUpdateResponse, error) {
	return nil, errors.New("not implemented")
}

// MemberList wraps member list API with following rules:
// - If 'advaddr' is not empty and 'prefix' is not empty, return registered member lists via resolver
// - If 'advaddr' is not empty and 'prefix' is not empty and registered grpc-proxy members haven't been fetched, return the 'advaddr'
// - If 'advaddr' is not empty and 'prefix' is empty, return 'advaddr' without forcing it to 'register'
// - If 'advaddr' is empty, forward to member list API
func (cp *clusterProxy) MemberList(ctx context.Context, r *pb.MemberListRequest) (*pb.MemberListResponse, error) {
	return nil, errors.New("not implemented")
}

func (cp *clusterProxy) MemberPromote(ctx context.Context, r *pb.MemberPromoteRequest) (*pb.MemberPromoteResponse, error) {
	// TODO: implement
	return nil, errors.New("not implemented")
}
