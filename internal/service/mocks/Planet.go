// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"
	model "projeto-star-wars-api-go/internal/model"

	mock "github.com/stretchr/testify/mock"
)

// Planet is an autogenerated mock type for the Planet type
type Planet struct {
	mock.Mock
}

// DeleteById provides a mock function with given fields: ctx, id
func (_m *Planet) DeleteById(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: ctx
func (_m *Planet) FindAll(ctx context.Context) ([]model.PlanetOut, error) {
	ret := _m.Called(ctx)

	var r0 []model.PlanetOut
	if rf, ok := ret.Get(0).(func(context.Context) []model.PlanetOut); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.PlanetOut)
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

// FindById provides a mock function with given fields: ctx, id
func (_m *Planet) FindById(ctx context.Context, id string) (*model.PlanetOut, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.PlanetOut
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.PlanetOut); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PlanetOut)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByName provides a mock function with given fields: ctx, name
func (_m *Planet) FindByName(ctx context.Context, name string) ([]model.PlanetOut, error) {
	ret := _m.Called(ctx, name)

	var r0 []model.PlanetOut
	if rf, ok := ret.Get(0).(func(context.Context, string) []model.PlanetOut); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.PlanetOut)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: parentContext, planet
func (_m *Planet) Save(parentContext context.Context, planet *model.PlanetIn) (string, error) {
	ret := _m.Called(parentContext, planet)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *model.PlanetIn) string); ok {
		r0 = rf(parentContext, planet)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.PlanetIn) error); ok {
		r1 = rf(parentContext, planet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateById provides a mock function with given fields: ctx, p, id
func (_m *Planet) UpdateById(ctx context.Context, p model.PlanetIn, id string) error {
	ret := _m.Called(ctx, p, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.PlanetIn, string) error); ok {
		r0 = rf(ctx, p, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
