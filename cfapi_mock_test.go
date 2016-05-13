// Automatically generated by MockGen. DO NOT EDIT!
// Source: cfapi.go

package main

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of CloudApi interface
type MockCloudApi struct {
	ctrl     *gomock.Controller
	recorder *_MockCloudApiRecorder
}

// Recorder for MockCloudApi (not exported)
type _MockCloudApiRecorder struct {
	mock *MockCloudApi
}

func NewMockCloudApi(ctrl *gomock.Controller) *MockCloudApi {
	mock := &MockCloudApi{ctrl: ctrl}
	mock.recorder = &_MockCloudApiRecorder{mock}
	return mock
}

func (_m *MockCloudApi) EXPECT() *_MockCloudApiRecorder {
	return _m.recorder
}

func (_m *MockCloudApi) GetOrgIdAndSpaceIdFromCfByServiceInstanceId(service_instance_id string) (string, string, error) {
	ret := _m.ctrl.Call(_m, "GetOrgIdAndSpaceIdFromCfByServiceInstanceId", service_instance_id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockCloudApiRecorder) GetOrgIdAndSpaceIdFromCfByServiceInstanceId(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetOrgIdAndSpaceIdFromCfByServiceInstanceId", arg0)
}

func (_m *MockCloudApi) GetSpaceDetailsFromCfBySpaceId(space_id string) (CfSpaceDetails, error) {
	ret := _m.ctrl.Call(_m, "GetSpaceDetailsFromCfBySpaceId", space_id)
	ret0, _ := ret[0].(CfSpaceDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCloudApiRecorder) GetSpaceDetailsFromCfBySpaceId(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSpaceDetailsFromCfBySpaceId", arg0)
}

func (_m *MockCloudApi) GetInstanceDetailsFromCfById(instance_id string) (CfInstanceDetails, error) {
	ret := _m.ctrl.Call(_m, "GetInstanceDetailsFromCfById", instance_id)
	ret0, _ := ret[0].(CfInstanceDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCloudApiRecorder) GetInstanceDetailsFromCfById(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetInstanceDetailsFromCfById", arg0)
}

func (_m *MockCloudApi) GetServiceBrokerByName(brokerName string) (FindServiceBrokerResponse, error) {
	ret := _m.ctrl.Call(_m, "GetServiceBrokerByName", brokerName)
	ret0, _ := ret[0].(FindServiceBrokerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCloudApiRecorder) GetServiceBrokerByName(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetServiceBrokerByName", arg0)
}

func (_m *MockCloudApi) UpdateServiceBroker() (ServiceBroker, error) {
	ret := _m.ctrl.Call(_m, "UpdateServiceBroker")
	ret0, _ := ret[0].(ServiceBroker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCloudApiRecorder) UpdateServiceBroker() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateServiceBroker")
}