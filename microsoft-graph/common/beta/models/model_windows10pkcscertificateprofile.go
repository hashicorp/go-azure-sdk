package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10PkcsCertificateProfile struct {
	Assignments                                 *[]DeviceConfigurationAssignment                               `json:"assignments,omitempty"`
	CertificateStore                            *Windows10PkcsCertificateProfileCertificateStore               `json:"certificateStore,omitempty"`
	CertificateTemplateName                     *string                                                        `json:"certificateTemplateName,omitempty"`
	CertificateValidityPeriodScale              *Windows10PkcsCertificateProfileCertificateValidityPeriodScale `json:"certificateValidityPeriodScale,omitempty"`
	CertificateValidityPeriodValue              *int64                                                         `json:"certificateValidityPeriodValue,omitempty"`
	CertificationAuthority                      *string                                                        `json:"certificationAuthority,omitempty"`
	CertificationAuthorityName                  *string                                                        `json:"certificationAuthorityName,omitempty"`
	CreatedDateTime                             *string                                                        `json:"createdDateTime,omitempty"`
	CustomSubjectAlternativeNames               *[]CustomSubjectAlternativeName                                `json:"customSubjectAlternativeNames,omitempty"`
	Description                                 *string                                                        `json:"description,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode                   `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition  *DeviceManagementApplicabilityRuleOsEdition                    `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion  *DeviceManagementApplicabilityRuleOsVersion                    `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceSettingStateSummaries                 *[]SettingStateDeviceSummary                                   `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                        *DeviceConfigurationDeviceOverview                             `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                              *[]DeviceConfigurationDeviceStatus                             `json:"deviceStatuses,omitempty"`
	DisplayName                                 *string                                                        `json:"displayName,omitempty"`
	ExtendedKeyUsages                           *[]ExtendedKeyUsage                                            `json:"extendedKeyUsages,omitempty"`
	GroupAssignments                            *[]DeviceConfigurationGroupAssignment                          `json:"groupAssignments,omitempty"`
	Id                                          *string                                                        `json:"id,omitempty"`
	KeyStorageProvider                          *Windows10PkcsCertificateProfileKeyStorageProvider             `json:"keyStorageProvider,omitempty"`
	LastModifiedDateTime                        *string                                                        `json:"lastModifiedDateTime,omitempty"`
	ManagedDeviceCertificateStates              *[]ManagedDeviceCertificateState                               `json:"managedDeviceCertificateStates,omitempty"`
	ODataType                                   *string                                                        `json:"@odata.type,omitempty"`
	RenewalThresholdPercentage                  *int64                                                         `json:"renewalThresholdPercentage,omitempty"`
	RoleScopeTagIds                             *[]string                                                      `json:"roleScopeTagIds,omitempty"`
	SubjectAlternativeNameFormatString          *string                                                        `json:"subjectAlternativeNameFormatString,omitempty"`
	SubjectAlternativeNameType                  *Windows10PkcsCertificateProfileSubjectAlternativeNameType     `json:"subjectAlternativeNameType,omitempty"`
	SubjectNameFormat                           *Windows10PkcsCertificateProfileSubjectNameFormat              `json:"subjectNameFormat,omitempty"`
	SubjectNameFormatString                     *string                                                        `json:"subjectNameFormatString,omitempty"`
	SupportsScopeTags                           *bool                                                          `json:"supportsScopeTags,omitempty"`
	UserStatusOverview                          *DeviceConfigurationUserOverview                               `json:"userStatusOverview,omitempty"`
	UserStatuses                                *[]DeviceConfigurationUserStatus                               `json:"userStatuses,omitempty"`
	Version                                     *int64                                                         `json:"version,omitempty"`
}
