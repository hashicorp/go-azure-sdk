package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobilityManagementPolicy struct {
	AppliesTo      *MobilityManagementPolicyAppliesTo `json:"appliesTo,omitempty"`
	ComplianceUrl  *string                            `json:"complianceUrl,omitempty"`
	Description    *string                            `json:"description,omitempty"`
	DiscoveryUrl   *string                            `json:"discoveryUrl,omitempty"`
	DisplayName    *string                            `json:"displayName,omitempty"`
	Id             *string                            `json:"id,omitempty"`
	IncludedGroups *[]Group                           `json:"includedGroups,omitempty"`
	IsValid        *bool                              `json:"isValid,omitempty"`
	ODataType      *string                            `json:"@odata.type,omitempty"`
	TermsOfUseUrl  *string                            `json:"termsOfUseUrl,omitempty"`
}
