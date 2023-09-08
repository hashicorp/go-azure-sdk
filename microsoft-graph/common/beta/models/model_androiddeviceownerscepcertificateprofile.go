package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerScepCertificateProfile struct {
	Assignments                                 *[]DeviceConfigurationAssignment                                        `json:"assignments,omitempty"`
	CertificateAccessType                       *AndroidDeviceOwnerScepCertificateProfileCertificateAccessType          `json:"certificateAccessType,omitempty"`
	CertificateStore                            *AndroidDeviceOwnerScepCertificateProfileCertificateStore               `json:"certificateStore,omitempty"`
	CertificateValidityPeriodScale              *AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale `json:"certificateValidityPeriodScale,omitempty"`
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
	HashAlgorithm                               *AndroidDeviceOwnerScepCertificateProfileHashAlgorithm                  `json:"hashAlgorithm,omitempty"`
	Id                                          *string                                                                 `json:"id,omitempty"`
	KeySize                                     *AndroidDeviceOwnerScepCertificateProfileKeySize                        `json:"keySize,omitempty"`
	KeyUsage                                    *AndroidDeviceOwnerScepCertificateProfileKeyUsage                       `json:"keyUsage,omitempty"`
	LastModifiedDateTime                        *string                                                                 `json:"lastModifiedDateTime,omitempty"`
	ManagedDeviceCertificateStates              *[]ManagedDeviceCertificateState                                        `json:"managedDeviceCertificateStates,omitempty"`
	ODataType                                   *string                                                                 `json:"@odata.type,omitempty"`
	RenewalThresholdPercentage                  *int64                                                                  `json:"renewalThresholdPercentage,omitempty"`
	RoleScopeTagIds                             *[]string                                                               `json:"roleScopeTagIds,omitempty"`
	RootCertificate                             *AndroidDeviceOwnerTrustedRootCertificate                               `json:"rootCertificate,omitempty"`
	ScepServerUrls                              *[]string                                                               `json:"scepServerUrls,omitempty"`
	SilentCertificateAccessDetails              *[]AndroidDeviceOwnerSilentCertificateAccess                            `json:"silentCertificateAccessDetails,omitempty"`
	SubjectAlternativeNameFormatString          *string                                                                 `json:"subjectAlternativeNameFormatString,omitempty"`
	SubjectAlternativeNameType                  *AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType     `json:"subjectAlternativeNameType,omitempty"`
	SubjectNameFormat                           *AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat              `json:"subjectNameFormat,omitempty"`
	SubjectNameFormatString                     *string                                                                 `json:"subjectNameFormatString,omitempty"`
	SupportsScopeTags                           *bool                                                                   `json:"supportsScopeTags,omitempty"`
	UserStatusOverview                          *DeviceConfigurationUserOverview                                        `json:"userStatusOverview,omitempty"`
	UserStatuses                                *[]DeviceConfigurationUserStatus                                        `json:"userStatuses,omitempty"`
	Version                                     *int64                                                                  `json:"version,omitempty"`
}
