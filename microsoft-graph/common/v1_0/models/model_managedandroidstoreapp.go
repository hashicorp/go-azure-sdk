package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAndroidStoreApp struct {
	AppAvailability                 *ManagedAndroidStoreAppAppAvailability `json:"appAvailability,omitempty"`
	AppStoreUrl                     *string                                `json:"appStoreUrl,omitempty"`
	Assignments                     *[]MobileAppAssignment                 `json:"assignments,omitempty"`
	Categories                      *[]MobileAppCategory                   `json:"categories,omitempty"`
	CreatedDateTime                 *string                                `json:"createdDateTime,omitempty"`
	Description                     *string                                `json:"description,omitempty"`
	Developer                       *string                                `json:"developer,omitempty"`
	DisplayName                     *string                                `json:"displayName,omitempty"`
	Id                              *string                                `json:"id,omitempty"`
	InformationUrl                  *string                                `json:"informationUrl,omitempty"`
	IsFeatured                      *bool                                  `json:"isFeatured,omitempty"`
	LargeIcon                       *MimeContent                           `json:"largeIcon,omitempty"`
	LastModifiedDateTime            *string                                `json:"lastModifiedDateTime,omitempty"`
	MinimumSupportedOperatingSystem *AndroidMinimumOperatingSystem         `json:"minimumSupportedOperatingSystem,omitempty"`
	Notes                           *string                                `json:"notes,omitempty"`
	ODataType                       *string                                `json:"@odata.type,omitempty"`
	Owner                           *string                                `json:"owner,omitempty"`
	PackageId                       *string                                `json:"packageId,omitempty"`
	PrivacyInformationUrl           *string                                `json:"privacyInformationUrl,omitempty"`
	Publisher                       *string                                `json:"publisher,omitempty"`
	PublishingState                 *ManagedAndroidStoreAppPublishingState `json:"publishingState,omitempty"`
	Version                         *string                                `json:"version,omitempty"`
}
