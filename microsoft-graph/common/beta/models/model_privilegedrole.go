package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedRole struct {
	Assignments *[]PrivilegedRoleAssignment `json:"assignments,omitempty"`
	Id          *string                     `json:"id,omitempty"`
	Name        *string                     `json:"name,omitempty"`
	ODataType   *string                     `json:"@odata.type,omitempty"`
	Settings    *PrivilegedRoleSettings     `json:"settings,omitempty"`
	Summary     *PrivilegedRoleSummary      `json:"summary,omitempty"`
}
