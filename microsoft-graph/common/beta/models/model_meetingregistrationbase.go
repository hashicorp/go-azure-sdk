package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingRegistrationBase struct {
	AllowedRegistrant *MeetingRegistrationBaseAllowedRegistrant `json:"allowedRegistrant,omitempty"`
	Id                *string                                   `json:"id,omitempty"`
	ODataType         *string                                   `json:"@odata.type,omitempty"`
	Registrants       *[]MeetingRegistrantBase                  `json:"registrants,omitempty"`
}
