package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryUnifiedGroupSource struct {
	CreatedBy       *IdentitySet                                 `json:"createdBy,omitempty"`
	CreatedDateTime *string                                      `json:"createdDateTime,omitempty"`
	DisplayName     *string                                      `json:"displayName,omitempty"`
	Group           *Group                                       `json:"group,omitempty"`
	HoldStatus      *EdiscoveryUnifiedGroupSourceHoldStatus      `json:"holdStatus,omitempty"`
	Id              *string                                      `json:"id,omitempty"`
	IncludedSources *EdiscoveryUnifiedGroupSourceIncludedSources `json:"includedSources,omitempty"`
	ODataType       *string                                      `json:"@odata.type,omitempty"`
}
