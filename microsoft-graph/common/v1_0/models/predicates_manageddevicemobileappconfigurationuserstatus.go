package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceMobileAppConfigurationUserStatusOperationPredicate struct {
	DevicesCount         *int64
	Id                   *string
	LastReportedDateTime *string
	ODataType            *string
	UserDisplayName      *string
	UserPrincipalName    *string
}

func (p ManagedDeviceMobileAppConfigurationUserStatusOperationPredicate) Matches(input ManagedDeviceMobileAppConfigurationUserStatus) bool {

	if p.DevicesCount != nil && (input.DevicesCount == nil || *p.DevicesCount != *input.DevicesCount) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastReportedDateTime != nil && (input.LastReportedDateTime == nil || *p.LastReportedDateTime != *input.LastReportedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UserDisplayName != nil && (input.UserDisplayName == nil || *p.UserDisplayName != *input.UserDisplayName) {
		return false
	}

	if p.UserPrincipalName != nil && (input.UserPrincipalName == nil || *p.UserPrincipalName != *input.UserPrincipalName) {
		return false
	}

	return true
}
