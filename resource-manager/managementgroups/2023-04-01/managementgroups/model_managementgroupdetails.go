package managementgroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagementGroupDetails struct {
	ManagementGroupAncestors      *[]string                     `json:"managementGroupAncestors,omitempty"`
	ManagementGroupAncestorsChain *[]ManagementGroupPathElement `json:"managementGroupAncestorsChain,omitempty"`
	Parent                        *ParentGroupInfo              `json:"parent,omitempty"`
	Path                          *[]ManagementGroupPathElement `json:"path,omitempty"`
	UpdatedBy                     *string                       `json:"updatedBy,omitempty"`
	UpdatedTime                   *string                       `json:"updatedTime,omitempty"`
	Version                       *float64                      `json:"version,omitempty"`
}
