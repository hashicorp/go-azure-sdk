package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityInfoProperties struct {
	DisplayName            *string                `json:"displayName,omitempty"`
	InheritedPermissions   *InheritedPermissions  `json:"inheritedPermissions,omitempty"`
	NumberOfChildGroups    *int64                 `json:"numberOfChildGroups,omitempty"`
	NumberOfChildren       *int64                 `json:"numberOfChildren,omitempty"`
	NumberOfDescendants    *int64                 `json:"numberOfDescendants,omitempty"`
	Parent                 *EntityParentGroupInfo `json:"parent,omitempty"`
	ParentDisplayNameChain *[]string              `json:"parentDisplayNameChain,omitempty"`
	ParentNameChain        *[]string              `json:"parentNameChain,omitempty"`
	Permissions            *Permissions           `json:"permissions,omitempty"`
	TenantId               *string                `json:"tenantId,omitempty"`
}
