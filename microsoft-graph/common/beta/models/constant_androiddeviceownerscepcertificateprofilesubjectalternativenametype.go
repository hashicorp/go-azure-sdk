package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType string

const (
	AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute      AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType = "CustomAzureADAttribute"
	AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypedomainNameService           AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType = "DomainNameService"
	AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypeemailAddress                AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType = "EmailAddress"
	AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypenone                        AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType = "None"
	AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType = "UniversalResourceIdentifier"
	AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypeuserPrincipalName           AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForAndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute),
		string(AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypedomainNameService),
		string(AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypeemailAddress),
		string(AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypenone),
		string(AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseAndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType(input string) (*AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypeemailAddress,
		"none":                        AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypenone,
		"universalresourceidentifier": AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerScepCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
