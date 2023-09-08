package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftStoreForBusinessApp struct {
	Assignments           *[]MobileAppAssignment                       `json:"assignments,omitempty"`
	Categories            *[]MobileAppCategory                         `json:"categories,omitempty"`
	CreatedDateTime       *string                                      `json:"createdDateTime,omitempty"`
	Description           *string                                      `json:"description,omitempty"`
	Developer             *string                                      `json:"developer,omitempty"`
	DisplayName           *string                                      `json:"displayName,omitempty"`
	Id                    *string                                      `json:"id,omitempty"`
	InformationUrl        *string                                      `json:"informationUrl,omitempty"`
	IsFeatured            *bool                                        `json:"isFeatured,omitempty"`
	LargeIcon             *MimeContent                                 `json:"largeIcon,omitempty"`
	LastModifiedDateTime  *string                                      `json:"lastModifiedDateTime,omitempty"`
	LicenseType           *MicrosoftStoreForBusinessAppLicenseType     `json:"licenseType,omitempty"`
	Notes                 *string                                      `json:"notes,omitempty"`
	ODataType             *string                                      `json:"@odata.type,omitempty"`
	Owner                 *string                                      `json:"owner,omitempty"`
	PackageIdentityName   *string                                      `json:"packageIdentityName,omitempty"`
	PrivacyInformationUrl *string                                      `json:"privacyInformationUrl,omitempty"`
	ProductKey            *string                                      `json:"productKey,omitempty"`
	Publisher             *string                                      `json:"publisher,omitempty"`
	PublishingState       *MicrosoftStoreForBusinessAppPublishingState `json:"publishingState,omitempty"`
	TotalLicenseCount     *int64                                       `json:"totalLicenseCount,omitempty"`
	UsedLicenseCount      *int64                                       `json:"usedLicenseCount,omitempty"`
}
