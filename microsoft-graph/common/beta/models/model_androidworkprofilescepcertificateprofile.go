package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileScepCertificateProfile struct {
	Assignments                                 *[]DeviceConfigurationAssignment                                        `json:"assignments,omitempty"`
	CertificateStore                            *AndroidWorkProfileScepCertificateProfileCertificateStore               `json:"certificateStore,omitempty"`
	CertificateValidityPeriodScale              *AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScale `json:"certificateValidityPeriodScale,omitempty"`
	CertificateValidityPeriodValue              *int64                                                                  `json:"certificateValidityPeriodValue,omitempty"`
	CreatedDateTime                             *string                                                                 `json:"createdDateTime,omitempty"`
	CustomSubjectAlternativeNames               *[]CustomSubjectAlternativeName                                         `json:"customSubjectAlternativeNames,omitempty"`
	Description                                 *string                                                                 `json:"description,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode                            `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition  *DeviceManagementApplicabilityRuleOsEdition                             `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion  *DeviceManagementApplicabilityRuleOsVersion                             `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceSettingStateSummaries                 *[]SettingStateDeviceSummary                                            `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                        *DeviceConfigurationDeviceOverview                                      `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                              *[]DeviceConfigurationDeviceStatus                                      `json:"deviceStatuses,omitempty"`
	DisplayName                                 *string                                                                 `json:"displayName,omitempty"`
	ExtendedKeyUsages                           *[]ExtendedKeyUsage                                                     `json:"extendedKeyUsages,omitempty"`
	GroupAssignments                            *[]DeviceConfigurationGroupAssignment                                   `json:"groupAssignments,omitempty"`
	HashAlgorithm                               *AndroidWorkProfileScepCertificateProfileHashAlgorithm                  `json:"hashAlgorithm,omitempty"`
	Id                                          *string                                                                 `json:"id,omitempty"`
	KeySize                                     *AndroidWorkProfileScepCertificateProfileKeySize                        `json:"keySize,omitempty"`
	KeyUsage                                    *AndroidWorkProfileScepCertificateProfileKeyUsage                       `json:"keyUsage,omitempty"`
	LastModifiedDateTime                        *string                                                                 `json:"lastModifiedDateTime,omitempty"`
	ManagedDeviceCertificateStates              *[]ManagedDeviceCertificateState                                        `json:"managedDeviceCertificateStates,omitempty"`
	ODataType                                   *string                                                                 `json:"@odata.type,omitempty"`
	RenewalThresholdPercentage                  *int64                                                                  `json:"renewalThresholdPercentage,omitempty"`
	RoleScopeTagIds                             *[]string                                                               `json:"roleScopeTagIds,omitempty"`
	RootCertificate                             *AndroidWorkProfileTrustedRootCertificate                               `json:"rootCertificate,omitempty"`
	ScepServerUrls                              *[]string                                                               `json:"scepServerUrls,omitempty"`
	SubjectAlternativeNameFormatString          *string                                                                 `json:"subjectAlternativeNameFormatString,omitempty"`
	SubjectAlternativeNameType                  *AndroidWorkProfileScepCertificateProfileSubjectAlternativeNameType     `json:"subjectAlternativeNameType,omitempty"`
	SubjectNameFormat                           *AndroidWorkProfileScepCertificateProfileSubjectNameFormat              `json:"subjectNameFormat,omitempty"`
	SubjectNameFormatString                     *string                                                                 `json:"subjectNameFormatString,omitempty"`
	SupportsScopeTags                           *bool                                                                   `json:"supportsScopeTags,omitempty"`
	UserStatusOverview                          *DeviceConfigurationUserOverview                                        `json:"userStatusOverview,omitempty"`
	UserStatuses                                *[]DeviceConfigurationUserStatus                                        `json:"userStatuses,omitempty"`
	Version                                     *int64                                                                  `json:"version,omitempty"`
}
