package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobApp struct {
	ApplicableArchitectures        *Win32LobAppApplicableArchitectures `json:"applicableArchitectures,omitempty"`
	Assignments                    *[]MobileAppAssignment              `json:"assignments,omitempty"`
	Categories                     *[]MobileAppCategory                `json:"categories,omitempty"`
	CommittedContentVersion        *string                             `json:"committedContentVersion,omitempty"`
	ContentVersions                *[]MobileAppContent                 `json:"contentVersions,omitempty"`
	CreatedDateTime                *string                             `json:"createdDateTime,omitempty"`
	Description                    *string                             `json:"description,omitempty"`
	Developer                      *string                             `json:"developer,omitempty"`
	DisplayName                    *string                             `json:"displayName,omitempty"`
	FileName                       *string                             `json:"fileName,omitempty"`
	Id                             *string                             `json:"id,omitempty"`
	InformationUrl                 *string                             `json:"informationUrl,omitempty"`
	InstallCommandLine             *string                             `json:"installCommandLine,omitempty"`
	InstallExperience              *Win32LobAppInstallExperience       `json:"installExperience,omitempty"`
	IsFeatured                     *bool                               `json:"isFeatured,omitempty"`
	LargeIcon                      *MimeContent                        `json:"largeIcon,omitempty"`
	LastModifiedDateTime           *string                             `json:"lastModifiedDateTime,omitempty"`
	MinimumCpuSpeedInMHz           *int64                              `json:"minimumCpuSpeedInMHz,omitempty"`
	MinimumFreeDiskSpaceInMB       *int64                              `json:"minimumFreeDiskSpaceInMB,omitempty"`
	MinimumMemoryInMB              *int64                              `json:"minimumMemoryInMB,omitempty"`
	MinimumNumberOfProcessors      *int64                              `json:"minimumNumberOfProcessors,omitempty"`
	MinimumSupportedWindowsRelease *string                             `json:"minimumSupportedWindowsRelease,omitempty"`
	MsiInformation                 *Win32LobAppMsiInformation          `json:"msiInformation,omitempty"`
	Notes                          *string                             `json:"notes,omitempty"`
	ODataType                      *string                             `json:"@odata.type,omitempty"`
	Owner                          *string                             `json:"owner,omitempty"`
	PrivacyInformationUrl          *string                             `json:"privacyInformationUrl,omitempty"`
	Publisher                      *string                             `json:"publisher,omitempty"`
	PublishingState                *Win32LobAppPublishingState         `json:"publishingState,omitempty"`
	ReturnCodes                    *[]Win32LobAppReturnCode            `json:"returnCodes,omitempty"`
	Rules                          *[]Win32LobAppRule                  `json:"rules,omitempty"`
	SetupFilePath                  *string                             `json:"setupFilePath,omitempty"`
	Size                           *int64                              `json:"size,omitempty"`
	UninstallCommandLine           *string                             `json:"uninstallCommandLine,omitempty"`
}
