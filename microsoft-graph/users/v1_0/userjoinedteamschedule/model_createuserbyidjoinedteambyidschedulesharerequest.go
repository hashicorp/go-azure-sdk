package userjoinedteamschedule

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserByIdJoinedTeamByIdScheduleShareRequest struct {
	EndDateTime   *string `json:"endDateTime,omitempty"`
	NotifyTeam    *bool   `json:"notifyTeam,omitempty"`
	StartDateTime *string `json:"startDateTime,omitempty"`
}
