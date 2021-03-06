// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"

// EventHandler is an autogenerated mock type for the EventHandler type
type EventHandler struct {
	mock.Mock
}

// OnAdd provides a mock function with given fields: obj
func (_m *EventHandler) OnAdd(obj interface{}) {
	_m.Called(obj)
}

// OnDelete provides a mock function with given fields: obj
func (_m *EventHandler) OnDelete(obj interface{}) {
	_m.Called(obj)
}

// OnStatus provides a mock function with given fields:
func (_m *EventHandler) OnStatus() {
	_m.Called()
}

// OnUpdate provides a mock function with given fields: oldObj, newObj
func (_m *EventHandler) OnUpdate(oldObj interface{}, newObj interface{}) {
	_m.Called(oldObj, newObj)
}
