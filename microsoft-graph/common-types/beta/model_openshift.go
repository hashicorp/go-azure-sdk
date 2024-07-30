package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenShift struct {
	CreatedBy            *IdentitySet   `json:"createdBy,omitempty"`
	CreatedDateTime      *string        `json:"createdDateTime,omitempty"`
	DraftOpenShift       *OpenShiftItem `json:"draftOpenShift,omitempty"`
	Id                   *string        `json:"id,omitempty"`
	IsStagedForDeletion  *bool          `json:"isStagedForDeletion,omitempty"`
	LastModifiedBy       *IdentitySet   `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string        `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string        `json:"@odata.type,omitempty"`
	SchedulingGroupId    *string        `json:"schedulingGroupId,omitempty"`
	SchedulingGroupName  *string        `json:"schedulingGroupName,omitempty"`
	SharedOpenShift      *OpenShiftItem `json:"sharedOpenShift,omitempty"`
	TeamId               *string        `json:"teamId,omitempty"`
	TeamName             *string        `json:"teamName,omitempty"`
}
