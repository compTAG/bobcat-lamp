// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import (
	msg "github.com/comptag/bobcat-lamp/internal/msg"
	mock "github.com/stretchr/testify/mock"
)

// Messenger is an autogenerated mock type for the Messenger type
type Messenger struct {
	mock.Mock
}

// Send provides a mock function with given fields: _a0
func (_m *Messenger) Send(_a0 msg.Message) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(msg.Message) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(msg.Message) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
