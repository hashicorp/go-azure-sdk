package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerAssignment struct {
	AssignedBy       *IdentitySet `json:"assignedBy,omitempty"`
	AssignedDateTime *string      `json:"assignedDateTime,omitempty"`
	ODataType        *string      `json:"@odata.type,omitempty"`
	OrderHint        *string      `json:"orderHint,omitempty"`
}
