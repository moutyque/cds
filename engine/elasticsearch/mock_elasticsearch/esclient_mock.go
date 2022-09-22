// Code generated by MockGen. DO NOT EDIT.
// Source: types.go

// Package mock_elasticsearch is a generated GoMock package.
package mock_elasticsearch

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	elastic "gopkg.in/olivere/elastic.v6"
)

// MockESClient is a mock of ESClient interface.
type MockESClient struct {
	ctrl     *gomock.Controller
	recorder *MockESClientMockRecorder
}

// MockESClientMockRecorder is the mock recorder for MockESClient.
type MockESClientMockRecorder struct {
	mock *MockESClient
}

// NewMockESClient creates a new mock instance.
func NewMockESClient(ctrl *gomock.Controller) *MockESClient {
	mock := &MockESClient{ctrl: ctrl}
	mock.recorder = &MockESClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockESClient) EXPECT() *MockESClientMockRecorder {
	return m.recorder
}

// IndexDoc mocks base method.
func (m *MockESClient) IndexDoc(ctx context.Context, index, docType, id string, body interface{}) (*elastic.IndexResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IndexDoc", ctx, index, docType, id, body)
	ret0, _ := ret[0].(*elastic.IndexResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IndexDoc indicates an expected call of IndexDoc.
func (mr *MockESClientMockRecorder) IndexDoc(ctx, index, docType, id, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexDoc", reflect.TypeOf((*MockESClient)(nil).IndexDoc), ctx, index, docType, id, body)
}

// Ping mocks base method.
func (m *MockESClient) Ping(ctx context.Context, url string) (*elastic.PingResult, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", ctx, url)
	ret0, _ := ret[0].(*elastic.PingResult)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Ping indicates an expected call of Ping.
func (mr *MockESClientMockRecorder) Ping(ctx, url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockESClient)(nil).Ping), ctx, url)
}

// SearchDoc mocks base method.
func (m *MockESClient) SearchDoc(ctx context.Context, indices []string, docType string, query elastic.Query, sorts []elastic.Sorter, from, size int) (*elastic.SearchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchDoc", ctx, indices, docType, query, sorts, from, size)
	ret0, _ := ret[0].(*elastic.SearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchDoc indicates an expected call of SearchDoc.
func (mr *MockESClientMockRecorder) SearchDoc(ctx, indices, docType, query, sorts, from, size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchDoc", reflect.TypeOf((*MockESClient)(nil).SearchDoc), ctx, indices, docType, query, sorts, from, size)
}