/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: ../scalesets.go

// Package mock_scalesets is a generated GoMock package.
package mock_scalesets

import (
	context "context"
	reflect "reflect"

	autorest "github.com/Azure/go-autorest/autorest"
	gomock "github.com/golang/mock/gomock"
	v1beta1 "sigs.k8s.io/cluster-api-provider-azure/api/v1beta1"
	azure "sigs.k8s.io/cluster-api-provider-azure/azure"
	v1beta10 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// MockScaleSetScope is a mock of ScaleSetScope interface.
type MockScaleSetScope struct {
	ctrl     *gomock.Controller
	recorder *MockScaleSetScopeMockRecorder
}

// MockScaleSetScopeMockRecorder is the mock recorder for MockScaleSetScope.
type MockScaleSetScopeMockRecorder struct {
	mock *MockScaleSetScope
}

// NewMockScaleSetScope creates a new mock instance.
func NewMockScaleSetScope(ctrl *gomock.Controller) *MockScaleSetScope {
	mock := &MockScaleSetScope{ctrl: ctrl}
	mock.recorder = &MockScaleSetScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScaleSetScope) EXPECT() *MockScaleSetScopeMockRecorder {
	return m.recorder
}

// AdditionalTags mocks base method.
func (m *MockScaleSetScope) AdditionalTags() v1beta1.Tags {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdditionalTags")
	ret0, _ := ret[0].(v1beta1.Tags)
	return ret0
}

// AdditionalTags indicates an expected call of AdditionalTags.
func (mr *MockScaleSetScopeMockRecorder) AdditionalTags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdditionalTags", reflect.TypeOf((*MockScaleSetScope)(nil).AdditionalTags))
}

// Authorizer mocks base method.
func (m *MockScaleSetScope) Authorizer() autorest.Authorizer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorizer")
	ret0, _ := ret[0].(autorest.Authorizer)
	return ret0
}

// Authorizer indicates an expected call of Authorizer.
func (mr *MockScaleSetScopeMockRecorder) Authorizer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorizer", reflect.TypeOf((*MockScaleSetScope)(nil).Authorizer))
}

// AvailabilitySetEnabled mocks base method.
func (m *MockScaleSetScope) AvailabilitySetEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvailabilitySetEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AvailabilitySetEnabled indicates an expected call of AvailabilitySetEnabled.
func (mr *MockScaleSetScopeMockRecorder) AvailabilitySetEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailabilitySetEnabled", reflect.TypeOf((*MockScaleSetScope)(nil).AvailabilitySetEnabled))
}

// BaseURI mocks base method.
func (m *MockScaleSetScope) BaseURI() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseURI")
	ret0, _ := ret[0].(string)
	return ret0
}

// BaseURI indicates an expected call of BaseURI.
func (mr *MockScaleSetScopeMockRecorder) BaseURI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseURI", reflect.TypeOf((*MockScaleSetScope)(nil).BaseURI))
}

// ClientID mocks base method.
func (m *MockScaleSetScope) ClientID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientID indicates an expected call of ClientID.
func (mr *MockScaleSetScopeMockRecorder) ClientID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientID", reflect.TypeOf((*MockScaleSetScope)(nil).ClientID))
}

// ClientSecret mocks base method.
func (m *MockScaleSetScope) ClientSecret() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientSecret")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientSecret indicates an expected call of ClientSecret.
func (mr *MockScaleSetScopeMockRecorder) ClientSecret() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientSecret", reflect.TypeOf((*MockScaleSetScope)(nil).ClientSecret))
}

// CloudEnvironment mocks base method.
func (m *MockScaleSetScope) CloudEnvironment() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudEnvironment")
	ret0, _ := ret[0].(string)
	return ret0
}

// CloudEnvironment indicates an expected call of CloudEnvironment.
func (mr *MockScaleSetScopeMockRecorder) CloudEnvironment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudEnvironment", reflect.TypeOf((*MockScaleSetScope)(nil).CloudEnvironment))
}

// CloudProviderConfigOverrides mocks base method.
func (m *MockScaleSetScope) CloudProviderConfigOverrides() *v1beta1.CloudProviderConfigOverrides {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudProviderConfigOverrides")
	ret0, _ := ret[0].(*v1beta1.CloudProviderConfigOverrides)
	return ret0
}

// CloudProviderConfigOverrides indicates an expected call of CloudProviderConfigOverrides.
func (mr *MockScaleSetScopeMockRecorder) CloudProviderConfigOverrides() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudProviderConfigOverrides", reflect.TypeOf((*MockScaleSetScope)(nil).CloudProviderConfigOverrides))
}

// ClusterName mocks base method.
func (m *MockScaleSetScope) ClusterName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClusterName indicates an expected call of ClusterName.
func (mr *MockScaleSetScopeMockRecorder) ClusterName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterName", reflect.TypeOf((*MockScaleSetScope)(nil).ClusterName))
}

