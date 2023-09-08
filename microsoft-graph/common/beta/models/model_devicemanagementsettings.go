package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSettings struct {
	AndroidDeviceAdministratorEnrollmentEnabled *bool                                              `json:"androidDeviceAdministratorEnrollmentEnabled,omitempty"`
	DerivedCredentialProvider                   *DeviceManagementSettingsDerivedCredentialProvider `json:"derivedCredentialProvider,omitempty"`
	DerivedCredentialUrl                        *string                                            `json:"derivedCredentialUrl,omitempty"`
	DeviceComplianceCheckinThresholdDays        *int64                                             `json:"deviceComplianceCheckinThresholdDays,omitempty"`
	DeviceInactivityBeforeRetirementInDay       *int64                                             `json:"deviceInactivityBeforeRetirementInDay,omitempty"`
	EnableAutopilotDiagnostics                  *bool                                              `json:"enableAutopilotDiagnostics,omitempty"`
	EnableDeviceGroupMembershipReport           *bool                                              `json:"enableDeviceGroupMembershipReport,omitempty"`
	EnableEnhancedTroubleshootingExperience     *bool                                              `json:"enableEnhancedTroubleshootingExperience,omitempty"`
	EnableLogCollection                         *bool                                              `json:"enableLogCollection,omitempty"`
	EnhancedJailBreak                           *bool                                              `json:"enhancedJailBreak,omitempty"`
	IgnoreDevicesForUnsupportedSettingsEnabled  *bool                                              `json:"ignoreDevicesForUnsupportedSettingsEnabled,omitempty"`
	IsScheduledActionEnabled                    *bool                                              `json:"isScheduledActionEnabled,omitempty"`
	ODataType                                   *string                                            `json:"@odata.type,omitempty"`
	SecureByDefault                             *bool                                              `json:"secureByDefault,omitempty"`
}
