// Code generated by mockery v1.0.0
package mocks

import failover "github.com/spotahome/redis-operator/pkg/failover"
import mock "github.com/stretchr/testify/mock"
import v1beta1 "k8s.io/client-go/pkg/apis/apps/v1beta1"

// Transformer is an autogenerated mock type for the Transformer type
type Transformer struct {
	mock.Mock
}

// DeploymentToSentinelSettings provides a mock function with given fields: _a0
func (_m *Transformer) DeploymentToSentinelSettings(_a0 *v1beta1.Deployment) (*failover.SentinelSettings, error) {
	ret := _m.Called(_a0)

	var r0 *failover.SentinelSettings
	if rf, ok := ret.Get(0).(func(*v1beta1.Deployment) *failover.SentinelSettings); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*failover.SentinelSettings)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1beta1.Deployment) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StatefulsetToRedisSettings provides a mock function with given fields: _a0
func (_m *Transformer) StatefulsetToRedisSettings(_a0 *v1beta1.StatefulSet) (*failover.RedisSettings, error) {
	ret := _m.Called(_a0)

	var r0 *failover.RedisSettings
	if rf, ok := ret.Get(0).(func(*v1beta1.StatefulSet) *failover.RedisSettings); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*failover.RedisSettings)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1beta1.StatefulSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
