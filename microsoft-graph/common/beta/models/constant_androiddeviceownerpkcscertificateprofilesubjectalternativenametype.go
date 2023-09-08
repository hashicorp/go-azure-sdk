package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType string

const (
	AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute      AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType = "CustomAzureADAttribute"
	AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypedomainNameService           AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType = "DomainNameService"
	AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypeemailAddress                AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType = "EmailAddress"
	AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypenone                        AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType = "None"
	AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType = "UniversalResourceIdentifier"
	AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypeuserPrincipalName           AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForAndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType() []string {
	return []string{
		string(AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute),
		string(AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypedomainNameService),
		string(AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypeemailAddress),
		string(AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypenone),
		string(AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseAndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType(input string) (*AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType, error) {
	vals := map[string]AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType{
		"customazureadattribute":      AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypeemailAddress,
		"none":                        AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypenone,
		"universalresourceidentifier": AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerPkcsCertificateProfileSubjectAlternativeNameType(input)
	return &out, nil
}
