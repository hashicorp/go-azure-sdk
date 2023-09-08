package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessDeviceUsageSummaryOperationPredicate struct {
	ActiveDeviceCount   *int64
	InactiveDeviceCount *int64
	ODataType           *string
	TotalDeviceCount    *int64
}

func (p NetworkaccessDeviceUsageSummaryOperationPredicate) Matches(input NetworkaccessDeviceUsageSummary) bool {

	if p.ActiveDeviceCount != nil && (input.ActiveDeviceCount == nil || *p.ActiveDeviceCount != *input.ActiveDeviceCount) {
		return false
	}

	if p.InactiveDeviceCount != nil && (input.InactiveDeviceCount == nil || *p.InactiveDeviceCount != *input.InactiveDeviceCount) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TotalDeviceCount != nil && (input.TotalDeviceCount == nil || *p.TotalDeviceCount != *input.TotalDeviceCount) {
		return false
	}

	return true
}
