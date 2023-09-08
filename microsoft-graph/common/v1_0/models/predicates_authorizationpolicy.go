package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationPolicyOperationPredicate struct {
	AllowEmailVerifiedUsersToJoinOrganization *bool
	AllowUserConsentForRiskyApps              *bool
	AllowedToSignUpEmailBasedSubscriptions    *bool
	AllowedToUseSSPR                          *bool
	BlockMsolPowerShell                       *bool
	DeletedDateTime                           *string
	Description                               *string
	DisplayName                               *string
	GuestUserRoleId                           *string
	Id                                        *string
	ODataType                                 *string
}

func (p AuthorizationPolicyOperationPredicate) Matches(input AuthorizationPolicy) bool {

	if p.AllowEmailVerifiedUsersToJoinOrganization != nil && (input.AllowEmailVerifiedUsersToJoinOrganization == nil || *p.AllowEmailVerifiedUsersToJoinOrganization != *input.AllowEmailVerifiedUsersToJoinOrganization) {
		return false
	}

	if p.AllowUserConsentForRiskyApps != nil && (input.AllowUserConsentForRiskyApps == nil || *p.AllowUserConsentForRiskyApps != *input.AllowUserConsentForRiskyApps) {
		return false
	}

	if p.AllowedToSignUpEmailBasedSubscriptions != nil && (input.AllowedToSignUpEmailBasedSubscriptions == nil || *p.AllowedToSignUpEmailBasedSubscriptions != *input.AllowedToSignUpEmailBasedSubscriptions) {
		return false
	}

	if p.AllowedToUseSSPR != nil && (input.AllowedToUseSSPR == nil || *p.AllowedToUseSSPR != *input.AllowedToUseSSPR) {
		return false
	}

	if p.BlockMsolPowerShell != nil && (input.BlockMsolPowerShell == nil || *p.BlockMsolPowerShell != *input.BlockMsolPowerShell) {
		return false
	}

	if p.DeletedDateTime != nil && (input.DeletedDateTime == nil || *p.DeletedDateTime != *input.DeletedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.GuestUserRoleId != nil && (input.GuestUserRoleId == nil || *p.GuestUserRoleId != *input.GuestUserRoleId) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