// DeleteLongRunningOperationState mocks base method.
func (m *MockScaleSetScope) DeleteLongRunningOperationState(arg0, arg1, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteLongRunningOperationState", arg0, arg1, arg2)
}

// DeleteLongRunningOperationState indicates an expected call of DeleteLongRunningOperationState.
func (mr *MockScaleSetScopeMockRecorder) DeleteLongRunningOperationState(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLongRunningOperationState", reflect.TypeOf((*MockScaleSetScope)(nil).DeleteLongRunningOperationState), arg0, arg1, arg2)
}

// ExtendedLocation mocks base method.
func (m *MockScaleSetScope) ExtendedLocation() *v1beta1.ExtendedLocationSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtendedLocation")
	ret0, _ := ret[0].(*v1beta1.ExtendedLocationSpec)
	return ret0
}

// ExtendedLocation indicates an expected call of ExtendedLocation.
func (mr *MockScaleSetScopeMockRecorder) ExtendedLocation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtendedLocation", reflect.TypeOf((*MockScaleSetScope)(nil).ExtendedLocation))
}

// FailureDomains mocks base method.
func (m *MockScaleSetScope) FailureDomains() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FailureDomains")
	ret0, _ := ret[0].([]string)
	return ret0
}

// FailureDomains indicates an expected call of FailureDomains.
func (mr *MockScaleSetScopeMockRecorder) FailureDomains() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailureDomains", reflect.TypeOf((*MockScaleSetScope)(nil).FailureDomains))
}

// GetBootstrapData mocks base method.
func (m *MockScaleSetScope) GetBootstrapData(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBootstrapData", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBootstrapData indicates an expected call of GetBootstrapData.
func (mr *MockScaleSetScopeMockRecorder) GetBootstrapData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBootstrapData", reflect.TypeOf((*MockScaleSetScope)(nil).GetBootstrapData), arg0)
}

// GetLongRunningOperationState mocks base method.
func (m *MockScaleSetScope) GetLongRunningOperationState(arg0, arg1, arg2 string) *v1beta1.Future {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLongRunningOperationState", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.Future)
	return ret0
}

// GetLongRunningOperationState indicates an expected call of GetLongRunningOperationState.
func (mr *MockScaleSetScopeMockRecorder) GetLongRunningOperationState(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLongRunningOperationState", reflect.TypeOf((*MockScaleSetScope)(nil).GetLongRunningOperationState), arg0, arg1, arg2)
}

// GetVMImage mocks base method.
func (m *MockScaleSetScope) GetVMImage(arg0 context.Context) (*v1beta1.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVMImage", arg0)
	ret0, _ := ret[0].(*v1beta1.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVMImage indicates an expected call of GetVMImage.
func (mr *MockScaleSetScopeMockRecorder) GetVMImage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVMImage", reflect.TypeOf((*MockScaleSetScope)(nil).GetVMImage), arg0)
}

// HashKey mocks base method.
func (m *MockScaleSetScope) HashKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// HashKey indicates an expected call of HashKey.
func (mr *MockScaleSetScopeMockRecorder) HashKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashKey", reflect.TypeOf((*MockScaleSetScope)(nil).HashKey))
}

// Location mocks base method.
func (m *MockScaleSetScope) Location() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Location")
	ret0, _ := ret[0].(string)
	return ret0
}

// Location indicates an expected call of Location.
func (mr *MockScaleSetScopeMockRecorder) Location() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Location", reflect.TypeOf((*MockScaleSetScope)(nil).Location))
}

// MaxSurge mocks base method.
func (m *MockScaleSetScope) MaxSurge() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxSurge")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MaxSurge indicates an expected call of MaxSurge.
func (mr *MockScaleSetScopeMockRecorder) MaxSurge() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxSurge", reflect.TypeOf((*MockScaleSetScope)(nil).MaxSurge))
}

// ResourceGroup mocks base method.
func (m *MockScaleSetScope) ResourceGroup() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceGroup")
	ret0, _ := ret[0].(string)
	return ret0
}

// ResourceGroup indicates an expected call of ResourceGroup.
func (mr *MockScaleSetScopeMockRecorder) ResourceGroup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceGroup", reflect.TypeOf((*MockScaleSetScope)(nil).ResourceGroup))
}

// SaveVMImageToStatus mocks base method.
func (m *MockScaleSetScope) SaveVMImageToStatus(arg0 *v1beta1.Image) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SaveVMImageToStatus", arg0)
}

// SaveVMImageToStatus indicates an expected call of SaveVMImageToStatus.
func (mr *MockScaleSetScopeMockRecorder) SaveVMImageToStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveVMImageToStatus", reflect.TypeOf((*MockScaleSetScope)(nil).SaveVMImageToStatus), arg0)
}

// ScaleSetSpec mocks base method.
func (m *MockScaleSetScope) ScaleSetSpec() azure.ScaleSetSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScaleSetSpec")
	ret0, _ := ret[0].(azure.ScaleSetSpec)
	return ret0
}

