package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedPCConfiguration struct {
	AccountManagerPolicy                        *SharedPCAccountManagerPolicy                `json:"accountManagerPolicy,omitempty"`
	AllowLocalStorage                           *bool                                        `json:"allowLocalStorage,omitempty"`
	AllowedAccounts                             *SharedPCConfigurationAllowedAccounts        `json:"allowedAccounts,omitempty"`
	Assignments                                 *[]DeviceConfigurationAssignment             `json:"assignments,omitempty"`
	CreatedDateTime                             *string                                      `json:"createdDateTime,omitempty"`
	Description                                 *string                                      `json:"description,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition  *DeviceManagementApplicabilityRuleOsEdition  `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion  *DeviceManagementApplicabilityRuleOsVersion  `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceSettingStateSummaries                 *[]SettingStateDeviceSummary                 `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                        *DeviceConfigurationDeviceOverview           `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                              *[]DeviceConfigurationDeviceStatus           `json:"deviceStatuses,omitempty"`
	DisableAccountManager                       *bool                                        `json:"disableAccountManager,omitempty"`
	DisableEduPolicies                          *bool                                        `json:"disableEduPolicies,omitempty"`
	DisablePowerPolicies                        *bool                                        `json:"disablePowerPolicies,omitempty"`
	DisableSignInOnResume                       *bool                                        `json:"disableSignInOnResume,omitempty"`
	DisplayName                                 *string                                      `json:"displayName,omitempty"`
	Enabled                                     *bool                                        `json:"enabled,omitempty"`
	FastFirstSignIn                             *SharedPCConfigurationFastFirstSignIn        `json:"fastFirstSignIn,omitempty"`
	GroupAssignments                            *[]DeviceConfigurationGroupAssignment        `json:"groupAssignments,omitempty"`
	Id                                          *string                                      `json:"id,omitempty"`
	IdleTimeBeforeSleepInSeconds                *int64                                       `json:"idleTimeBeforeSleepInSeconds,omitempty"`
	KioskAppDisplayName                         *string                                      `json:"kioskAppDisplayName,omitempty"`
	KioskAppUserModelId                         *string                                      `json:"kioskAppUserModelId,omitempty"`
	LastModifiedDateTime                        *string                                      `json:"lastModifiedDateTime,omitempty"`
	LocalStorage                                *SharedPCConfigurationLocalStorage           `json:"localStorage,omitempty"`
	MaintenanceStartTime                        *string                                      `json:"maintenanceStartTime,omitempty"`
	ODataType                                   *string                                      `json:"@odata.type,omitempty"`
	RoleScopeTagIds                             *[]string                                    `json:"roleScopeTagIds,omitempty"`
	SetAccountManager                           *SharedPCConfigurationSetAccountManager      `json:"setAccountManager,omitempty"`
	SetEduPolicies                              *SharedPCConfigurationSetEduPolicies         `json:"setEduPolicies,omitempty"`
	SetPowerPolicies                            *SharedPCConfigurationSetPowerPolicies       `json:"setPowerPolicies,omitempty"`
	SignInOnResume                              *SharedPCConfigurationSignInOnResume         `json:"signInOnResume,omitempty"`
	SupportsScopeTags                           *bool                                        `json:"supportsScopeTags,omitempty"`
	UserStatusOverview                          *DeviceConfigurationUserOverview             `json:"userStatusOverview,omitempty"`
	UserStatuses                                *[]DeviceConfigurationUserStatus             `json:"userStatuses,omitempty"`
	Version                                     *int64                                       `json:"version,omitempty"`
}
