package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenShift struct {
	CreatedDateTime      *string        `json:"createdDateTime,omitempty"`
	DraftOpenShift       *OpenShiftItem `json:"draftOpenShift,omitempty"`
	Id                   *string        `json:"id,omitempty"`
	LastModifiedBy       *IdentitySet   `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string        `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string        `json:"@odata.type,omitempty"`
	SchedulingGroupId    *string        `json:"schedulingGroupId,omitempty"`
	SharedOpenShift      *OpenShiftItem `json:"sharedOpenShift,omitempty"`
}
