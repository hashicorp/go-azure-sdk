package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType string

const (
	AndroidForWorkCertificateProfileBaseSubjectAlternativeNameTypecustomAzureADAttribute      AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType = "CustomAzureADAttribute"
	AndroidForWorkCertificateProfileBaseSubjectAlternativeNameTypedomainNameService           AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType = "DomainNameService"
	AndroidForWorkCertificateProfileBaseSubjectAlternativeNameTypeemailAddress                AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType = "EmailAddress"
	AndroidForWorkCertificateProfileBaseSubjectAlternativeNameTypenone                        AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType = "None"
	AndroidForWorkCertificateProfileBaseSubjectAlternativeNameTypeuniversalResourceIdentifier AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType = "UniversalResourceIdentifier"
	AndroidForWorkCertificateProfileBaseSubjectAlternativeNameTypeuserPrincipalName           AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForAndroidForWorkCertificateProfileBaseSubjectAlternativeNameType() []string {
	return []string{
		string(AndroidForWorkCertificateProfileBaseSubjectAlternativeNameTypecustomAzureADAttribute),
		string(AndroidForWorkCertificateProfileBaseSubjectAlternativeNameTypedomainNameService),
		string(AndroidForWorkCertificateProfileBaseSubjectAlternativeNameTypeemailAddress),
		string(AndroidForWorkCertificateProfileBaseSubjectAlternativeNameTypenone),
		string(AndroidForWorkCertificateProfileBaseSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(AndroidForWorkCertificateProfileBaseSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseAndroidForWorkCertificateProfileBaseSubjectAlternativeNameType(input string) (*AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType, error) {
	vals := map[string]AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType{
		"customazureadattribute":      AndroidForWorkCertificateProfileBaseSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           AndroidForWorkCertificateProfileBaseSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                AndroidForWorkCertificateProfileBaseSubjectAlternativeNameTypeemailAddress,
		"none":                        AndroidForWorkCertificateProfileBaseSubjectAlternativeNameTypenone,
		"universalresourceidentifier": AndroidForWorkCertificateProfileBaseSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           AndroidForWorkCertificateProfileBaseSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkCertificateProfileBaseSubjectAlternativeNameType(input)
	return &out, nil
}
