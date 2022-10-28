// Code generated by MockGen. DO NOT EDIT.
// Source: poi.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockConnection is a mock of Connection interface.
type MockConnection struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionMockRecorder
}

// MockConnectionMockRecorder is the mock recorder for MockConnection.
type MockConnectionMockRecorder struct {
	mock *MockConnection
}

// NewMockConnection creates a new mock instance.
func NewMockConnection(ctrl *gomock.Controller) *MockConnection {
	mock := &MockConnection{ctrl: ctrl}
	mock.recorder = &MockConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnection) EXPECT() *MockConnectionMockRecorder {
	return m.recorder
}

// CheckRepos mocks base method.
func (m *MockConnection) CheckRepos(ctx context.Context, hostname string, repoNames []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckRepos", ctx, hostname, repoNames)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckRepos indicates an expected call of CheckRepos.
func (mr *MockConnectionMockRecorder) CheckRepos(ctx, hostname, repoNames interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckRepos", reflect.TypeOf((*MockConnection)(nil).CheckRepos), ctx, hostname, repoNames)
}

// CheckoutBranch mocks base method.
func (m *MockConnection) CheckoutBranch(ctx context.Context, branchName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckoutBranch", ctx, branchName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckoutBranch indicates an expected call of CheckoutBranch.
func (mr *MockConnectionMockRecorder) CheckoutBranch(ctx, branchName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckoutBranch", reflect.TypeOf((*MockConnection)(nil).CheckoutBranch), ctx, branchName)
}

// DeleteBranches mocks base method.
func (m *MockConnection) DeleteBranches(ctx context.Context, branchNames []string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBranches", ctx, branchNames)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBranches indicates an expected call of DeleteBranches.
func (mr *MockConnectionMockRecorder) DeleteBranches(ctx, branchNames interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBranches", reflect.TypeOf((*MockConnection)(nil).DeleteBranches), ctx, branchNames)
}

// GetAssociatedRefNames mocks base method.
func (m *MockConnection) GetAssociatedRefNames(ctx context.Context, oid string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAssociatedRefNames", ctx, oid)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAssociatedRefNames indicates an expected call of GetAssociatedRefNames.
func (mr *MockConnectionMockRecorder) GetAssociatedRefNames(ctx, oid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAssociatedRefNames", reflect.TypeOf((*MockConnection)(nil).GetAssociatedRefNames), ctx, oid)
}

// GetBranchNames mocks base method.
func (m *MockConnection) GetBranchNames(ctx context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBranchNames", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBranchNames indicates an expected call of GetBranchNames.
func (mr *MockConnectionMockRecorder) GetBranchNames(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBranchNames", reflect.TypeOf((*MockConnection)(nil).GetBranchNames), ctx)
}

// GetConfig mocks base method.
func (m *MockConnection) GetConfig(ctx context.Context, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig", ctx, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockConnectionMockRecorder) GetConfig(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockConnection)(nil).GetConfig), ctx, key)
}

// GetLog mocks base method.
func (m *MockConnection) GetLog(ctx context.Context, branchName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLog", ctx, branchName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLog indicates an expected call of GetLog.
func (mr *MockConnectionMockRecorder) GetLog(ctx, branchName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLog", reflect.TypeOf((*MockConnection)(nil).GetLog), ctx, branchName)
}

// GetMergedBranchNames mocks base method.
func (m *MockConnection) GetMergedBranchNames(ctx context.Context, remoteName, branchName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMergedBranchNames", ctx, remoteName, branchName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMergedBranchNames indicates an expected call of GetMergedBranchNames.
func (mr *MockConnectionMockRecorder) GetMergedBranchNames(ctx, remoteName, branchName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMergedBranchNames", reflect.TypeOf((*MockConnection)(nil).GetMergedBranchNames), ctx, remoteName, branchName)
}

// GetPullRequests mocks base method.
func (m *MockConnection) GetPullRequests(ctx context.Context, hostname string, repoNames []string, queryHashes string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPullRequests", ctx, hostname, repoNames, queryHashes)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPullRequests indicates an expected call of GetPullRequests.
func (mr *MockConnectionMockRecorder) GetPullRequests(ctx, hostname, repoNames, queryHashes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPullRequests", reflect.TypeOf((*MockConnection)(nil).GetPullRequests), ctx, hostname, repoNames, queryHashes)
}

// GetRemoteNames mocks base method.
func (m *MockConnection) GetRemoteNames(ctx context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRemoteNames", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRemoteNames indicates an expected call of GetRemoteNames.
func (mr *MockConnectionMockRecorder) GetRemoteNames(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRemoteNames", reflect.TypeOf((*MockConnection)(nil).GetRemoteNames), ctx)
}

// GetRepoNames mocks base method.
func (m *MockConnection) GetRepoNames(ctx context.Context, hostname, repoName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepoNames", ctx, hostname, repoName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepoNames indicates an expected call of GetRepoNames.
func (mr *MockConnectionMockRecorder) GetRepoNames(ctx, hostname, repoName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepoNames", reflect.TypeOf((*MockConnection)(nil).GetRepoNames), ctx, hostname, repoName)
}

// GetSshConfig mocks base method.
func (m *MockConnection) GetSshConfig(ctx context.Context, name string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSshConfig", ctx, name)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSshConfig indicates an expected call of GetSshConfig.
func (mr *MockConnectionMockRecorder) GetSshConfig(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSshConfig", reflect.TypeOf((*MockConnection)(nil).GetSshConfig), ctx, name)
}

// GetUncommittedChanges mocks base method.
func (m *MockConnection) GetUncommittedChanges(ctx context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUncommittedChanges", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUncommittedChanges indicates an expected call of GetUncommittedChanges.
func (mr *MockConnectionMockRecorder) GetUncommittedChanges(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUncommittedChanges", reflect.TypeOf((*MockConnection)(nil).GetUncommittedChanges), ctx)
}
