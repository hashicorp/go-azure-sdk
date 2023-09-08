package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RetireScheduledManagedDeviceOperationPredicate struct {
	DeviceCompliancePolicyId   *string
	DeviceCompliancePolicyName *string
	Id                         *string
	ManagedDeviceId            *string
	ManagedDeviceName          *string
	ODataType                  *string
	RetireAfterDateTime        *string
}

func (p RetireScheduledManagedDeviceOperationPredicate) Matches(input RetireScheduledManagedDevice) bool {

	if p.DeviceCompliancePolicyId != nil && (input.DeviceCompliancePolicyId == nil || *p.DeviceCompliancePolicyId != *input.DeviceCompliancePolicyId) {
		return false
	}

	if p.DeviceCompliancePolicyName != nil && (input.DeviceCompliancePolicyName == nil || *p.DeviceCompliancePolicyName != *input.DeviceCompliancePolicyName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ManagedDeviceId != nil && (input.ManagedDeviceId == nil || *p.ManagedDeviceId != *input.ManagedDeviceId) {
		return false
	}

	if p.ManagedDeviceName != nil && (input.ManagedDeviceName == nil || *p.ManagedDeviceName != *input.ManagedDeviceName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RetireAfterDateTime != nil && (input.RetireAfterDateTime == nil || *p.RetireAfterDateTime != *input.RetireAfterDateTime) {
		return false
	}

	return true
}
