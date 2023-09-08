package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedEmailDomainInvitation struct {
	ExpiryTime       *string `json:"expiryTime,omitempty"`
	Id               *string `json:"id,omitempty"`
	InvitationDomain *string `json:"invitationDomain,omitempty"`
	InvitationStatus *string `json:"invitationStatus,omitempty"`
	ODataType        *string `json:"@odata.type,omitempty"`
}
