package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VppToken struct {
	AppleId                             *string                      `json:"appleId,omitempty"`
	AutomaticallyUpdateApps             *bool                        `json:"automaticallyUpdateApps,omitempty"`
	ClaimTokenManagementFromExternalMdm *bool                        `json:"claimTokenManagementFromExternalMdm,omitempty"`
	CountryOrRegion                     *string                      `json:"countryOrRegion,omitempty"`
	DataSharingConsentGranted           *bool                        `json:"dataSharingConsentGranted,omitempty"`
	DisplayName                         *string                      `json:"displayName,omitempty"`
	ExpirationDateTime                  *string                      `json:"expirationDateTime,omitempty"`
	Id                                  *string                      `json:"id,omitempty"`
	LastModifiedDateTime                *string                      `json:"lastModifiedDateTime,omitempty"`
	LastSyncDateTime                    *string                      `json:"lastSyncDateTime,omitempty"`
	LastSyncStatus                      *VppTokenLastSyncStatus      `json:"lastSyncStatus,omitempty"`
	LocationName                        *string                      `json:"locationName,omitempty"`
	ODataType                           *string                      `json:"@odata.type,omitempty"`
	OrganizationName                    *string                      `json:"organizationName,omitempty"`
	RoleScopeTagIds                     *[]string                    `json:"roleScopeTagIds,omitempty"`
	State                               *VppTokenState               `json:"state,omitempty"`
	Token                               *string                      `json:"token,omitempty"`
	TokenActionResults                  *[]VppTokenActionResult      `json:"tokenActionResults,omitempty"`
	VppTokenAccountType                 *VppTokenVppTokenAccountType `json:"vppTokenAccountType,omitempty"`
}
