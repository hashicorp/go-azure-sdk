package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftStoreForBusinessApp struct {
	Assignments           *[]MobileAppAssignment                       `json:"assignments,omitempty"`
	Categories            *[]MobileAppCategory                         `json:"categories,omitempty"`
	ContainedApps         *[]MobileContainedApp                        `json:"containedApps,omitempty"`
	CreatedDateTime       *string                                      `json:"createdDateTime,omitempty"`
	DependentAppCount     *int64                                       `json:"dependentAppCount,omitempty"`
	Description           *string                                      `json:"description,omitempty"`
	Developer             *string                                      `json:"developer,omitempty"`
	DisplayName           *string                                      `json:"displayName,omitempty"`
	Id                    *string                                      `json:"id,omitempty"`
	InformationUrl        *string                                      `json:"informationUrl,omitempty"`
	IsAssigned            *bool                                        `json:"isAssigned,omitempty"`
	IsFeatured            *bool                                        `json:"isFeatured,omitempty"`
	LargeIcon             *MimeContent                                 `json:"largeIcon,omitempty"`
	LastModifiedDateTime  *string                                      `json:"lastModifiedDateTime,omitempty"`
	LicenseType           *MicrosoftStoreForBusinessAppLicenseType     `json:"licenseType,omitempty"`
	LicensingType         *VppLicensingType                            `json:"licensingType,omitempty"`
	Notes                 *string                                      `json:"notes,omitempty"`
	ODataType             *string                                      `json:"@odata.type,omitempty"`
	Owner                 *string                                      `json:"owner,omitempty"`
	PackageIdentityName   *string                                      `json:"packageIdentityName,omitempty"`
	PrivacyInformationUrl *string                                      `json:"privacyInformationUrl,omitempty"`
	ProductKey            *string                                      `json:"productKey,omitempty"`
	Publisher             *string                                      `json:"publisher,omitempty"`
	PublishingState       *MicrosoftStoreForBusinessAppPublishingState `json:"publishingState,omitempty"`
	Relationships         *[]MobileAppRelationship                     `json:"relationships,omitempty"`
	RoleScopeTagIds       *[]string                                    `json:"roleScopeTagIds,omitempty"`
	SupersededAppCount    *int64                                       `json:"supersededAppCount,omitempty"`
	SupersedingAppCount   *int64                                       `json:"supersedingAppCount,omitempty"`
	TotalLicenseCount     *int64                                       `json:"totalLicenseCount,omitempty"`
	UploadState           *int64                                       `json:"uploadState,omitempty"`
	UsedLicenseCount      *int64                                       `json:"usedLicenseCount,omitempty"`
}
