package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttendeeAvailability struct {
	Attendee     *AttendeeBase                     `json:"attendee,omitempty"`
	Availability *AttendeeAvailabilityAvailability `json:"availability,omitempty"`
	ODataType    *string                           `json:"@odata.type,omitempty"`
}
