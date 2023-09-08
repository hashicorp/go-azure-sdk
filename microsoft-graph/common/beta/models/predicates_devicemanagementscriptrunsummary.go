package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementScriptRunSummaryOperationPredicate struct {
	ErrorDeviceCount   *int64
	ErrorUserCount     *int64
	Id                 *string
	ODataType          *string
	SuccessDeviceCount *int64
	SuccessUserCount   *int64
}

func (p DeviceManagementScriptRunSummaryOperationPredicate) Matches(input DeviceManagementScriptRunSummary) bool {

	if p.ErrorDeviceCount != nil && (input.ErrorDeviceCount == nil || *p.ErrorDeviceCount != *input.ErrorDeviceCount) {
		return false
	}

	if p.ErrorUserCount != nil && (input.ErrorUserCount == nil || *p.ErrorUserCount != *input.ErrorUserCount) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SuccessDeviceCount != nil && (input.SuccessDeviceCount == nil || *p.SuccessDeviceCount != *input.SuccessDeviceCount) {
		return false
	}

	if p.SuccessUserCount != nil && (input.SuccessUserCount == nil || *p.SuccessUserCount != *input.SuccessUserCount) {
		return false
	}

	return true
}
