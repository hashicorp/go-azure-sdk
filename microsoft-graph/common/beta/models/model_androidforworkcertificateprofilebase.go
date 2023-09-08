package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkCertificateProfileBase struct {
	Assignments                                 *[]DeviceConfigurationAssignment                                    `json:"assignments,omitempty"`
	CertificateValidityPeriodScale              *AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScale `json:"certificateValidityPeriodScale,omitempty"`
	CertificateValidityPeriodValue              *int64                                                              `json:"certificateValidityPeriodValue,omitempty"`
	CreatedDateTime                             *string                                                             `json:"createdDateTime,omitempty"`
	Description                                 *string                                                             `json:"description,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode                        `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition  *DeviceManagementApplicabilityRuleOsEdition                         `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion  *DeviceManagementApplicabilityRuleOsVersion                         `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceSettingStateSummaries                 *[]SettingStateDeviceSummary                                        `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                        *DeviceConfigurationDeviceOverview                                  `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                              *[]DeviceConfigurationDeviceStatus                                  `json:"deviceStatuses,omitempty"`
	DisplayName                                 *string                                                             `json:"displayName,omitempty"`
	ExtendedKeyUsages                           *[]ExtendedKeyUsage                                                 `json:"extendedKeyUsages,omitempty"`
	GroupAssignments                            *[]DeviceConfigurationGroupAssignment                               `json:"groupAssignments,omitempty"`
	Id                                          *string                                                             `json:"id,omitempty"`
	LastModifiedDateTime                        *string                                                             `json:"lastModifiedDateTime,omitempty"`
	ODataType                                   *string                                                             `json:"@odata.type,omitempty"`
	RenewalThresholdPercentage                  *int64                                                              `json:"renewalThresholdPercentage,omitempty"`
	RoleScopeTagIds                             *[]string                                                           `json:"roleScopeTagIds,omitempty"`
	RootCertificate                             *AndroidForWorkTrustedRootCertificate                               `json:"rootCertificate,omitempty"`
	SubjectAlternativeNameType                  *AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType     `json:"subjectAlternativeNameType,omitempty"`
	SubjectNameFormat                           *AndroidForWorkCertificateProfileBaseSubjectNameFormat              `json:"subjectNameFormat,omitempty"`
	SupportsScopeTags                           *bool                                                               `json:"supportsScopeTags,omitempty"`
	UserStatusOverview                          *DeviceConfigurationUserOverview                                    `json:"userStatusOverview,omitempty"`
	UserStatuses                                *[]DeviceConfigurationUserStatus                                    `json:"userStatuses,omitempty"`
	Version                                     *int64                                                              `json:"version,omitempty"`
}
