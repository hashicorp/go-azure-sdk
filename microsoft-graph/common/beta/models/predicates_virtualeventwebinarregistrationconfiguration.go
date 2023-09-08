package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventWebinarRegistrationConfigurationOperationPredicate struct {
	Capacity                *int64
	Id                      *string
	IsManualApprovalEnabled *bool
	IsWaitlistEnabled       *bool
	ODataType               *string
	RegistrationWebUrl      *string
}

func (p VirtualEventWebinarRegistrationConfigurationOperationPredicate) Matches(input VirtualEventWebinarRegistrationConfiguration) bool {

	if p.Capacity != nil && (input.Capacity == nil || *p.Capacity != *input.Capacity) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsManualApprovalEnabled != nil && (input.IsManualApprovalEnabled == nil || *p.IsManualApprovalEnabled != *input.IsManualApprovalEnabled) {
		return false
	}

	if p.IsWaitlistEnabled != nil && (input.IsWaitlistEnabled == nil || *p.IsWaitlistEnabled != *input.IsWaitlistEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RegistrationWebUrl != nil && (input.RegistrationWebUrl == nil || *p.RegistrationWebUrl != *input.RegistrationWebUrl) {
		return false
	}

	return true
}
