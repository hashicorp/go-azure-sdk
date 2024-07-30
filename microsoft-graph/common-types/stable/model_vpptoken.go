package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VppToken struct {
	AppleId                 *string                      `json:"appleId,omitempty"`
	AutomaticallyUpdateApps *bool                        `json:"automaticallyUpdateApps,omitempty"`
	CountryOrRegion         *string                      `json:"countryOrRegion,omitempty"`
	ExpirationDateTime      *string                      `json:"expirationDateTime,omitempty"`
	Id                      *string                      `json:"id,omitempty"`
	LastModifiedDateTime    *string                      `json:"lastModifiedDateTime,omitempty"`
	LastSyncDateTime        *string                      `json:"lastSyncDateTime,omitempty"`
	LastSyncStatus          *VppTokenLastSyncStatus      `json:"lastSyncStatus,omitempty"`
	ODataType               *string                      `json:"@odata.type,omitempty"`
	OrganizationName        *string                      `json:"organizationName,omitempty"`
	State                   *VppTokenState               `json:"state,omitempty"`
	Token                   *string                      `json:"token,omitempty"`
	VppTokenAccountType     *VppTokenVppTokenAccountType `json:"vppTokenAccountType,omitempty"`
}
