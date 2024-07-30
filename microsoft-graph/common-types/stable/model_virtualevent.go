package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEvent struct {
	CreatedBy     *CommunicationsIdentitySet `json:"createdBy,omitempty"`
	Description   *ItemBody                  `json:"description,omitempty"`
	DisplayName   *string                    `json:"displayName,omitempty"`
	EndDateTime   *DateTimeTimeZone          `json:"endDateTime,omitempty"`
	Id            *string                    `json:"id,omitempty"`
	ODataType     *string                    `json:"@odata.type,omitempty"`
	Sessions      *[]VirtualEventSession     `json:"sessions,omitempty"`
	StartDateTime *DateTimeTimeZone          `json:"startDateTime,omitempty"`
	Status        *VirtualEventStatus        `json:"status,omitempty"`
}
