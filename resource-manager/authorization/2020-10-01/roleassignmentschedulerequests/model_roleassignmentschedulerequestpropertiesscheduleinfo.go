package roleassignmentschedulerequests

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleAssignmentScheduleRequestPropertiesScheduleInfo struct {
	Expiration    *RoleAssignmentScheduleRequestPropertiesScheduleInfoExpiration `json:"expiration,omitempty"`
	StartDateTime *string                                                        `json:"startDateTime,omitempty"`
}
