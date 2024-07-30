package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeOffItem struct {
	EndDateTime     *string           `json:"endDateTime,omitempty"`
	ODataType       *string           `json:"@odata.type,omitempty"`
	StartDateTime   *string           `json:"startDateTime,omitempty"`
	Theme           *TimeOffItemTheme `json:"theme,omitempty"`
	TimeOffReasonId *string           `json:"timeOffReasonId,omitempty"`
}
