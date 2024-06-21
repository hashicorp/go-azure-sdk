package roleeligibilityschedulerequests

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleEligibilityScheduleRequestPropertiesScheduleInfo struct {
	Expiration    *RoleEligibilityScheduleRequestPropertiesScheduleInfoExpiration `json:"expiration,omitempty"`
	StartDateTime *string                                                         `json:"startDateTime,omitempty"`
}
