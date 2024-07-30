package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryCustodian struct {
	AcknowledgedDateTime *string                                `json:"acknowledgedDateTime,omitempty"`
	CreatedDateTime      *string                                `json:"createdDateTime,omitempty"`
	DisplayName          *string                                `json:"displayName,omitempty"`
	Email                *string                                `json:"email,omitempty"`
	HoldStatus           *SecurityEdiscoveryCustodianHoldStatus `json:"holdStatus,omitempty"`
	Id                   *string                                `json:"id,omitempty"`
	LastIndexOperation   *SecurityEdiscoveryIndexOperation      `json:"lastIndexOperation,omitempty"`
	LastModifiedDateTime *string                                `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                                `json:"@odata.type,omitempty"`
	ReleasedDateTime     *string                                `json:"releasedDateTime,omitempty"`
	SiteSources          *[]SecuritySiteSource                  `json:"siteSources,omitempty"`
	Status               *SecurityEdiscoveryCustodianStatus     `json:"status,omitempty"`
	UnifiedGroupSources  *[]SecurityUnifiedGroupSource          `json:"unifiedGroupSources,omitempty"`
	UserSources          *[]SecurityUserSource                  `json:"userSources,omitempty"`
}
