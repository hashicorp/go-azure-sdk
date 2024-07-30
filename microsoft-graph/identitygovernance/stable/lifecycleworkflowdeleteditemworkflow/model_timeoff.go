package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeOff struct {
	CreatedDateTime      *string      `json:"createdDateTime,omitempty"`
	DraftTimeOff         *TimeOffItem `json:"draftTimeOff,omitempty"`
	Id                   *string      `json:"id,omitempty"`
	LastModifiedBy       *IdentitySet `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string      `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string      `json:"@odata.type,omitempty"`
	SharedTimeOff        *TimeOffItem `json:"sharedTimeOff,omitempty"`
	UserId               *string      `json:"userId,omitempty"`
}
