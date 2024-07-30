package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerRelationshipBasedUserType struct {
	ODataType *string                                   `json:"@odata.type,omitempty"`
	Role      *PlannerRelationshipBasedUserTypeRole     `json:"role,omitempty"`
	RoleKind  *PlannerRelationshipBasedUserTypeRoleKind `json:"roleKind,omitempty"`
}
