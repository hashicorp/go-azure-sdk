package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Office365ActiveUserDetail struct {
	AssignedProducts                  *[]string `json:"assignedProducts,omitempty"`
	DeletedDate                       *string   `json:"deletedDate,omitempty"`
	DisplayName                       *string   `json:"displayName,omitempty"`
	ExchangeLastActivityDate          *string   `json:"exchangeLastActivityDate,omitempty"`
	ExchangeLicenseAssignDate         *string   `json:"exchangeLicenseAssignDate,omitempty"`
	HasExchangeLicense                *bool     `json:"hasExchangeLicense,omitempty"`
	HasOneDriveLicense                *bool     `json:"hasOneDriveLicense,omitempty"`
	HasSharePointLicense              *bool     `json:"hasSharePointLicense,omitempty"`
	HasSkypeForBusinessLicense        *bool     `json:"hasSkypeForBusinessLicense,omitempty"`
	HasTeamsLicense                   *bool     `json:"hasTeamsLicense,omitempty"`
	HasYammerLicense                  *bool     `json:"hasYammerLicense,omitempty"`
	Id                                *string   `json:"id,omitempty"`
	IsDeleted                         *bool     `json:"isDeleted,omitempty"`
	ODataType                         *string   `json:"@odata.type,omitempty"`
	OneDriveLastActivityDate          *string   `json:"oneDriveLastActivityDate,omitempty"`
	OneDriveLicenseAssignDate         *string   `json:"oneDriveLicenseAssignDate,omitempty"`
	ReportRefreshDate                 *string   `json:"reportRefreshDate,omitempty"`
	SharePointLastActivityDate        *string   `json:"sharePointLastActivityDate,omitempty"`
	SharePointLicenseAssignDate       *string   `json:"sharePointLicenseAssignDate,omitempty"`
	SkypeForBusinessLastActivityDate  *string   `json:"skypeForBusinessLastActivityDate,omitempty"`
	SkypeForBusinessLicenseAssignDate *string   `json:"skypeForBusinessLicenseAssignDate,omitempty"`
	TeamsLastActivityDate             *string   `json:"teamsLastActivityDate,omitempty"`
	TeamsLicenseAssignDate            *string   `json:"teamsLicenseAssignDate,omitempty"`
	UserPrincipalName                 *string   `json:"userPrincipalName,omitempty"`
	YammerLastActivityDate            *string   `json:"yammerLastActivityDate,omitempty"`
	YammerLicenseAssignDate           *string   `json:"yammerLicenseAssignDate,omitempty"`
}
