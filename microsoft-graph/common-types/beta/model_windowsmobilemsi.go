package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsMobileMSI struct {
	Assignments             *[]MobileAppAssignment           `json:"assignments,omitempty"`
	Categories              *[]MobileAppCategory             `json:"categories,omitempty"`
	CommandLine             *string                          `json:"commandLine,omitempty"`
	CommittedContentVersion *string                          `json:"committedContentVersion,omitempty"`
	ContentVersions         *[]MobileAppContent              `json:"contentVersions,omitempty"`
	CreatedDateTime         *string                          `json:"createdDateTime,omitempty"`
	DependentAppCount       *int64                           `json:"dependentAppCount,omitempty"`
	Description             *string                          `json:"description,omitempty"`
	Developer               *string                          `json:"developer,omitempty"`
	DisplayName             *string                          `json:"displayName,omitempty"`
	FileName                *string                          `json:"fileName,omitempty"`
	Id                      *string                          `json:"id,omitempty"`
	IdentityVersion         *string                          `json:"identityVersion,omitempty"`
	IgnoreVersionDetection  *bool                            `json:"ignoreVersionDetection,omitempty"`
	InformationUrl          *string                          `json:"informationUrl,omitempty"`
	IsAssigned              *bool                            `json:"isAssigned,omitempty"`
	IsFeatured              *bool                            `json:"isFeatured,omitempty"`
	LargeIcon               *MimeContent                     `json:"largeIcon,omitempty"`
	LastModifiedDateTime    *string                          `json:"lastModifiedDateTime,omitempty"`
	Notes                   *string                          `json:"notes,omitempty"`
	ODataType               *string                          `json:"@odata.type,omitempty"`
	Owner                   *string                          `json:"owner,omitempty"`
	PrivacyInformationUrl   *string                          `json:"privacyInformationUrl,omitempty"`
	ProductCode             *string                          `json:"productCode,omitempty"`
	ProductVersion          *string                          `json:"productVersion,omitempty"`
	Publisher               *string                          `json:"publisher,omitempty"`
	PublishingState         *WindowsMobileMSIPublishingState `json:"publishingState,omitempty"`
	Relationships           *[]MobileAppRelationship         `json:"relationships,omitempty"`
	RoleScopeTagIds         *[]string                        `json:"roleScopeTagIds,omitempty"`
	Size                    *int64                           `json:"size,omitempty"`
	SupersededAppCount      *int64                           `json:"supersededAppCount,omitempty"`
	SupersedingAppCount     *int64                           `json:"supersedingAppCount,omitempty"`
	UploadState             *int64                           `json:"uploadState,omitempty"`
	UseDeviceContext        *bool                            `json:"useDeviceContext,omitempty"`
}
