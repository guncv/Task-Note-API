// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/guncv/tech-exam-software-engineering/models"
	mock "github.com/stretchr/testify/mock"
)

// MockIUserRepository is an autogenerated mock type for the IUserRepository type
type MockIUserRepository struct {
	mock.Mock
}

type MockIUserRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIUserRepository) EXPECT() *MockIUserRepository_Expecter {
	return &MockIUserRepository_Expecter{mock: &_m.Mock}
}

// GetUser provides a mock function with given fields: ctx, email
func (_m *MockIUserRepository) GetUser(ctx context.Context, email string) (*models.User, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for GetUser")
	}

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.User); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIUserRepository_GetUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUser'
type MockIUserRepository_GetUser_Call struct {
	*mock.Call
}

// GetUser is a helper method to define mock.On call
//   - ctx context.Context
//   - email string
func (_e *MockIUserRepository_Expecter) GetUser(ctx interface{}, email interface{}) *MockIUserRepository_GetUser_Call {
	return &MockIUserRepository_GetUser_Call{Call: _e.mock.On("GetUser", ctx, email)}
}

func (_c *MockIUserRepository_GetUser_Call) Run(run func(ctx context.Context, email string)) *MockIUserRepository_GetUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockIUserRepository_GetUser_Call) Return(_a0 *models.User, _a1 error) *MockIUserRepository_GetUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIUserRepository_GetUser_Call) RunAndReturn(run func(context.Context, string) (*models.User, error)) *MockIUserRepository_GetUser_Call {
	_c.Call.Return(run)
	return _c
}

// RegisterUser provides a mock function with given fields: ctx, user
func (_m *MockIUserRepository) RegisterUser(ctx context.Context, user *models.User) error {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for RegisterUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIUserRepository_RegisterUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegisterUser'
type MockIUserRepository_RegisterUser_Call struct {
	*mock.Call
}

// RegisterUser is a helper method to define mock.On call
//   - ctx context.Context
//   - user *models.User
func (_e *MockIUserRepository_Expecter) RegisterUser(ctx interface{}, user interface{}) *MockIUserRepository_RegisterUser_Call {
	return &MockIUserRepository_RegisterUser_Call{Call: _e.mock.On("RegisterUser", ctx, user)}
}

func (_c *MockIUserRepository_RegisterUser_Call) Run(run func(ctx context.Context, user *models.User)) *MockIUserRepository_RegisterUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.User))
	})
	return _c
}

func (_c *MockIUserRepository_RegisterUser_Call) Return(_a0 error) *MockIUserRepository_RegisterUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIUserRepository_RegisterUser_Call) RunAndReturn(run func(context.Context, *models.User) error) *MockIUserRepository_RegisterUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockIUserRepository creates a new instance of MockIUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIUserRepository {
	mock := &MockIUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
