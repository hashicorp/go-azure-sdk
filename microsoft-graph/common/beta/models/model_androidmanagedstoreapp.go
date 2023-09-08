package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreApp struct {
	AppIdentifier         *string                                `json:"appIdentifier,omitempty"`
	AppStoreUrl           *string                                `json:"appStoreUrl,omitempty"`
	AppTracks             *[]AndroidManagedStoreAppTrack         `json:"appTracks,omitempty"`
	Assignments           *[]MobileAppAssignment                 `json:"assignments,omitempty"`
	Categories            *[]MobileAppCategory                   `json:"categories,omitempty"`
	CreatedDateTime       *string                                `json:"createdDateTime,omitempty"`
	DependentAppCount     *int64                                 `json:"dependentAppCount,omitempty"`
	Description           *string                                `json:"description,omitempty"`
	Developer             *string                                `json:"developer,omitempty"`
	DisplayName           *string                                `json:"displayName,omitempty"`
	Id                    *string                                `json:"id,omitempty"`
	InformationUrl        *string                                `json:"informationUrl,omitempty"`
	IsAssigned            *bool                                  `json:"isAssigned,omitempty"`
	IsFeatured            *bool                                  `json:"isFeatured,omitempty"`
	IsPrivate             *bool                                  `json:"isPrivate,omitempty"`
	IsSystemApp           *bool                                  `json:"isSystemApp,omitempty"`
	LargeIcon             *MimeContent                           `json:"largeIcon,omitempty"`
	LastModifiedDateTime  *string                                `json:"lastModifiedDateTime,omitempty"`
	Notes                 *string                                `json:"notes,omitempty"`
	ODataType             *string                                `json:"@odata.type,omitempty"`
	Owner                 *string                                `json:"owner,omitempty"`
	PackageId             *string                                `json:"packageId,omitempty"`
	PrivacyInformationUrl *string                                `json:"privacyInformationUrl,omitempty"`
	Publisher             *string                                `json:"publisher,omitempty"`
	PublishingState       *AndroidManagedStoreAppPublishingState `json:"publishingState,omitempty"`
	Relationships         *[]MobileAppRelationship               `json:"relationships,omitempty"`
	RoleScopeTagIds       *[]string                              `json:"roleScopeTagIds,omitempty"`
	SupersededAppCount    *int64                                 `json:"supersededAppCount,omitempty"`
	SupersedingAppCount   *int64                                 `json:"supersedingAppCount,omitempty"`
	SupportsOemConfig     *bool                                  `json:"supportsOemConfig,omitempty"`
	TotalLicenseCount     *int64                                 `json:"totalLicenseCount,omitempty"`
	UploadState           *int64                                 `json:"uploadState,omitempty"`
	UsedLicenseCount      *int64                                 `json:"usedLicenseCount,omitempty"`
}
