package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSLobApp struct {
	Assignments                     *[]MobileAppAssignment       `json:"assignments,omitempty"`
	BuildNumber                     *string                      `json:"buildNumber,omitempty"`
	BundleId                        *string                      `json:"bundleId,omitempty"`
	Categories                      *[]MobileAppCategory         `json:"categories,omitempty"`
	ChildApps                       *[]MacOSLobChildApp          `json:"childApps,omitempty"`
	CommittedContentVersion         *string                      `json:"committedContentVersion,omitempty"`
	ContentVersions                 *[]MobileAppContent          `json:"contentVersions,omitempty"`
	CreatedDateTime                 *string                      `json:"createdDateTime,omitempty"`
	Description                     *string                      `json:"description,omitempty"`
	Developer                       *string                      `json:"developer,omitempty"`
	DisplayName                     *string                      `json:"displayName,omitempty"`
	FileName                        *string                      `json:"fileName,omitempty"`
	Id                              *string                      `json:"id,omitempty"`
	IgnoreVersionDetection          *bool                        `json:"ignoreVersionDetection,omitempty"`
	InformationUrl                  *string                      `json:"informationUrl,omitempty"`
	InstallAsManaged                *bool                        `json:"installAsManaged,omitempty"`
	IsFeatured                      *bool                        `json:"isFeatured,omitempty"`
	LargeIcon                       *MimeContent                 `json:"largeIcon,omitempty"`
	LastModifiedDateTime            *string                      `json:"lastModifiedDateTime,omitempty"`
	Md5Hash                         *[]string                    `json:"md5Hash,omitempty"`
	Md5HashChunkSize                *int64                       `json:"md5HashChunkSize,omitempty"`
	MinimumSupportedOperatingSystem *MacOSMinimumOperatingSystem `json:"minimumSupportedOperatingSystem,omitempty"`
	Notes                           *string                      `json:"notes,omitempty"`
	ODataType                       *string                      `json:"@odata.type,omitempty"`
	Owner                           *string                      `json:"owner,omitempty"`
	PrivacyInformationUrl           *string                      `json:"privacyInformationUrl,omitempty"`
	Publisher                       *string                      `json:"publisher,omitempty"`
	PublishingState                 *MacOSLobAppPublishingState  `json:"publishingState,omitempty"`
	Size                            *int64                       `json:"size,omitempty"`
	VersionNumber                   *string                      `json:"versionNumber,omitempty"`
}
