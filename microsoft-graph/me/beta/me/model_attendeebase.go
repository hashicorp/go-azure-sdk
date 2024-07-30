package me

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttendeeBase struct {
	EmailAddress *EmailAddress     `json:"emailAddress,omitempty"`
	ODataType    *string           `json:"@odata.type,omitempty"`
	Type         *AttendeeBaseType `json:"type,omitempty"`
}
