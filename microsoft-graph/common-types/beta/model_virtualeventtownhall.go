package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventTownhall struct {
	Audience         *VirtualEventTownhallAudience `json:"audience,omitempty"`
	CoOrganizers     *[]CommunicationsUserIdentity `json:"coOrganizers,omitempty"`
	CreatedBy        *CommunicationsIdentitySet    `json:"createdBy,omitempty"`
	Description      *ItemBody                     `json:"description,omitempty"`
	DisplayName      *string                       `json:"displayName,omitempty"`
	EndDateTime      *DateTimeTimeZone             `json:"endDateTime,omitempty"`
	Id               *string                       `json:"id,omitempty"`
	InvitedAttendees *[]CommunicationsUserIdentity `json:"invitedAttendees,omitempty"`
	IsInviteOnly     *bool                         `json:"isInviteOnly,omitempty"`
	ODataType        *string                       `json:"@odata.type,omitempty"`
	Presenters       *[]VirtualEventPresenter      `json:"presenters,omitempty"`
	Sessions         *[]VirtualEventSession        `json:"sessions,omitempty"`
	StartDateTime    *DateTimeTimeZone             `json:"startDateTime,omitempty"`
	Status           *VirtualEventTownhallStatus   `json:"status,omitempty"`
}
