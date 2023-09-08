package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VppTokenLicenseSummary struct {
	AppleId               *string `json:"appleId,omitempty"`
	AvailableLicenseCount *int64  `json:"availableLicenseCount,omitempty"`
	ODataType             *string `json:"@odata.type,omitempty"`
	OrganizationName      *string `json:"organizationName,omitempty"`
	UsedLicenseCount      *int64  `json:"usedLicenseCount,omitempty"`
	VppTokenId            *string `json:"vppTokenId,omitempty"`
}
