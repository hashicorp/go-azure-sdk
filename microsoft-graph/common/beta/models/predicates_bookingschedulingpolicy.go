package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingSchedulingPolicyOperationPredicate struct {
	AllowStaffSelection      *bool
	MaximumAdvance           *string
	MinimumLeadTime          *string
	ODataType                *string
	SendConfirmationsToOwner *bool
	TimeSlotInterval         *string
}

func (p BookingSchedulingPolicyOperationPredicate) Matches(input BookingSchedulingPolicy) bool {

	if p.AllowStaffSelection != nil && (input.AllowStaffSelection == nil || *p.AllowStaffSelection != *input.AllowStaffSelection) {
		return false
	}

	if p.MaximumAdvance != nil && (input.MaximumAdvance == nil || *p.MaximumAdvance != *input.MaximumAdvance) {
		return false
	}

	if p.MinimumLeadTime != nil && (input.MinimumLeadTime == nil || *p.MinimumLeadTime != *input.MinimumLeadTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SendConfirmationsToOwner != nil && (input.SendConfirmationsToOwner == nil || *p.SendConfirmationsToOwner != *input.SendConfirmationsToOwner) {
		return false
	}

	if p.TimeSlotInterval != nil && (input.TimeSlotInterval == nil || *p.TimeSlotInterval != *input.TimeSlotInterval) {
		return false
	}

	return true
}
