package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Invitation struct {
	Id                      *string                 `json:"id,omitempty"`
	InviteRedeemUrl         *string                 `json:"inviteRedeemUrl,omitempty"`
	InviteRedirectUrl       *string                 `json:"inviteRedirectUrl,omitempty"`
	InvitedUser             *User                   `json:"invitedUser,omitempty"`
	InvitedUserDisplayName  *string                 `json:"invitedUserDisplayName,omitempty"`
	InvitedUserEmailAddress *string                 `json:"invitedUserEmailAddress,omitempty"`
	InvitedUserMessageInfo  *InvitedUserMessageInfo `json:"invitedUserMessageInfo,omitempty"`
	InvitedUserType         *string                 `json:"invitedUserType,omitempty"`
	ODataType               *string                 `json:"@odata.type,omitempty"`
	ResetRedemption         *bool                   `json:"resetRedemption,omitempty"`
	SendInvitationMessage   *bool                   `json:"sendInvitationMessage,omitempty"`
	Status                  *string                 `json:"status,omitempty"`
}
