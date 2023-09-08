package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsDeviceCompliancePolicySettingStateSummaryOperationPredicate struct {
	ConflictDeviceCount      *int64
	ErrorDeviceCount         *int64
	FailedDeviceCount        *int64
	Id                       *string
	IntuneAccountId          *string
	IntuneSettingId          *string
	LastRefreshedDateTime    *string
	NotApplicableDeviceCount *int64
	ODataType                *string
	PendingDeviceCount       *int64
	PolicyType               *string
	SettingName              *string
	SucceededDeviceCount     *int64
	TenantDisplayName        *string
	TenantId                 *string
}

func (p ManagedTenantsDeviceCompliancePolicySettingStateSummaryOperationPredicate) Matches(input ManagedTenantsDeviceCompliancePolicySettingStateSummary) bool {

	if p.ConflictDeviceCount != nil && (input.ConflictDeviceCount == nil || *p.ConflictDeviceCount != *input.ConflictDeviceCount) {
		return false
	}

	if p.ErrorDeviceCount != nil && (input.ErrorDeviceCount == nil || *p.ErrorDeviceCount != *input.ErrorDeviceCount) {
		return false
	}

	if p.FailedDeviceCount != nil && (input.FailedDeviceCount == nil || *p.FailedDeviceCount != *input.FailedDeviceCount) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IntuneAccountId != nil && (input.IntuneAccountId == nil || *p.IntuneAccountId != *input.IntuneAccountId) {
		return false
	}

	if p.IntuneSettingId != nil && (input.IntuneSettingId == nil || *p.IntuneSettingId != *input.IntuneSettingId) {
		return false
	}

	if p.LastRefreshedDateTime != nil && (input.LastRefreshedDateTime == nil || *p.LastRefreshedDateTime != *input.LastRefreshedDateTime) {
		return false
	}

	if p.NotApplicableDeviceCount != nil && (input.NotApplicableDeviceCount == nil || *p.NotApplicableDeviceCount != *input.NotApplicableDeviceCount) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PendingDeviceCount != nil && (input.PendingDeviceCount == nil || *p.PendingDeviceCount != *input.PendingDeviceCount) {
		return false
	}

	if p.PolicyType != nil && (input.PolicyType == nil || *p.PolicyType != *input.PolicyType) {
		return false
	}

	if p.SettingName != nil && (input.SettingName == nil || *p.SettingName != *input.SettingName) {
		return false
	}

	if p.SucceededDeviceCount != nil && (input.SucceededDeviceCount == nil || *p.SucceededDeviceCount != *input.SucceededDeviceCount) {
		return false
	}

	if p.TenantDisplayName != nil && (input.TenantDisplayName == nil || *p.TenantDisplayName != *input.TenantDisplayName) {
		return false
	}

	if p.TenantId != nil && (input.TenantId == nil || *p.TenantId != *input.TenantId) {
		return false
	}

	return true
}
