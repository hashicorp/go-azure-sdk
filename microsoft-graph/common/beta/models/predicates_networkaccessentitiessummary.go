package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessEntitiesSummaryOperationPredicate struct {
	DeviceCount   *int64
	ODataType     *string
	UserCount     *int64
	WorkloadCount *int64
}

func (p NetworkaccessEntitiesSummaryOperationPredicate) Matches(input NetworkaccessEntitiesSummary) bool {

	if p.DeviceCount != nil && (input.DeviceCount == nil || *p.DeviceCount != *input.DeviceCount) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UserCount != nil && (input.UserCount == nil || *p.UserCount != *input.UserCount) {
		return false
	}

	if p.WorkloadCount != nil && (input.WorkloadCount == nil || *p.WorkloadCount != *input.WorkloadCount) {
		return false
	}

	return true
}
