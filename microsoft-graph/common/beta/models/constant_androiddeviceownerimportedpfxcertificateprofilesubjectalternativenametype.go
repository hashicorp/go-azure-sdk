package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType string

const (
	AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute      AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType = "CustomAzureADAttribute"
	AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameTypedomainNameService           AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType = "DomainNameService"
	AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameTypeemailAddress                AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType = "EmailAddress"
	AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameTypenone                        AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType = "None"
	AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType = "UniversalResourceIdentifier"
	AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameTypeuserPrincipalName           AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForAndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameTypedomainNameService),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameTypeemailAddress),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameTypenone),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseAndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType(input string) (*AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameTypeemailAddress,
		"none":                        AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameTypenone,
		"universalresourceidentifier": AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerImportedPFXCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
