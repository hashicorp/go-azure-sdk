package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AvailabilityItem struct {
	EndDateTime   *DateTimeTimeZone       `json:"endDateTime,omitempty"`
	ODataType     *string                 `json:"@odata.type,omitempty"`
	ServiceId     *string                 `json:"serviceId,omitempty"`
	StartDateTime *DateTimeTimeZone       `json:"startDateTime,omitempty"`
	Status        *AvailabilityItemStatus `json:"status,omitempty"`
}
