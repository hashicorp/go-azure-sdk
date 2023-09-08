package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvitationOperationPredicate struct {
	Id                      *string
	InviteRedeemUrl         *string
	InviteRedirectUrl       *string
	InvitedUserDisplayName  *string
	InvitedUserEmailAddress *string
	InvitedUserType         *string
	ODataType               *string
	ResetRedemption         *bool
	SendInvitationMessage   *bool
	Status                  *string
}

func (p InvitationOperationPredicate) Matches(input Invitation) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.InviteRedeemUrl != nil && (input.InviteRedeemUrl == nil || *p.InviteRedeemUrl != *input.InviteRedeemUrl) {
		return false
	}

	if p.InviteRedirectUrl != nil && (input.InviteRedirectUrl == nil || *p.InviteRedirectUrl != *input.InviteRedirectUrl) {
		return false
	}

	if p.InvitedUserDisplayName != nil && (input.InvitedUserDisplayName == nil || *p.InvitedUserDisplayName != *input.InvitedUserDisplayName) {
		return false
	}

	if p.InvitedUserEmailAddress != nil && (input.InvitedUserEmailAddress == nil || *p.InvitedUserEmailAddress != *input.InvitedUserEmailAddress) {
		return false
	}

	if p.InvitedUserType != nil && (input.InvitedUserType == nil || *p.InvitedUserType != *input.InvitedUserType) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ResetRedemption != nil && (input.ResetRedemption == nil || *p.ResetRedemption != *input.ResetRedemption) {
		return false
	}

	if p.SendInvitationMessage != nil && (input.SendInvitationMessage == nil || *p.SendInvitationMessage != *input.SendInvitationMessage) {
		return false
	}

	if p.Status != nil && (input.Status == nil || *p.Status != *input.Status) {
		return false
	}

	return true
}
