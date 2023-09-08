package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsPstnOnlineMeetingDialoutReportOperationPredicate struct {
	Currency           *string
	DestinationContext *string
	ODataType          *string
	TotalCallCharge    *float64
	TotalCallSeconds   *int64
	TotalCalls         *int64
	UsageLocation      *string
	UserDisplayName    *string
	UserId             *string
	UserPrincipalName  *string
}

func (p CallRecordsPstnOnlineMeetingDialoutReportOperationPredicate) Matches(input CallRecordsPstnOnlineMeetingDialoutReport) bool {

	if p.Currency != nil && (input.Currency == nil || *p.Currency != *input.Currency) {
		return false
	}

	if p.DestinationContext != nil && (input.DestinationContext == nil || *p.DestinationContext != *input.DestinationContext) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TotalCallCharge != nil && (input.TotalCallCharge == nil || *p.TotalCallCharge != *input.TotalCallCharge) {
		return false
	}

	if p.TotalCallSeconds != nil && (input.TotalCallSeconds == nil || *p.TotalCallSeconds != *input.TotalCallSeconds) {
		return false
	}

	if p.TotalCalls != nil && (input.TotalCalls == nil || *p.TotalCalls != *input.TotalCalls) {
		return false
	}

	if p.UsageLocation != nil && (input.UsageLocation == nil || *p.UsageLocation != *input.UsageLocation) {
		return false
	}

	if p.UserDisplayName != nil && (input.UserDisplayName == nil || *p.UserDisplayName != *input.UserDisplayName) {
		return false
	}

	if p.UserId != nil && (input.UserId == nil || *p.UserId != *input.UserId) {
		return false
	}

	if p.UserPrincipalName != nil && (input.UserPrincipalName == nil || *p.UserPrincipalName != *input.UserPrincipalName) {
		return false
	}

	return true
}
