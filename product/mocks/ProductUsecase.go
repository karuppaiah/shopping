// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import model "github.com/karuppaiah/shopping/model"

// ProductUsecase is an autogenerated mock type for the ProductUsecase type
type ProductUsecase struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, id
func (_m *ProductUsecase) Delete(ctx context.Context, id int) (bool, error) {
	ret := _m.Called(ctx, id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, int) bool); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: ctx
func (_m *ProductUsecase) Fetch(ctx context.Context) ([]*model.Product, error) {
	ret := _m.Called(ctx)

	var r0 []*model.Product
	if rf, ok := ret.Get(0).(func(context.Context) []*model.Product); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, a
func (_m *ProductUsecase) Store(ctx context.Context, a *model.Product) (int64, error) {
	ret := _m.Called(ctx, a)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *model.Product) int64); ok {
		r0 = rf(ctx, a)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.Product) error); ok {
		r1 = rf(ctx, a)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
