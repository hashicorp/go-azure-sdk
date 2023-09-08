package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AllDevicesAssignmentTargetOperationPredicate struct {
	DeviceAndAppManagementAssignmentFilterId *string
	ODataType                                *string
}

func (p AllDevicesAssignmentTargetOperationPredicate) Matches(input AllDevicesAssignmentTarget) bool {

	if p.DeviceAndAppManagementAssignmentFilterId != nil && (input.DeviceAndAppManagementAssignmentFilterId == nil || *p.DeviceAndAppManagementAssignmentFilterId != *input.DeviceAndAppManagementAssignmentFilterId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
