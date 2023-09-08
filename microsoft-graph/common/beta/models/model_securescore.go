package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecureScore struct {
	ActiveUserCount          *int64                     `json:"activeUserCount,omitempty"`
	AverageComparativeScores *[]AverageComparativeScore `json:"averageComparativeScores,omitempty"`
	AzureTenantId            *string                    `json:"azureTenantId,omitempty"`
	ControlScores            *[]ControlScore            `json:"controlScores,omitempty"`
	CreatedDateTime          *string                    `json:"createdDateTime,omitempty"`
	EnabledServices          *[]string                  `json:"enabledServices,omitempty"`
	Id                       *string                    `json:"id,omitempty"`
	LicensedUserCount        *int64                     `json:"licensedUserCount,omitempty"`
	ODataType                *string                    `json:"@odata.type,omitempty"`
	VendorInformation        *SecurityVendorInformation `json:"vendorInformation,omitempty"`
}
