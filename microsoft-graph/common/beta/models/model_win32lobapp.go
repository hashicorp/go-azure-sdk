package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobApp struct {
	AllowAvailableUninstall         *bool                               `json:"allowAvailableUninstall,omitempty"`
	ApplicableArchitectures         *Win32LobAppApplicableArchitectures `json:"applicableArchitectures,omitempty"`
	Assignments                     *[]MobileAppAssignment              `json:"assignments,omitempty"`
	Categories                      *[]MobileAppCategory                `json:"categories,omitempty"`
	CommittedContentVersion         *string                             `json:"committedContentVersion,omitempty"`
	ContentVersions                 *[]MobileAppContent                 `json:"contentVersions,omitempty"`
	CreatedDateTime                 *string                             `json:"createdDateTime,omitempty"`
	DependentAppCount               *int64                              `json:"dependentAppCount,omitempty"`
	Description                     *string                             `json:"description,omitempty"`
	DetectionRules                  *[]Win32LobAppDetection             `json:"detectionRules,omitempty"`
	Developer                       *string                             `json:"developer,omitempty"`
	DisplayName                     *string                             `json:"displayName,omitempty"`
	DisplayVersion                  *string                             `json:"displayVersion,omitempty"`
	FileName                        *string                             `json:"fileName,omitempty"`
	Id                              *string                             `json:"id,omitempty"`
	InformationUrl                  *string                             `json:"informationUrl,omitempty"`
	InstallCommandLine              *string                             `json:"installCommandLine,omitempty"`
	InstallExperience               *Win32LobAppInstallExperience       `json:"installExperience,omitempty"`
	IsAssigned                      *bool                               `json:"isAssigned,omitempty"`
	IsFeatured                      *bool                               `json:"isFeatured,omitempty"`
	LargeIcon                       *MimeContent                        `json:"largeIcon,omitempty"`
	LastModifiedDateTime            *string                             `json:"lastModifiedDateTime,omitempty"`
	MinimumCpuSpeedInMHz            *int64                              `json:"minimumCpuSpeedInMHz,omitempty"`
	MinimumFreeDiskSpaceInMB        *int64                              `json:"minimumFreeDiskSpaceInMB,omitempty"`
	MinimumMemoryInMB               *int64                              `json:"minimumMemoryInMB,omitempty"`
	MinimumNumberOfProcessors       *int64                              `json:"minimumNumberOfProcessors,omitempty"`
	MinimumSupportedOperatingSystem *WindowsMinimumOperatingSystem      `json:"minimumSupportedOperatingSystem,omitempty"`
	MinimumSupportedWindowsRelease  *string                             `json:"minimumSupportedWindowsRelease,omitempty"`
	MsiInformation                  *Win32LobAppMsiInformation          `json:"msiInformation,omitempty"`
	Notes                           *string                             `json:"notes,omitempty"`
	ODataType                       *string                             `json:"@odata.type,omitempty"`
	Owner                           *string                             `json:"owner,omitempty"`
	PrivacyInformationUrl           *string                             `json:"privacyInformationUrl,omitempty"`
	Publisher                       *string                             `json:"publisher,omitempty"`
	PublishingState                 *Win32LobAppPublishingState         `json:"publishingState,omitempty"`
	Relationships                   *[]MobileAppRelationship            `json:"relationships,omitempty"`
	RequirementRules                *[]Win32LobAppRequirement           `json:"requirementRules,omitempty"`
	ReturnCodes                     *[]Win32LobAppReturnCode            `json:"returnCodes,omitempty"`
	RoleScopeTagIds                 *[]string                           `json:"roleScopeTagIds,omitempty"`
	Rules                           *[]Win32LobAppRule                  `json:"rules,omitempty"`
	SetupFilePath                   *string                             `json:"setupFilePath,omitempty"`
	Size                            *int64                              `json:"size,omitempty"`
	SupersededAppCount              *int64                              `json:"supersededAppCount,omitempty"`
	SupersedingAppCount             *int64                              `json:"supersedingAppCount,omitempty"`
	UninstallCommandLine            *string                             `json:"uninstallCommandLine,omitempty"`
	UploadState                     *int64                              `json:"uploadState,omitempty"`
}
