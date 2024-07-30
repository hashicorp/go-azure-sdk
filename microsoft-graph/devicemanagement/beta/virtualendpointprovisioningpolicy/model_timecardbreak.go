package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeCardBreak struct {
	BreakId   *string        `json:"breakId,omitempty"`
	End       *TimeCardEvent `json:"end,omitempty"`
	Notes     *ItemBody      `json:"notes,omitempty"`
	ODataType *string        `json:"@odata.type,omitempty"`
	Start     *TimeCardEvent `json:"start,omitempty"`
}
