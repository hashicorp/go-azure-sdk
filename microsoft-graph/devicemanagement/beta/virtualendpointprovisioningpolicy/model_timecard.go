package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeCard struct {
	Breaks               *[]TimeCardBreak     `json:"breaks,omitempty"`
	ClockInEvent         *TimeCardEvent       `json:"clockInEvent,omitempty"`
	ClockOutEvent        *TimeCardEvent       `json:"clockOutEvent,omitempty"`
	ConfirmedBy          *TimeCardConfirmedBy `json:"confirmedBy,omitempty"`
	CreatedBy            *IdentitySet         `json:"createdBy,omitempty"`
	CreatedDateTime      *string              `json:"createdDateTime,omitempty"`
	Id                   *string              `json:"id,omitempty"`
	LastModifiedBy       *IdentitySet         `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string              `json:"lastModifiedDateTime,omitempty"`
	Notes                *ItemBody            `json:"notes,omitempty"`
	ODataType            *string              `json:"@odata.type,omitempty"`
	OriginalEntry        *TimeCardEntry       `json:"originalEntry,omitempty"`
	State                *TimeCardState       `json:"state,omitempty"`
	UserId               *string              `json:"userId,omitempty"`
}
