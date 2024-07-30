package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUserSource struct {
	CreatedBy       *IdentitySet                       `json:"createdBy,omitempty"`
	CreatedDateTime *string                            `json:"createdDateTime,omitempty"`
	DisplayName     *string                            `json:"displayName,omitempty"`
	Email           *string                            `json:"email,omitempty"`
	HoldStatus      *SecurityUserSourceHoldStatus      `json:"holdStatus,omitempty"`
	Id              *string                            `json:"id,omitempty"`
	IncludedSources *SecurityUserSourceIncludedSources `json:"includedSources,omitempty"`
	ODataType       *string                            `json:"@odata.type,omitempty"`
	SiteWebUrl      *string                            `json:"siteWebUrl,omitempty"`
}
