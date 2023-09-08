package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagementAppHealthSummaryOperationPredicate struct {
	HealthyDeviceCount   *int64
	Id                   *string
	ODataType            *string
	UnhealthyDeviceCount *int64
	UnknownDeviceCount   *int64
}

func (p WindowsManagementAppHealthSummaryOperationPredicate) Matches(input WindowsManagementAppHealthSummary) bool {

	if p.HealthyDeviceCount != nil && (input.HealthyDeviceCount == nil || *p.HealthyDeviceCount != *input.HealthyDeviceCount) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UnhealthyDeviceCount != nil && (input.UnhealthyDeviceCount == nil || *p.UnhealthyDeviceCount != *input.UnhealthyDeviceCount) {
		return false
	}

	if p.UnknownDeviceCount != nil && (input.UnknownDeviceCount == nil || *p.UnknownDeviceCount != *input.UnknownDeviceCount) {
		return false
	}

	return true
}
