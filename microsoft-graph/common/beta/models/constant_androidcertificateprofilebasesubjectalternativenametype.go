package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidCertificateProfileBaseSubjectAlternativeNameType string

const (
	AndroidCertificateProfileBaseSubjectAlternativeNameTypecustomAzureADAttribute      AndroidCertificateProfileBaseSubjectAlternativeNameType = "CustomAzureADAttribute"
	AndroidCertificateProfileBaseSubjectAlternativeNameTypedomainNameService           AndroidCertificateProfileBaseSubjectAlternativeNameType = "DomainNameService"
	AndroidCertificateProfileBaseSubjectAlternativeNameTypeemailAddress                AndroidCertificateProfileBaseSubjectAlternativeNameType = "EmailAddress"
	AndroidCertificateProfileBaseSubjectAlternativeNameTypenone                        AndroidCertificateProfileBaseSubjectAlternativeNameType = "None"
	AndroidCertificateProfileBaseSubjectAlternativeNameTypeuniversalResourceIdentifier AndroidCertificateProfileBaseSubjectAlternativeNameType = "UniversalResourceIdentifier"
	AndroidCertificateProfileBaseSubjectAlternativeNameTypeuserPrincipalName           AndroidCertificateProfileBaseSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForAndroidCertificateProfileBaseSubjectAlternativeNameType() []string {
	return []string{
		string(AndroidCertificateProfileBaseSubjectAlternativeNameTypecustomAzureADAttribute),
		string(AndroidCertificateProfileBaseSubjectAlternativeNameTypedomainNameService),
		string(AndroidCertificateProfileBaseSubjectAlternativeNameTypeemailAddress),
		string(AndroidCertificateProfileBaseSubjectAlternativeNameTypenone),
		string(AndroidCertificateProfileBaseSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(AndroidCertificateProfileBaseSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseAndroidCertificateProfileBaseSubjectAlternativeNameType(input string) (*AndroidCertificateProfileBaseSubjectAlternativeNameType, error) {
	vals := map[string]AndroidCertificateProfileBaseSubjectAlternativeNameType{
		"customazureadattribute":      AndroidCertificateProfileBaseSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           AndroidCertificateProfileBaseSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                AndroidCertificateProfileBaseSubjectAlternativeNameTypeemailAddress,
		"none":                        AndroidCertificateProfileBaseSubjectAlternativeNameTypenone,
		"universalresourceidentifier": AndroidCertificateProfileBaseSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           AndroidCertificateProfileBaseSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidCertificateProfileBaseSubjectAlternativeNameType(input)
	return &out, nil
}
