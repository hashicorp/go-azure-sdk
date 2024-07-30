package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicySet struct {
	Assignments          *[]PolicySetAssignment `json:"assignments,omitempty"`
	CreatedDateTime      *string                `json:"createdDateTime,omitempty"`
	Description          *string                `json:"description,omitempty"`
	DisplayName          *string                `json:"displayName,omitempty"`
	ErrorCode            *PolicySetErrorCode    `json:"errorCode,omitempty"`
	GuidedDeploymentTags *[]string              `json:"guidedDeploymentTags,omitempty"`
	Id                   *string                `json:"id,omitempty"`
	Items                *[]PolicySetItem       `json:"items,omitempty"`
	LastModifiedDateTime *string                `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                `json:"@odata.type,omitempty"`
	RoleScopeTags        *[]string              `json:"roleScopeTags,omitempty"`
	Status               *PolicySetStatus       `json:"status,omitempty"`
}
