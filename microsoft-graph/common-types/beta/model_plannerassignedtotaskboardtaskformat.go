package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerAssignedToTaskBoardTaskFormat struct {
	Id                   *string                      `json:"id,omitempty"`
	ODataType            *string                      `json:"@odata.type,omitempty"`
	OrderHintsByAssignee *PlannerOrderHintsByAssignee `json:"orderHintsByAssignee,omitempty"`
	UnassignedOrderHint  *string                      `json:"unassignedOrderHint,omitempty"`
}
