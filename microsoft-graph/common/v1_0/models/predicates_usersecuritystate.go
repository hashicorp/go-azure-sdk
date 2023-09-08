package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSecurityStateOperationPredicate struct {
	AadUserId                    *string
	AccountName                  *string
	DomainName                   *string
	IsVpn                        *bool
	LogonDateTime                *string
	LogonId                      *string
	LogonIp                      *string
	LogonLocation                *string
	ODataType                    *string
	OnPremisesSecurityIdentifier *string
	RiskScore                    *string
	UserPrincipalName            *string
}

func (p UserSecurityStateOperationPredicate) Matches(input UserSecurityState) bool {

	if p.AadUserId != nil && (input.AadUserId == nil || *p.AadUserId != *input.AadUserId) {
		return false
	}

	if p.AccountName != nil && (input.AccountName == nil || *p.AccountName != *input.AccountName) {
		return false
	}

	if p.DomainName != nil && (input.DomainName == nil || *p.DomainName != *input.DomainName) {
		return false
	}

	if p.IsVpn != nil && (input.IsVpn == nil || *p.IsVpn != *input.IsVpn) {
		return false
	}

	if p.LogonDateTime != nil && (input.LogonDateTime == nil || *p.LogonDateTime != *input.LogonDateTime) {
		return false
	}

	if p.LogonId != nil && (input.LogonId == nil || *p.LogonId != *input.LogonId) {
		return false
	}

	if p.LogonIp != nil && (input.LogonIp == nil || *p.LogonIp != *input.LogonIp) {
		return false
	}

	if p.LogonLocation != nil && (input.LogonLocation == nil || *p.LogonLocation != *input.LogonLocation) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OnPremisesSecurityIdentifier != nil && (input.OnPremisesSecurityIdentifier == nil || *p.OnPremisesSecurityIdentifier != *input.OnPremisesSecurityIdentifier) {
		return false
	}

	if p.RiskScore != nil && (input.RiskScore == nil || *p.RiskScore != *input.RiskScore) {
		return false
	}

	if p.UserPrincipalName != nil && (input.UserPrincipalName == nil || *p.UserPrincipalName != *input.UserPrincipalName) {
		return false
	}

	return true
}
