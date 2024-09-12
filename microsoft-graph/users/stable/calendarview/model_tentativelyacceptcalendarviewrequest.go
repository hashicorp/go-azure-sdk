package calendarview

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TentativelyAcceptCalendarViewRequest struct {
	Comment         nullable.Type[string] `json:"Comment,omitempty"`
	ProposedNewTime *stable.TimeSlot      `json:"ProposedNewTime,omitempty"`
	SendResponse    nullable.Type[bool]   `json:"SendResponse,omitempty"`
}
