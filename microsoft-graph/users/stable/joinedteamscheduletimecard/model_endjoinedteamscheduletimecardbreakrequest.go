package joinedteamscheduletimecard

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndJoinedTeamScheduleTimeCardBreakRequest struct {
	IsAtApprovedLocation nullable.Type[bool] `json:"isAtApprovedLocation,omitempty"`
	Notes                *stable.ItemBody    `json:"notes,omitempty"`
}
