package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerPkcsCertificateProfile struct {
	Assignments                                 *[]DeviceConfigurationAssignment                                        `json:"assignments,omitempty"`
	CertificateAccessType                       *AndroidDeviceOwnerPkcsCertificateProfileCertificateAccessType          `json:"certificateAccessType,omitempty"`
	CertificateStore                            *AndroidDeviceOwnerPkcsCertificateProfileCertificateStore               `json:"certificateStore,omitempty"`
	CertificateTemplateName                     *string                                                                 `json:"certificateTemplateName,omitempty"`
	CertificateValidityPeriodScale              *AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale `json:"certificateValidityPeriodScale,omitempty"`
	CertificateValidityPeriodValue              *int64                                                                  `json:"certificateValidityPeriodValue,omitempty"`
	CertificationAuthority                      *string                                                                 `json:"certificationAuthority,omitempty"`
	CertificationAuthorityName                  *string                                                                 `json:"certificationAuthorityName,omitempty"`
	CertificationAuthorityType                  *AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityType     `json:"certificationAuthorityType,omitempty"`
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
	Id                                          *string                                                                 `json:"id,omitempty"`
	LastModifiedDateTime                        *string                                                                 `json:"lastModifiedDateTime,omitempty"`
	ManagedDeviceCertificateStates              *[]ManagedDeviceCertificateState                                        `json:"managedDeviceCertificateStates,omitempty"`
	ODataType                                   *string                                                                 `json:"@odata.type,omitempty"`
	RenewalThresholdPercentage                  *int64                                                                  `json:"renewalThresholdPercentage,omitempty"`
	RoleScopeTagIds                             *[]string                                                               `json:"roleScopeTagIds,omitempty"`
	RootCertificate                             *AndroidDeviceOwnerTrustedRootCertificate                               `json:"rootCertificate,omitempty"`
	SilentCertificateAccessDetails              *[]AndroidDeviceOwnerSilentCertificateAccess                            `json:"silentCertificateAccessDetails,omitempty"`
	SubjectAlternativeNameFormatString          *string                                                                 `json:"subjectAlternativeNameFormatString,omitempty"`
	SubjectAlternativeNameType                  *AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType     `json:"subjectAlternativeNameType,omitempty"`
	SubjectNameFormat                           *AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat              `json:"subjectNameFormat,omitempty"`
	SubjectNameFormatString                     *string                                                                 `json:"subjectNameFormatString,omitempty"`
	SupportsScopeTags                           *bool                                                                   `json:"supportsScopeTags,omitempty"`
	UserStatusOverview                          *DeviceConfigurationUserOverview                                        `json:"userStatusOverview,omitempty"`
	UserStatuses                                *[]DeviceConfigurationUserStatus                                        `json:"userStatuses,omitempty"`
	Version                                     *int64                                                                  `json:"version,omitempty"`
}
