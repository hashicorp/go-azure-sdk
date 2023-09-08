package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType string

const (
	AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypecustomAzureADAttribute      AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType = "CustomAzureADAttribute"
	AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypedomainNameService           AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType = "DomainNameService"
	AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypeemailAddress                AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType = "EmailAddress"
	AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypenone                        AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType = "None"
	AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypeuniversalResourceIdentifier AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType = "UniversalResourceIdentifier"
	AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypeuserPrincipalName           AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForAndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType() []string {
	return []string{
		string(AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypecustomAzureADAttribute),
		string(AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypedomainNameService),
		string(AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypeemailAddress),
		string(AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypenone),
		string(AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseAndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType(input string) (*AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType, error) {
	vals := map[string]AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType{
		"customazureadattribute":      AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypeemailAddress,
		"none":                        AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypenone,
		"universalresourceidentifier": AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerCertificateProfileBaseSubjectAlternativeNameType(input)
	return &out, nil
}
