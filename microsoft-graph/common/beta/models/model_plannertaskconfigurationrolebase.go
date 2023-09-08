package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskConfigurationRoleBase struct {
	ODataType *string                                   `json:"@odata.type,omitempty"`
	RoleKind  *PlannerTaskConfigurationRoleBaseRoleKind `json:"roleKind,omitempty"`
}
