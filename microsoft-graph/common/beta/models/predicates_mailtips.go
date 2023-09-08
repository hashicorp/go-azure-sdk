package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailTipsOperationPredicate struct {
	CustomMailTip       *string
	DeliveryRestricted  *bool
	ExternalMemberCount *int64
	IsModerated         *bool
	MailboxFull         *bool
	MaxMessageSize      *int64
	ODataType           *string
	TotalMemberCount    *int64
}

func (p MailTipsOperationPredicate) Matches(input MailTips) bool {

	if p.CustomMailTip != nil && (input.CustomMailTip == nil || *p.CustomMailTip != *input.CustomMailTip) {
		return false
	}

	if p.DeliveryRestricted != nil && (input.DeliveryRestricted == nil || *p.DeliveryRestricted != *input.DeliveryRestricted) {
		return false
	}

	if p.ExternalMemberCount != nil && (input.ExternalMemberCount == nil || *p.ExternalMemberCount != *input.ExternalMemberCount) {
		return false
	}

	if p.IsModerated != nil && (input.IsModerated == nil || *p.IsModerated != *input.IsModerated) {
		return false
	}

	if p.MailboxFull != nil && (input.MailboxFull == nil || *p.MailboxFull != *input.MailboxFull) {
		return false
	}

	if p.MaxMessageSize != nil && (input.MaxMessageSize == nil || *p.MaxMessageSize != *input.MaxMessageSize) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TotalMemberCount != nil && (input.TotalMemberCount == nil || *p.TotalMemberCount != *input.TotalMemberCount) {
		return false
	}

	return true
}
