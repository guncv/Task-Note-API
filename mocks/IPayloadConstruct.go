// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"

	log "github.com/guncv/tech-exam-software-engineering/infras/log"
	mock "github.com/stretchr/testify/mock"

	time "time"

	utils "github.com/guncv/tech-exam-software-engineering/utils"
)

// MockIPayloadConstruct is an autogenerated mock type for the IPayloadConstruct type
type MockIPayloadConstruct struct {
	mock.Mock
}

type MockIPayloadConstruct_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIPayloadConstruct) EXPECT() *MockIPayloadConstruct_Expecter {
	return &MockIPayloadConstruct_Expecter{mock: &_m.Mock}
}

// GetAuthPayload provides a mock function with given fields: ctx, _a1
func (_m *MockIPayloadConstruct) GetAuthPayload(ctx context.Context, _a1 *log.Logger) (*utils.Payload, error) {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetAuthPayload")
	}

	var r0 *utils.Payload
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *log.Logger) (*utils.Payload, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *log.Logger) *utils.Payload); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*utils.Payload)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *log.Logger) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIPayloadConstruct_GetAuthPayload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAuthPayload'
type MockIPayloadConstruct_GetAuthPayload_Call struct {
	*mock.Call
}

// GetAuthPayload is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 *log.Logger
func (_e *MockIPayloadConstruct_Expecter) GetAuthPayload(ctx interface{}, _a1 interface{}) *MockIPayloadConstruct_GetAuthPayload_Call {
	return &MockIPayloadConstruct_GetAuthPayload_Call{Call: _e.mock.On("GetAuthPayload", ctx, _a1)}
}

func (_c *MockIPayloadConstruct_GetAuthPayload_Call) Run(run func(ctx context.Context, _a1 *log.Logger)) *MockIPayloadConstruct_GetAuthPayload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*log.Logger))
	})
	return _c
}

func (_c *MockIPayloadConstruct_GetAuthPayload_Call) Return(_a0 *utils.Payload, _a1 error) *MockIPayloadConstruct_GetAuthPayload_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIPayloadConstruct_GetAuthPayload_Call) RunAndReturn(run func(context.Context, *log.Logger) (*utils.Payload, error)) *MockIPayloadConstruct_GetAuthPayload_Call {
	_c.Call.Return(run)
	return _c
}

// NewCreatePayload provides a mock function with given fields: userId, duration
func (_m *MockIPayloadConstruct) NewCreatePayload(userId string, duration time.Duration) (*utils.Payload, error) {
	ret := _m.Called(userId, duration)

	if len(ret) == 0 {
		panic("no return value specified for NewCreatePayload")
	}

	var r0 *utils.Payload
	var r1 error
	if rf, ok := ret.Get(0).(func(string, time.Duration) (*utils.Payload, error)); ok {
		return rf(userId, duration)
	}
	if rf, ok := ret.Get(0).(func(string, time.Duration) *utils.Payload); ok {
		r0 = rf(userId, duration)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*utils.Payload)
		}
	}

	if rf, ok := ret.Get(1).(func(string, time.Duration) error); ok {
		r1 = rf(userId, duration)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIPayloadConstruct_NewCreatePayload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewCreatePayload'
type MockIPayloadConstruct_NewCreatePayload_Call struct {
	*mock.Call
}

// NewCreatePayload is a helper method to define mock.On call
//   - userId string
//   - duration time.Duration
func (_e *MockIPayloadConstruct_Expecter) NewCreatePayload(userId interface{}, duration interface{}) *MockIPayloadConstruct_NewCreatePayload_Call {
	return &MockIPayloadConstruct_NewCreatePayload_Call{Call: _e.mock.On("NewCreatePayload", userId, duration)}
}

func (_c *MockIPayloadConstruct_NewCreatePayload_Call) Run(run func(userId string, duration time.Duration)) *MockIPayloadConstruct_NewCreatePayload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(time.Duration))
	})
	return _c
}

func (_c *MockIPayloadConstruct_NewCreatePayload_Call) Return(_a0 *utils.Payload, _a1 error) *MockIPayloadConstruct_NewCreatePayload_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIPayloadConstruct_NewCreatePayload_Call) RunAndReturn(run func(string, time.Duration) (*utils.Payload, error)) *MockIPayloadConstruct_NewCreatePayload_Call {
	_c.Call.Return(run)
	return _c
}

// Valid provides a mock function with given fields: payload
func (_m *MockIPayloadConstruct) Valid(payload *utils.Payload) error {
	ret := _m.Called(payload)

	if len(ret) == 0 {
		panic("no return value specified for Valid")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*utils.Payload) error); ok {
		r0 = rf(payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIPayloadConstruct_Valid_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Valid'
type MockIPayloadConstruct_Valid_Call struct {
	*mock.Call
}

// Valid is a helper method to define mock.On call
//   - payload *utils.Payload
func (_e *MockIPayloadConstruct_Expecter) Valid(payload interface{}) *MockIPayloadConstruct_Valid_Call {
	return &MockIPayloadConstruct_Valid_Call{Call: _e.mock.On("Valid", payload)}
}

func (_c *MockIPayloadConstruct_Valid_Call) Run(run func(payload *utils.Payload)) *MockIPayloadConstruct_Valid_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*utils.Payload))
	})
	return _c
}

func (_c *MockIPayloadConstruct_Valid_Call) Return(_a0 error) *MockIPayloadConstruct_Valid_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIPayloadConstruct_Valid_Call) RunAndReturn(run func(*utils.Payload) error) *MockIPayloadConstruct_Valid_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockIPayloadConstruct creates a new instance of MockIPayloadConstruct. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIPayloadConstruct(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIPayloadConstruct {
	mock := &MockIPayloadConstruct{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
