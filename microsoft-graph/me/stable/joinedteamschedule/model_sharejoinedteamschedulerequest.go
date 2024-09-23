package joinedteamschedule

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ShareJoinedTeamScheduleRequest struct {
	EndDateTime   nullable.Type[string] `json:"endDateTime,omitempty"`
	NotifyTeam    nullable.Type[bool]   `json:"notifyTeam,omitempty"`
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`
}
