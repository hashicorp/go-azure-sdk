package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharingInvitation struct {
	Email          *string      `json:"email,omitempty"`
	InvitedBy      *IdentitySet `json:"invitedBy,omitempty"`
	ODataType      *string      `json:"@odata.type,omitempty"`
	RedeemedBy     *string      `json:"redeemedBy,omitempty"`
	SignInRequired *bool        `json:"signInRequired,omitempty"`
}
