package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroup struct {
	AssignmentApprovals          *[]Approval                                         `json:"assignmentApprovals,omitempty"`
	AssignmentScheduleInstances  *[]PrivilegedAccessGroupAssignmentScheduleInstance  `json:"assignmentScheduleInstances,omitempty"`
	AssignmentScheduleRequests   *[]PrivilegedAccessGroupAssignmentScheduleRequest   `json:"assignmentScheduleRequests,omitempty"`
	AssignmentSchedules          *[]PrivilegedAccessGroupAssignmentSchedule          `json:"assignmentSchedules,omitempty"`
	EligibilityScheduleInstances *[]PrivilegedAccessGroupEligibilityScheduleInstance `json:"eligibilityScheduleInstances,omitempty"`
	EligibilityScheduleRequests  *[]PrivilegedAccessGroupEligibilityScheduleRequest  `json:"eligibilityScheduleRequests,omitempty"`
	EligibilitySchedules         *[]PrivilegedAccessGroupEligibilitySchedule         `json:"eligibilitySchedules,omitempty"`
	Id                           *string                                             `json:"id,omitempty"`
	ODataType                    *string                                             `json:"@odata.type,omitempty"`
}
