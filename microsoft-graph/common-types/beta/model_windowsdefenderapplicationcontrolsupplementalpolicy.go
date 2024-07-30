package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDefenderApplicationControlSupplementalPolicy struct {
	Assignments          *[]WindowsDefenderApplicationControlSupplementalPolicyAssignment       `json:"assignments,omitempty"`
	Content              *string                                                                `json:"content,omitempty"`
	ContentFileName      *string                                                                `json:"contentFileName,omitempty"`
	CreationDateTime     *string                                                                `json:"creationDateTime,omitempty"`
	DeploySummary        *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary  `json:"deploySummary,omitempty"`
	Description          *string                                                                `json:"description,omitempty"`
	DeviceStatuses       *[]WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus `json:"deviceStatuses,omitempty"`
	DisplayName          *string                                                                `json:"displayName,omitempty"`
	Id                   *string                                                                `json:"id,omitempty"`
	LastModifiedDateTime *string                                                                `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                                                                `json:"@odata.type,omitempty"`
	RoleScopeTagIds      *[]string                                                              `json:"roleScopeTagIds,omitempty"`
	Version              *string                                                                `json:"version,omitempty"`
}
