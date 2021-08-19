// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	command "github.com/bitrise-io/go-utils/command"
	mock "github.com/stretchr/testify/mock"
)

// Factory is an autogenerated mock type for the Factory type
type Factory struct {
	mock.Mock
}

// Create provides a mock function with given fields: name, args, opts
func (_m *Factory) Create(name string, args []string, opts *command.Opts) command.Command {
	ret := _m.Called(name, args, opts)

	var r0 command.Command
	if rf, ok := ret.Get(0).(func(string, []string, *command.Opts) command.Command); ok {
		r0 = rf(name, args, opts)
	} else {
		if ret.Get(0) != nil {
			r0, ok = ret.Get(0).(command.Command)
			if !ok {
				return nil
			}
		}
	}

	return r0
}
