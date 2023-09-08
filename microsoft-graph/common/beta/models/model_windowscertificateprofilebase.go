package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsCertificateProfileBase struct {
	Assignments                                 *[]DeviceConfigurationAssignment                             `json:"assignments,omitempty"`
	CertificateValidityPeriodScale              *WindowsCertificateProfileBaseCertificateValidityPeriodScale `json:"certificateValidityPeriodScale,omitempty"`
	CertificateValidityPeriodValue              *int64                                                       `json:"certificateValidityPeriodValue,omitempty"`
	CreatedDateTime                             *string                                                      `json:"createdDateTime,omitempty"`
	Description                                 *string                                                      `json:"description,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode                 `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition  *DeviceManagementApplicabilityRuleOsEdition                  `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion  *DeviceManagementApplicabilityRuleOsVersion                  `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceSettingStateSummaries                 *[]SettingStateDeviceSummary                                 `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                        *DeviceConfigurationDeviceOverview                           `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                              *[]DeviceConfigurationDeviceStatus                           `json:"deviceStatuses,omitempty"`
	DisplayName                                 *string                                                      `json:"displayName,omitempty"`
	GroupAssignments                            *[]DeviceConfigurationGroupAssignment                        `json:"groupAssignments,omitempty"`
	Id                                          *string                                                      `json:"id,omitempty"`
	KeyStorageProvider                          *WindowsCertificateProfileBaseKeyStorageProvider             `json:"keyStorageProvider,omitempty"`
	LastModifiedDateTime                        *string                                                      `json:"lastModifiedDateTime,omitempty"`
	ODataType                                   *string                                                      `json:"@odata.type,omitempty"`
	RenewalThresholdPercentage                  *int64                                                       `json:"renewalThresholdPercentage,omitempty"`
	RoleScopeTagIds                             *[]string                                                    `json:"roleScopeTagIds,omitempty"`
	SubjectAlternativeNameType                  *WindowsCertificateProfileBaseSubjectAlternativeNameType     `json:"subjectAlternativeNameType,omitempty"`
	SubjectNameFormat                           *WindowsCertificateProfileBaseSubjectNameFormat              `json:"subjectNameFormat,omitempty"`
	SupportsScopeTags                           *bool                                                        `json:"supportsScopeTags,omitempty"`
	UserStatusOverview                          *DeviceConfigurationUserOverview                             `json:"userStatusOverview,omitempty"`
	UserStatuses                                *[]DeviceConfigurationUserStatus                             `json:"userStatuses,omitempty"`
	Version                                     *int64                                                       `json:"version,omitempty"`
}
