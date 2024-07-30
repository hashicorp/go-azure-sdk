package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ShiftActivity struct {
	Code          *string             `json:"code,omitempty"`
	DisplayName   *string             `json:"displayName,omitempty"`
	EndDateTime   *string             `json:"endDateTime,omitempty"`
	IsPaid        *bool               `json:"isPaid,omitempty"`
	ODataType     *string             `json:"@odata.type,omitempty"`
	StartDateTime *string             `json:"startDateTime,omitempty"`
	Theme         *ShiftActivityTheme `json:"theme,omitempty"`
}
