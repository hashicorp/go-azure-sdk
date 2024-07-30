package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfficeSuiteApp struct {
	Assignments                          *[]MobileAppAssignment                         `json:"assignments,omitempty"`
	AutoAcceptEula                       *bool                                          `json:"autoAcceptEula,omitempty"`
	Categories                           *[]MobileAppCategory                           `json:"categories,omitempty"`
	CreatedDateTime                      *string                                        `json:"createdDateTime,omitempty"`
	DependentAppCount                    *int64                                         `json:"dependentAppCount,omitempty"`
	Description                          *string                                        `json:"description,omitempty"`
	Developer                            *string                                        `json:"developer,omitempty"`
	DisplayName                          *string                                        `json:"displayName,omitempty"`
	ExcludedApps                         *ExcludedApps                                  `json:"excludedApps,omitempty"`
	Id                                   *string                                        `json:"id,omitempty"`
	InformationUrl                       *string                                        `json:"informationUrl,omitempty"`
	InstallProgressDisplayLevel          *OfficeSuiteAppInstallProgressDisplayLevel     `json:"installProgressDisplayLevel,omitempty"`
	IsAssigned                           *bool                                          `json:"isAssigned,omitempty"`
	IsFeatured                           *bool                                          `json:"isFeatured,omitempty"`
	LargeIcon                            *MimeContent                                   `json:"largeIcon,omitempty"`
	LastModifiedDateTime                 *string                                        `json:"lastModifiedDateTime,omitempty"`
	LocalesToInstall                     *[]string                                      `json:"localesToInstall,omitempty"`
	Notes                                *string                                        `json:"notes,omitempty"`
	ODataType                            *string                                        `json:"@odata.type,omitempty"`
	OfficeConfigurationXml               *string                                        `json:"officeConfigurationXml,omitempty"`
	OfficePlatformArchitecture           *OfficeSuiteAppOfficePlatformArchitecture      `json:"officePlatformArchitecture,omitempty"`
	OfficeSuiteAppDefaultFileFormat      *OfficeSuiteAppOfficeSuiteAppDefaultFileFormat `json:"officeSuiteAppDefaultFileFormat,omitempty"`
	Owner                                *string                                        `json:"owner,omitempty"`
	PrivacyInformationUrl                *string                                        `json:"privacyInformationUrl,omitempty"`
	ProductIds                           *OfficeSuiteAppProductIds                      `json:"productIds,omitempty"`
	Publisher                            *string                                        `json:"publisher,omitempty"`
	PublishingState                      *OfficeSuiteAppPublishingState                 `json:"publishingState,omitempty"`
	Relationships                        *[]MobileAppRelationship                       `json:"relationships,omitempty"`
	RoleScopeTagIds                      *[]string                                      `json:"roleScopeTagIds,omitempty"`
	ShouldUninstallOlderVersionsOfOffice *bool                                          `json:"shouldUninstallOlderVersionsOfOffice,omitempty"`
	SupersededAppCount                   *int64                                         `json:"supersededAppCount,omitempty"`
	SupersedingAppCount                  *int64                                         `json:"supersedingAppCount,omitempty"`
	TargetVersion                        *string                                        `json:"targetVersion,omitempty"`
	UpdateChannel                        *OfficeSuiteAppUpdateChannel                   `json:"updateChannel,omitempty"`
	UpdateVersion                        *string                                        `json:"updateVersion,omitempty"`
	UploadState                          *int64                                         `json:"uploadState,omitempty"`
	UseSharedComputerActivation          *bool                                          `json:"useSharedComputerActivation,omitempty"`
}
