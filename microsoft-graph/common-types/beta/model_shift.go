package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Shift struct {
	CreatedBy            *IdentitySet `json:"createdBy,omitempty"`
	CreatedDateTime      *string      `json:"createdDateTime,omitempty"`
	DraftShift           *ShiftItem   `json:"draftShift,omitempty"`
	Id                   *string      `json:"id,omitempty"`
	IsStagedForDeletion  *bool        `json:"isStagedForDeletion,omitempty"`
	LastModifiedBy       *IdentitySet `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string      `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string      `json:"@odata.type,omitempty"`
	SchedulingGroupId    *string      `json:"schedulingGroupId,omitempty"`
	SharedShift          *ShiftItem   `json:"sharedShift,omitempty"`
	UserId               *string      `json:"userId,omitempty"`
}
