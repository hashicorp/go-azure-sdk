package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventWebinar struct {
	Audience                  *VirtualEventWebinarAudience           `json:"audience,omitempty"`
	CoOrganizers              *[]CommunicationsUserIdentity          `json:"coOrganizers,omitempty"`
	CreatedBy                 *CommunicationsIdentitySet             `json:"createdBy,omitempty"`
	Description               *string                                `json:"description,omitempty"`
	DisplayName               *string                                `json:"displayName,omitempty"`
	EndDateTime               *DateTimeTimeZone                      `json:"endDateTime,omitempty"`
	Id                        *string                                `json:"id,omitempty"`
	ODataType                 *string                                `json:"@odata.type,omitempty"`
	Presenters                *[]VirtualEventPresenter               `json:"presenters,omitempty"`
	RegistrationConfiguration *VirtualEventRegistrationConfiguration `json:"registrationConfiguration,omitempty"`
	Registrations             *[]VirtualEventRegistration            `json:"registrations,omitempty"`
	Sessions                  *[]VirtualEventSession                 `json:"sessions,omitempty"`
	StartDateTime             *DateTimeTimeZone                      `json:"startDateTime,omitempty"`
	Status                    *VirtualEventWebinarStatus             `json:"status,omitempty"`
}
