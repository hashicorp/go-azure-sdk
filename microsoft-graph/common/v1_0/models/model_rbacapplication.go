package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RbacApplication struct {
	Id                               *string                                   `json:"id,omitempty"`
	ODataType                        *string                                   `json:"@odata.type,omitempty"`
	ResourceNamespaces               *[]UnifiedRbacResourceNamespace           `json:"resourceNamespaces,omitempty"`
	RoleAssignmentScheduleInstances  *[]UnifiedRoleAssignmentScheduleInstance  `json:"roleAssignmentScheduleInstances,omitempty"`
	RoleAssignmentScheduleRequests   *[]UnifiedRoleAssignmentScheduleRequest   `json:"roleAssignmentScheduleRequests,omitempty"`
	RoleAssignmentSchedules          *[]UnifiedRoleAssignmentSchedule          `json:"roleAssignmentSchedules,omitempty"`
	RoleAssignments                  *[]UnifiedRoleAssignment                  `json:"roleAssignments,omitempty"`
	RoleDefinitions                  *[]UnifiedRoleDefinition                  `json:"roleDefinitions,omitempty"`
	RoleEligibilityScheduleInstances *[]UnifiedRoleEligibilityScheduleInstance `json:"roleEligibilityScheduleInstances,omitempty"`
	RoleEligibilityScheduleRequests  *[]UnifiedRoleEligibilityScheduleRequest  `json:"roleEligibilityScheduleRequests,omitempty"`
	RoleEligibilitySchedules         *[]UnifiedRoleEligibilitySchedule         `json:"roleEligibilitySchedules,omitempty"`
}
