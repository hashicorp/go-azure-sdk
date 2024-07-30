package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DepOnboardingSetting struct {
	AppleIdentifier                     *string                        `json:"appleIdentifier,omitempty"`
	DataSharingConsentGranted           *bool                          `json:"dataSharingConsentGranted,omitempty"`
	DefaultIosEnrollmentProfile         *DepIOSEnrollmentProfile       `json:"defaultIosEnrollmentProfile,omitempty"`
	DefaultMacOsEnrollmentProfile       *DepMacOSEnrollmentProfile     `json:"defaultMacOsEnrollmentProfile,omitempty"`
	EnrollmentProfiles                  *[]EnrollmentProfile           `json:"enrollmentProfiles,omitempty"`
	Id                                  *string                        `json:"id,omitempty"`
	ImportedAppleDeviceIdentities       *[]ImportedAppleDeviceIdentity `json:"importedAppleDeviceIdentities,omitempty"`
	LastModifiedDateTime                *string                        `json:"lastModifiedDateTime,omitempty"`
	LastSuccessfulSyncDateTime          *string                        `json:"lastSuccessfulSyncDateTime,omitempty"`
	LastSyncErrorCode                   *int64                         `json:"lastSyncErrorCode,omitempty"`
	LastSyncTriggeredDateTime           *string                        `json:"lastSyncTriggeredDateTime,omitempty"`
	ODataType                           *string                        `json:"@odata.type,omitempty"`
	RoleScopeTagIds                     *[]string                      `json:"roleScopeTagIds,omitempty"`
	ShareTokenWithSchoolDataSyncService *bool                          `json:"shareTokenWithSchoolDataSyncService,omitempty"`
	SyncedDeviceCount                   *int64                         `json:"syncedDeviceCount,omitempty"`
	TokenExpirationDateTime             *string                        `json:"tokenExpirationDateTime,omitempty"`
	TokenName                           *string                        `json:"tokenName,omitempty"`
	TokenType                           *DepOnboardingSettingTokenType `json:"tokenType,omitempty"`
}
