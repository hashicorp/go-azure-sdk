package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10XSCEPCertificateProfile struct {
	Assignments                    *[]DeviceManagementResourceAccessProfileAssignment              `json:"assignments,omitempty"`
	CertificateStore               *Windows10XSCEPCertificateProfileCertificateStore               `json:"certificateStore,omitempty"`
	CertificateValidityPeriodScale *Windows10XSCEPCertificateProfileCertificateValidityPeriodScale `json:"certificateValidityPeriodScale,omitempty"`
	CertificateValidityPeriodValue *int64                                                          `json:"certificateValidityPeriodValue,omitempty"`
	CreationDateTime               *string                                                         `json:"creationDateTime,omitempty"`
	Description                    *string                                                         `json:"description,omitempty"`
	DisplayName                    *string                                                         `json:"displayName,omitempty"`
	ExtendedKeyUsages              *[]ExtendedKeyUsage                                             `json:"extendedKeyUsages,omitempty"`
	HashAlgorithm                  *[]Windows10XSCEPCertificateProfileHashAlgorithm                `json:"hashAlgorithm,omitempty"`
	Id                             *string                                                         `json:"id,omitempty"`
	KeySize                        *Windows10XSCEPCertificateProfileKeySize                        `json:"keySize,omitempty"`
	KeyStorageProvider             *Windows10XSCEPCertificateProfileKeyStorageProvider             `json:"keyStorageProvider,omitempty"`
	KeyUsage                       *Windows10XSCEPCertificateProfileKeyUsage                       `json:"keyUsage,omitempty"`
	LastModifiedDateTime           *string                                                         `json:"lastModifiedDateTime,omitempty"`
	ODataType                      *string                                                         `json:"@odata.type,omitempty"`
	RenewalThresholdPercentage     *int64                                                          `json:"renewalThresholdPercentage,omitempty"`
	RoleScopeTagIds                *[]string                                                       `json:"roleScopeTagIds,omitempty"`
	RootCertificateId              *string                                                         `json:"rootCertificateId,omitempty"`
	ScepServerUrls                 *[]string                                                       `json:"scepServerUrls,omitempty"`
	SubjectAlternativeNameFormats  *[]Windows10XCustomSubjectAlternativeName                       `json:"subjectAlternativeNameFormats,omitempty"`
	SubjectNameFormatString        *string                                                         `json:"subjectNameFormatString,omitempty"`
	Version                        *int64                                                          `json:"version,omitempty"`
}
