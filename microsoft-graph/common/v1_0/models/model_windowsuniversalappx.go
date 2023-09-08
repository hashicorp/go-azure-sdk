package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUniversalAppX struct {
	ApplicableArchitectures         *WindowsUniversalAppXApplicableArchitectures `json:"applicableArchitectures,omitempty"`
	ApplicableDeviceTypes           *WindowsUniversalAppXApplicableDeviceTypes   `json:"applicableDeviceTypes,omitempty"`
	Assignments                     *[]MobileAppAssignment                       `json:"assignments,omitempty"`
	Categories                      *[]MobileAppCategory                         `json:"categories,omitempty"`
	CommittedContainedApps          *[]MobileContainedApp                        `json:"committedContainedApps,omitempty"`
	CommittedContentVersion         *string                                      `json:"committedContentVersion,omitempty"`
	ContentVersions                 *[]MobileAppContent                          `json:"contentVersions,omitempty"`
	CreatedDateTime                 *string                                      `json:"createdDateTime,omitempty"`
	Description                     *string                                      `json:"description,omitempty"`
	Developer                       *string                                      `json:"developer,omitempty"`
	DisplayName                     *string                                      `json:"displayName,omitempty"`
	FileName                        *string                                      `json:"fileName,omitempty"`
	Id                              *string                                      `json:"id,omitempty"`
	IdentityName                    *string                                      `json:"identityName,omitempty"`
	IdentityPublisherHash           *string                                      `json:"identityPublisherHash,omitempty"`
	IdentityResourceIdentifier      *string                                      `json:"identityResourceIdentifier,omitempty"`
	IdentityVersion                 *string                                      `json:"identityVersion,omitempty"`
	InformationUrl                  *string                                      `json:"informationUrl,omitempty"`
	IsBundle                        *bool                                        `json:"isBundle,omitempty"`
	IsFeatured                      *bool                                        `json:"isFeatured,omitempty"`
	LargeIcon                       *MimeContent                                 `json:"largeIcon,omitempty"`
	LastModifiedDateTime            *string                                      `json:"lastModifiedDateTime,omitempty"`
	MinimumSupportedOperatingSystem *WindowsMinimumOperatingSystem               `json:"minimumSupportedOperatingSystem,omitempty"`
	Notes                           *string                                      `json:"notes,omitempty"`
	ODataType                       *string                                      `json:"@odata.type,omitempty"`
	Owner                           *string                                      `json:"owner,omitempty"`
	PrivacyInformationUrl           *string                                      `json:"privacyInformationUrl,omitempty"`
	Publisher                       *string                                      `json:"publisher,omitempty"`
	PublishingState                 *WindowsUniversalAppXPublishingState         `json:"publishingState,omitempty"`
	Size                            *int64                                       `json:"size,omitempty"`
}
