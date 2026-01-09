package calendarevent

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TentativelyAcceptCalendarEventRequest struct {
	Comment         nullable.Type[string] `json:"Comment,omitempty"`
	ProposedNewTime *beta.TimeSlot        `json:"ProposedNewTime,omitempty"`
	SendResponse    nullable.Type[bool]   `json:"SendResponse,omitempty"`
}
