package teamscheduletimecard

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateTeamScheduleTimeCardClockInRequest struct {
	AtApprovedLocation   nullable.Type[bool]   `json:"atApprovedLocation,omitempty"`
	IsAtApprovedLocation nullable.Type[bool]   `json:"isAtApprovedLocation,omitempty"`
	Notes                *beta.ItemBody        `json:"notes,omitempty"`
	OnBehalfOfUserId     nullable.Type[string] `json:"onBehalfOfUserId,omitempty"`
}
