package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeCardEvent struct {
	AtApprovedLocation *bool     `json:"atApprovedLocation,omitempty"`
	DateTime           *string   `json:"dateTime,omitempty"`
	Notes              *ItemBody `json:"notes,omitempty"`
	ODataType          *string   `json:"@odata.type,omitempty"`
}
