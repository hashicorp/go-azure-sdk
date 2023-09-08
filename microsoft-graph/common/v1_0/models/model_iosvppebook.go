package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVppEBook struct {
	AppleId               *string                    `json:"appleId,omitempty"`
	Assignments           *[]ManagedEBookAssignment  `json:"assignments,omitempty"`
	CreatedDateTime       *string                    `json:"createdDateTime,omitempty"`
	Description           *string                    `json:"description,omitempty"`
	DeviceStates          *[]DeviceInstallState      `json:"deviceStates,omitempty"`
	DisplayName           *string                    `json:"displayName,omitempty"`
	Genres                *[]string                  `json:"genres,omitempty"`
	Id                    *string                    `json:"id,omitempty"`
	InformationUrl        *string                    `json:"informationUrl,omitempty"`
	InstallSummary        *EBookInstallSummary       `json:"installSummary,omitempty"`
	Language              *string                    `json:"language,omitempty"`
	LargeCover            *MimeContent               `json:"largeCover,omitempty"`
	LastModifiedDateTime  *string                    `json:"lastModifiedDateTime,omitempty"`
	ODataType             *string                    `json:"@odata.type,omitempty"`
	PrivacyInformationUrl *string                    `json:"privacyInformationUrl,omitempty"`
	PublishedDateTime     *string                    `json:"publishedDateTime,omitempty"`
	Publisher             *string                    `json:"publisher,omitempty"`
	Seller                *string                    `json:"seller,omitempty"`
	TotalLicenseCount     *int64                     `json:"totalLicenseCount,omitempty"`
	UsedLicenseCount      *int64                     `json:"usedLicenseCount,omitempty"`
	UserStateSummary      *[]UserInstallStateSummary `json:"userStateSummary,omitempty"`
	VppOrganizationName   *string                    `json:"vppOrganizationName,omitempty"`
	VppTokenId            *string                    `json:"vppTokenId,omitempty"`
}
