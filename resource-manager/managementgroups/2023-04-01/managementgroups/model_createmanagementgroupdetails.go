package managementgroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateManagementGroupDetails struct {
	Parent      *CreateParentGroupInfo `json:"parent,omitempty"`
	UpdatedBy   *string                `json:"updatedBy,omitempty"`
	UpdatedTime *string                `json:"updatedTime,omitempty"`
	Version     *float64               `json:"version,omitempty"`
}