// ScaleSetSpec indicates an expected call of ScaleSetSpec.
func (mr *MockScaleSetScopeMockRecorder) ScaleSetSpec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScaleSetSpec", reflect.TypeOf((*MockScaleSetScope)(nil).ScaleSetSpec))
}

// SetAnnotation mocks base method.
func (m *MockScaleSetScope) SetAnnotation(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAnnotation", arg0, arg1)
}

// SetAnnotation indicates an expected call of SetAnnotation.
func (mr *MockScaleSetScopeMockRecorder) SetAnnotation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAnnotation", reflect.TypeOf((*MockScaleSetScope)(nil).SetAnnotation), arg0, arg1)
}

// SetLongRunningOperationState mocks base method.
func (m *MockScaleSetScope) SetLongRunningOperationState(arg0 *v1beta1.Future) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLongRunningOperationState", arg0)
}

// SetLongRunningOperationState indicates an expected call of SetLongRunningOperationState.
func (mr *MockScaleSetScopeMockRecorder) SetLongRunningOperationState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLongRunningOperationState", reflect.TypeOf((*MockScaleSetScope)(nil).SetLongRunningOperationState), arg0)
}

// SetProviderID mocks base method.
func (m *MockScaleSetScope) SetProviderID(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetProviderID", arg0)
}

// SetProviderID indicates an expected call of SetProviderID.
func (mr *MockScaleSetScopeMockRecorder) SetProviderID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetProviderID", reflect.TypeOf((*MockScaleSetScope)(nil).SetProviderID), arg0)
}

// SetVMSSState mocks base method.
func (m *MockScaleSetScope) SetVMSSState(arg0 *azure.VMSS) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetVMSSState", arg0)
}

// SetVMSSState indicates an expected call of SetVMSSState.
func (mr *MockScaleSetScopeMockRecorder) SetVMSSState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVMSSState", reflect.TypeOf((*MockScaleSetScope)(nil).SetVMSSState), arg0)
}

// SubscriptionID mocks base method.
func (m *MockScaleSetScope) SubscriptionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscriptionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SubscriptionID indicates an expected call of SubscriptionID.
func (mr *MockScaleSetScopeMockRecorder) SubscriptionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionID", reflect.TypeOf((*MockScaleSetScope)(nil).SubscriptionID))
}

// TenantID mocks base method.
func (m *MockScaleSetScope) TenantID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TenantID")
	ret0, _ := ret[0].(string)
	return ret0
}

// TenantID indicates an expected call of TenantID.
func (mr *MockScaleSetScopeMockRecorder) TenantID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantID", reflect.TypeOf((*MockScaleSetScope)(nil).TenantID))
}

// UpdateDeleteStatus mocks base method.
func (m *MockScaleSetScope) UpdateDeleteStatus(arg0 v1beta10.ConditionType, arg1 string, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateDeleteStatus", arg0, arg1, arg2)
}

// UpdateDeleteStatus indicates an expected call of UpdateDeleteStatus.
func (mr *MockScaleSetScopeMockRecorder) UpdateDeleteStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeleteStatus", reflect.TypeOf((*MockScaleSetScope)(nil).UpdateDeleteStatus), arg0, arg1, arg2)
}

// UpdatePatchStatus mocks base method.
func (m *MockScaleSetScope) UpdatePatchStatus(arg0 v1beta10.ConditionType, arg1 string, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatePatchStatus", arg0, arg1, arg2)
}

// UpdatePatchStatus indicates an expected call of UpdatePatchStatus.
func (mr *MockScaleSetScopeMockRecorder) UpdatePatchStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePatchStatus", reflect.TypeOf((*MockScaleSetScope)(nil).UpdatePatchStatus), arg0, arg1, arg2)
}

// UpdatePutStatus mocks base method.
func (m *MockScaleSetScope) UpdatePutStatus(arg0 v1beta10.ConditionType, arg1 string, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatePutStatus", arg0, arg1, arg2)
}

// UpdatePutStatus indicates an expected call of UpdatePutStatus.
func (mr *MockScaleSetScopeMockRecorder) UpdatePutStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePutStatus", reflect.TypeOf((*MockScaleSetScope)(nil).UpdatePutStatus), arg0, arg1, arg2)
}

// VMSSExtensionSpecs mocks base method.
func (m *MockScaleSetScope) VMSSExtensionSpecs() []azure.ResourceSpecGetter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VMSSExtensionSpecs")
	ret0, _ := ret[0].([]azure.ResourceSpecGetter)
	return ret0
}

// VMSSExtensionSpecs indicates an expected call of VMSSExtensionSpecs.
func (mr *MockScaleSetScopeMockRecorder) VMSSExtensionSpecs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VMSSExtensionSpecs", reflect.TypeOf((*MockScaleSetScope)(nil).VMSSExtensionSpecs))
}
