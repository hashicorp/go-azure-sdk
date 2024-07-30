package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoverySiteSource struct {
	CreatedBy       *IdentitySet                    `json:"createdBy,omitempty"`
	CreatedDateTime *string                         `json:"createdDateTime,omitempty"`
	DisplayName     *string                         `json:"displayName,omitempty"`
	HoldStatus      *EdiscoverySiteSourceHoldStatus `json:"holdStatus,omitempty"`
	Id              *string                         `json:"id,omitempty"`
	ODataType       *string                         `json:"@odata.type,omitempty"`
	Site            *Site                           `json:"site,omitempty"`
}
