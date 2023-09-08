package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedIOSStoreApp struct {
	AppAvailability                 *ManagedIOSStoreAppAppAvailability `json:"appAvailability,omitempty"`
	AppStoreUrl                     *string                            `json:"appStoreUrl,omitempty"`
	ApplicableDeviceType            *IosDeviceType                     `json:"applicableDeviceType,omitempty"`
	Assignments                     *[]MobileAppAssignment             `json:"assignments,omitempty"`
	BundleId                        *string                            `json:"bundleId,omitempty"`
	Categories                      *[]MobileAppCategory               `json:"categories,omitempty"`
	CreatedDateTime                 *string                            `json:"createdDateTime,omitempty"`
	Description                     *string                            `json:"description,omitempty"`
	Developer                       *string                            `json:"developer,omitempty"`
	DisplayName                     *string                            `json:"displayName,omitempty"`
	Id                              *string                            `json:"id,omitempty"`
	InformationUrl                  *string                            `json:"informationUrl,omitempty"`
	IsFeatured                      *bool                              `json:"isFeatured,omitempty"`
	LargeIcon                       *MimeContent                       `json:"largeIcon,omitempty"`
	LastModifiedDateTime            *string                            `json:"lastModifiedDateTime,omitempty"`
	MinimumSupportedOperatingSystem *IosMinimumOperatingSystem         `json:"minimumSupportedOperatingSystem,omitempty"`
	Notes                           *string                            `json:"notes,omitempty"`
	ODataType                       *string                            `json:"@odata.type,omitempty"`
	Owner                           *string                            `json:"owner,omitempty"`
	PrivacyInformationUrl           *string                            `json:"privacyInformationUrl,omitempty"`
	Publisher                       *string                            `json:"publisher,omitempty"`
	PublishingState                 *ManagedIOSStoreAppPublishingState `json:"publishingState,omitempty"`
	Version                         *string                            `json:"version,omitempty"`
}
