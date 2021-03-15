// Code generated by MockGen. DO NOT EDIT.
// Source: calculator.go

// Package mock is a generated GoMock package.
package mock

import (
	parser "github.com/alexandear/scalc/internal/parser"
	scalc "github.com/alexandear/scalc/pkg/scalc"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockParser is a mock of Parser interface
type MockParser struct {
	ctrl     *gomock.Controller
	recorder *MockParserMockRecorder
}

// MockParserMockRecorder is the mock recorder for MockParser
type MockParserMockRecorder struct {
	mock *MockParser
}

// NewMockParser creates a new mock instance
func NewMockParser(ctrl *gomock.Controller) *MockParser {
	mock := &MockParser{ctrl: ctrl}
	mock.recorder = &MockParserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockParser) EXPECT() *MockParserMockRecorder {
	return m.recorder
}

// Parse mocks base method
func (m *MockParser) Parse(s string) (*parser.Expression, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse", s)
	ret0, _ := ret[0].(*parser.Expression)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parse indicates an expected call of Parse
func (mr *MockParserMockRecorder) Parse(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockParser)(nil).Parse), s)
}

// MockFileToIterator is a mock of FileToIterator interface
type MockFileToIterator struct {
	ctrl     *gomock.Controller
	recorder *MockFileToIteratorMockRecorder
}

// MockFileToIteratorMockRecorder is the mock recorder for MockFileToIterator
type MockFileToIteratorMockRecorder struct {
	mock *MockFileToIterator
}

// NewMockFileToIterator creates a new mock instance
func NewMockFileToIterator(ctrl *gomock.Controller) *MockFileToIterator {
	mock := &MockFileToIterator{ctrl: ctrl}
	mock.recorder = &MockFileToIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFileToIterator) EXPECT() *MockFileToIteratorMockRecorder {
	return m.recorder
}

// Iterator mocks base method
func (m *MockFileToIterator) Iterator(file string) (scalc.Iterator, io.Closer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iterator", file)
	ret0, _ := ret[0].(scalc.Iterator)
	ret1, _ := ret[1].(io.Closer)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Iterator indicates an expected call of Iterator
func (mr *MockFileToIteratorMockRecorder) Iterator(file interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iterator", reflect.TypeOf((*MockFileToIterator)(nil).Iterator), file)
}
