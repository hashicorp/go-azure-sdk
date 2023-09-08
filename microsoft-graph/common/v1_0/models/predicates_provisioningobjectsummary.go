package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningObjectSummaryOperationPredicate struct {
	ActivityDateTime       *string
	ChangeId               *string
	CycleId                *string
	DurationInMilliseconds *int64
	Id                     *string
	JobId                  *string
	ODataType              *string
	TenantId               *string
}

func (p ProvisioningObjectSummaryOperationPredicate) Matches(input ProvisioningObjectSummary) bool {

	if p.ActivityDateTime != nil && (input.ActivityDateTime == nil || *p.ActivityDateTime != *input.ActivityDateTime) {
		return false
	}

	if p.ChangeId != nil && (input.ChangeId == nil || *p.ChangeId != *input.ChangeId) {
		return false
	}

	if p.CycleId != nil && (input.CycleId == nil || *p.CycleId != *input.CycleId) {
		return false
	}

	if p.DurationInMilliseconds != nil && (input.DurationInMilliseconds == nil || *p.DurationInMilliseconds != *input.DurationInMilliseconds) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.JobId != nil && (input.JobId == nil || *p.JobId != *input.JobId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TenantId != nil && (input.TenantId == nil || *p.TenantId != *input.TenantId) {
		return false
	}

	return true
}
