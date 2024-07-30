package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosiPadOSWebClip struct {
	AppUrl                            *string                          `json:"appUrl,omitempty"`
	Assignments                       *[]MobileAppAssignment           `json:"assignments,omitempty"`
	Categories                        *[]MobileAppCategory             `json:"categories,omitempty"`
	CreatedDateTime                   *string                          `json:"createdDateTime,omitempty"`
	DependentAppCount                 *int64                           `json:"dependentAppCount,omitempty"`
	Description                       *string                          `json:"description,omitempty"`
	Developer                         *string                          `json:"developer,omitempty"`
	DisplayName                       *string                          `json:"displayName,omitempty"`
	FullScreenEnabled                 *bool                            `json:"fullScreenEnabled,omitempty"`
	Id                                *string                          `json:"id,omitempty"`
	IgnoreManifestScope               *bool                            `json:"ignoreManifestScope,omitempty"`
	InformationUrl                    *string                          `json:"informationUrl,omitempty"`
	IsAssigned                        *bool                            `json:"isAssigned,omitempty"`
	IsFeatured                        *bool                            `json:"isFeatured,omitempty"`
	LargeIcon                         *MimeContent                     `json:"largeIcon,omitempty"`
	LastModifiedDateTime              *string                          `json:"lastModifiedDateTime,omitempty"`
	Notes                             *string                          `json:"notes,omitempty"`
	ODataType                         *string                          `json:"@odata.type,omitempty"`
	Owner                             *string                          `json:"owner,omitempty"`
	PreComposedIconEnabled            *bool                            `json:"preComposedIconEnabled,omitempty"`
	PrivacyInformationUrl             *string                          `json:"privacyInformationUrl,omitempty"`
	Publisher                         *string                          `json:"publisher,omitempty"`
	PublishingState                   *IosiPadOSWebClipPublishingState `json:"publishingState,omitempty"`
	Relationships                     *[]MobileAppRelationship         `json:"relationships,omitempty"`
	RoleScopeTagIds                   *[]string                        `json:"roleScopeTagIds,omitempty"`
	SupersededAppCount                *int64                           `json:"supersededAppCount,omitempty"`
	SupersedingAppCount               *int64                           `json:"supersedingAppCount,omitempty"`
	TargetApplicationBundleIdentifier *string                          `json:"targetApplicationBundleIdentifier,omitempty"`
	UploadState                       *int64                           `json:"uploadState,omitempty"`
	UseManagedBrowser                 *bool                            `json:"useManagedBrowser,omitempty"`
}
