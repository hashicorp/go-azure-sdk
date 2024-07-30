package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUnifiedGroupSource struct {
	CreatedBy       *IdentitySet                               `json:"createdBy,omitempty"`
	CreatedDateTime *string                                    `json:"createdDateTime,omitempty"`
	DisplayName     *string                                    `json:"displayName,omitempty"`
	Group           *Group                                     `json:"group,omitempty"`
	HoldStatus      *SecurityUnifiedGroupSourceHoldStatus      `json:"holdStatus,omitempty"`
	Id              *string                                    `json:"id,omitempty"`
	IncludedSources *SecurityUnifiedGroupSourceIncludedSources `json:"includedSources,omitempty"`
	ODataType       *string                                    `json:"@odata.type,omitempty"`
}
