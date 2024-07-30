package teamscheduletimecard

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateTeamScheduleTimeCardClockInRequest struct {
	AtApprovedLocation *bool     `json:"atApprovedLocation,omitempty"`
	Notes              *ItemBody `json:"notes,omitempty"`
	OnBehalfOfUserId   *string   `json:"onBehalfOfUserId,omitempty"`
}
