package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsPstnOnlineMeetingDialoutReport struct {
	Currency           *string  `json:"currency,omitempty"`
	DestinationContext *string  `json:"destinationContext,omitempty"`
	ODataType          *string  `json:"@odata.type,omitempty"`
	TotalCallCharge    *float64 `json:"totalCallCharge,omitempty"`
	TotalCallSeconds   *int64   `json:"totalCallSeconds,omitempty"`
	TotalCalls         *int64   `json:"totalCalls,omitempty"`
	UsageLocation      *string  `json:"usageLocation,omitempty"`
	UserDisplayName    *string  `json:"userDisplayName,omitempty"`
	UserId             *string  `json:"userId,omitempty"`
	UserPrincipalName  *string  `json:"userPrincipalName,omitempty"`
}
