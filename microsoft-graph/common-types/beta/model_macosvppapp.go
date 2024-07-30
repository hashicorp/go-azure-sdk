package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOsVppApp struct {
	AppStoreUrl                *string                                  `json:"appStoreUrl,omitempty"`
	AssignedLicenses           *[]MacOsVppAppAssignedLicense            `json:"assignedLicenses,omitempty"`
	Assignments                *[]MobileAppAssignment                   `json:"assignments,omitempty"`
	BundleId                   *string                                  `json:"bundleId,omitempty"`
	Categories                 *[]MobileAppCategory                     `json:"categories,omitempty"`
	CreatedDateTime            *string                                  `json:"createdDateTime,omitempty"`
	DependentAppCount          *int64                                   `json:"dependentAppCount,omitempty"`
	Description                *string                                  `json:"description,omitempty"`
	Developer                  *string                                  `json:"developer,omitempty"`
	DisplayName                *string                                  `json:"displayName,omitempty"`
	Id                         *string                                  `json:"id,omitempty"`
	InformationUrl             *string                                  `json:"informationUrl,omitempty"`
	IsAssigned                 *bool                                    `json:"isAssigned,omitempty"`
	IsFeatured                 *bool                                    `json:"isFeatured,omitempty"`
	LargeIcon                  *MimeContent                             `json:"largeIcon,omitempty"`
	LastModifiedDateTime       *string                                  `json:"lastModifiedDateTime,omitempty"`
	LicensingType              *VppLicensingType                        `json:"licensingType,omitempty"`
	Notes                      *string                                  `json:"notes,omitempty"`
	ODataType                  *string                                  `json:"@odata.type,omitempty"`
	Owner                      *string                                  `json:"owner,omitempty"`
	PrivacyInformationUrl      *string                                  `json:"privacyInformationUrl,omitempty"`
	Publisher                  *string                                  `json:"publisher,omitempty"`
	PublishingState            *MacOsVppAppPublishingState              `json:"publishingState,omitempty"`
	Relationships              *[]MobileAppRelationship                 `json:"relationships,omitempty"`
	ReleaseDateTime            *string                                  `json:"releaseDateTime,omitempty"`
	RevokeLicenseActionResults *[]MacOsVppAppRevokeLicensesActionResult `json:"revokeLicenseActionResults,omitempty"`
	RoleScopeTagIds            *[]string                                `json:"roleScopeTagIds,omitempty"`
	SupersededAppCount         *int64                                   `json:"supersededAppCount,omitempty"`
	SupersedingAppCount        *int64                                   `json:"supersedingAppCount,omitempty"`
	TotalLicenseCount          *int64                                   `json:"totalLicenseCount,omitempty"`
	UploadState                *int64                                   `json:"uploadState,omitempty"`
	UsedLicenseCount           *int64                                   `json:"usedLicenseCount,omitempty"`
	VppTokenAccountType        *MacOsVppAppVppTokenAccountType          `json:"vppTokenAccountType,omitempty"`
	VppTokenAppleId            *string                                  `json:"vppTokenAppleId,omitempty"`
	VppTokenId                 *string                                  `json:"vppTokenId,omitempty"`
	VppTokenOrganizationName   *string                                  `json:"vppTokenOrganizationName,omitempty"`
}
