package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsPstnCallLogRowOperationPredicate struct {
	CallId             *string
	CallType           *string
	CalleeNumber       *string
	CallerNumber       *string
	Charge             *float64
	ConferenceId       *string
	ConnectionCharge   *float64
	Currency           *string
	DestinationContext *string
	DestinationName    *string
	Duration           *int64
	EndDateTime        *string
	Id                 *string
	InventoryType      *string
	LicenseCapability  *string
	ODataType          *string
	Operator           *string
	StartDateTime      *string
	TenantCountryCode  *string
	UsageCountryCode   *string
	UserDisplayName    *string
	UserId             *string
	UserPrincipalName  *string
}

func (p CallRecordsPstnCallLogRowOperationPredicate) Matches(input CallRecordsPstnCallLogRow) bool {

	if p.CallId != nil && (input.CallId == nil || *p.CallId != *input.CallId) {
		return false
	}

	if p.CallType != nil && (input.CallType == nil || *p.CallType != *input.CallType) {
		return false
	}

	if p.CalleeNumber != nil && (input.CalleeNumber == nil || *p.CalleeNumber != *input.CalleeNumber) {
		return false
	}

	if p.CallerNumber != nil && (input.CallerNumber == nil || *p.CallerNumber != *input.CallerNumber) {
		return false
	}

	if p.Charge != nil && (input.Charge == nil || *p.Charge != *input.Charge) {
		return false
	}

	if p.ConferenceId != nil && (input.ConferenceId == nil || *p.ConferenceId != *input.ConferenceId) {
		return false
	}

	if p.ConnectionCharge != nil && (input.ConnectionCharge == nil || *p.ConnectionCharge != *input.ConnectionCharge) {
		return false
	}

	if p.Currency != nil && (input.Currency == nil || *p.Currency != *input.Currency) {
		return false
	}

	if p.DestinationContext != nil && (input.DestinationContext == nil || *p.DestinationContext != *input.DestinationContext) {
		return false
	}

	if p.DestinationName != nil && (input.DestinationName == nil || *p.DestinationName != *input.DestinationName) {
		return false
	}

	if p.Duration != nil && (input.Duration == nil || *p.Duration != *input.Duration) {
		return false
	}

	if p.EndDateTime != nil && (input.EndDateTime == nil || *p.EndDateTime != *input.EndDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.InventoryType != nil && (input.InventoryType == nil || *p.InventoryType != *input.InventoryType) {
		return false
	}

	if p.LicenseCapability != nil && (input.LicenseCapability == nil || *p.LicenseCapability != *input.LicenseCapability) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Operator != nil && (input.Operator == nil || *p.Operator != *input.Operator) {
		return false
	}

	if p.StartDateTime != nil && (input.StartDateTime == nil || *p.StartDateTime != *input.StartDateTime) {
		return false
	}

	if p.TenantCountryCode != nil && (input.TenantCountryCode == nil || *p.TenantCountryCode != *input.TenantCountryCode) {
		return false
	}

	if p.UsageCountryCode != nil && (input.UsageCountryCode == nil || *p.UsageCountryCode != *input.UsageCountryCode) {
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
