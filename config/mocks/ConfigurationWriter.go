// Code generated by mockery v1.0.0
package mocks

import config "github.com/CloudHub360/ch360.go/config"
import mock "github.com/stretchr/testify/mock"

// ConfigurationWriter is an autogenerated mock type for the ConfigurationWriter type
type ConfigurationWriter struct {
	mock.Mock
}

// WriteConfiguration provides a mock function with given fields: configuration
func (_m *ConfigurationWriter) WriteConfiguration(configuration *config.Configuration) error {
	ret := _m.Called(configuration)

	var r0 error
	if rf, ok := ret.Get(0).(func(*config.Configuration) error); ok {
		r0 = rf(configuration)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
