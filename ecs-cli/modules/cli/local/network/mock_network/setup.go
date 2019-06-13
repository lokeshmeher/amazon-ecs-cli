// Copyright 2015-2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-cli/ecs-cli/modules/cli/local/network (interfaces: LocalEndpointsStarter)

// Package mock_network is a generated GoMock package.
package mock_network

import (
	reflect "reflect"

	types "github.com/docker/docker/api/types"
	container "github.com/docker/docker/api/types/container"
	network "github.com/docker/docker/api/types/network"
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
)

// MockLocalEndpointsStarter is a mock of LocalEndpointsStarter interface
type MockLocalEndpointsStarter struct {
	ctrl     *gomock.Controller
	recorder *MockLocalEndpointsStarterMockRecorder
}

// MockLocalEndpointsStarterMockRecorder is the mock recorder for MockLocalEndpointsStarter
type MockLocalEndpointsStarterMockRecorder struct {
	mock *MockLocalEndpointsStarter
}

// NewMockLocalEndpointsStarter creates a new mock instance
func NewMockLocalEndpointsStarter(ctrl *gomock.Controller) *MockLocalEndpointsStarter {
	mock := &MockLocalEndpointsStarter{ctrl: ctrl}
	mock.recorder = &MockLocalEndpointsStarterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLocalEndpointsStarter) EXPECT() *MockLocalEndpointsStarterMockRecorder {
	return m.recorder
}

// ContainerCreate mocks base method
func (m *MockLocalEndpointsStarter) ContainerCreate(arg0 context.Context, arg1 *container.Config, arg2 *container.HostConfig, arg3 *network.NetworkingConfig, arg4 string) (container.ContainerCreateCreatedBody, error) {
	ret := m.ctrl.Call(m, "ContainerCreate", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(container.ContainerCreateCreatedBody)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerCreate indicates an expected call of ContainerCreate
func (mr *MockLocalEndpointsStarterMockRecorder) ContainerCreate(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerCreate", reflect.TypeOf((*MockLocalEndpointsStarter)(nil).ContainerCreate), arg0, arg1, arg2, arg3, arg4)
}

// ContainerList mocks base method
func (m *MockLocalEndpointsStarter) ContainerList(arg0 context.Context, arg1 types.ContainerListOptions) ([]types.Container, error) {
	ret := m.ctrl.Call(m, "ContainerList", arg0, arg1)
	ret0, _ := ret[0].([]types.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerList indicates an expected call of ContainerList
func (mr *MockLocalEndpointsStarterMockRecorder) ContainerList(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerList", reflect.TypeOf((*MockLocalEndpointsStarter)(nil).ContainerList), arg0, arg1)
}

// ContainerStart mocks base method
func (m *MockLocalEndpointsStarter) ContainerStart(arg0 context.Context, arg1 string, arg2 types.ContainerStartOptions) error {
	ret := m.ctrl.Call(m, "ContainerStart", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ContainerStart indicates an expected call of ContainerStart
func (mr *MockLocalEndpointsStarterMockRecorder) ContainerStart(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerStart", reflect.TypeOf((*MockLocalEndpointsStarter)(nil).ContainerStart), arg0, arg1, arg2)
}

// NetworkCreate mocks base method
func (m *MockLocalEndpointsStarter) NetworkCreate(arg0 context.Context, arg1 string, arg2 types.NetworkCreate) (types.NetworkCreateResponse, error) {
	ret := m.ctrl.Call(m, "NetworkCreate", arg0, arg1, arg2)
	ret0, _ := ret[0].(types.NetworkCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetworkCreate indicates an expected call of NetworkCreate
func (mr *MockLocalEndpointsStarterMockRecorder) NetworkCreate(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkCreate", reflect.TypeOf((*MockLocalEndpointsStarter)(nil).NetworkCreate), arg0, arg1, arg2)
}

// NetworkInspect mocks base method
func (m *MockLocalEndpointsStarter) NetworkInspect(arg0 context.Context, arg1 string, arg2 types.NetworkInspectOptions) (types.NetworkResource, error) {
	ret := m.ctrl.Call(m, "NetworkInspect", arg0, arg1, arg2)
	ret0, _ := ret[0].(types.NetworkResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetworkInspect indicates an expected call of NetworkInspect
func (mr *MockLocalEndpointsStarterMockRecorder) NetworkInspect(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkInspect", reflect.TypeOf((*MockLocalEndpointsStarter)(nil).NetworkInspect), arg0, arg1, arg2)
}