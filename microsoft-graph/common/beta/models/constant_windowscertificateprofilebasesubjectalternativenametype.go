package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsCertificateProfileBaseSubjectAlternativeNameType string

const (
	WindowsCertificateProfileBaseSubjectAlternativeNameTypecustomAzureADAttribute      WindowsCertificateProfileBaseSubjectAlternativeNameType = "CustomAzureADAttribute"
	WindowsCertificateProfileBaseSubjectAlternativeNameTypedomainNameService           WindowsCertificateProfileBaseSubjectAlternativeNameType = "DomainNameService"
	WindowsCertificateProfileBaseSubjectAlternativeNameTypeemailAddress                WindowsCertificateProfileBaseSubjectAlternativeNameType = "EmailAddress"
	WindowsCertificateProfileBaseSubjectAlternativeNameTypenone                        WindowsCertificateProfileBaseSubjectAlternativeNameType = "None"
	WindowsCertificateProfileBaseSubjectAlternativeNameTypeuniversalResourceIdentifier WindowsCertificateProfileBaseSubjectAlternativeNameType = "UniversalResourceIdentifier"
	WindowsCertificateProfileBaseSubjectAlternativeNameTypeuserPrincipalName           WindowsCertificateProfileBaseSubjectAlternativeNameType = "UserPrincipalName"
)

func PossibleValuesForWindowsCertificateProfileBaseSubjectAlternativeNameType() []string {
	return []string{
		string(WindowsCertificateProfileBaseSubjectAlternativeNameTypecustomAzureADAttribute),
		string(WindowsCertificateProfileBaseSubjectAlternativeNameTypedomainNameService),
		string(WindowsCertificateProfileBaseSubjectAlternativeNameTypeemailAddress),
		string(WindowsCertificateProfileBaseSubjectAlternativeNameTypenone),
		string(WindowsCertificateProfileBaseSubjectAlternativeNameTypeuniversalResourceIdentifier),
		string(WindowsCertificateProfileBaseSubjectAlternativeNameTypeuserPrincipalName),
	}
}

func parseWindowsCertificateProfileBaseSubjectAlternativeNameType(input string) (*WindowsCertificateProfileBaseSubjectAlternativeNameType, error) {
	vals := map[string]WindowsCertificateProfileBaseSubjectAlternativeNameType{
		"customazureadattribute":      WindowsCertificateProfileBaseSubjectAlternativeNameTypecustomAzureADAttribute,
		"domainnameservice":           WindowsCertificateProfileBaseSubjectAlternativeNameTypedomainNameService,
		"emailaddress":                WindowsCertificateProfileBaseSubjectAlternativeNameTypeemailAddress,
		"none":                        WindowsCertificateProfileBaseSubjectAlternativeNameTypenone,
		"universalresourceidentifier": WindowsCertificateProfileBaseSubjectAlternativeNameTypeuniversalResourceIdentifier,
		"userprincipalname":           WindowsCertificateProfileBaseSubjectAlternativeNameTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsCertificateProfileBaseSubjectAlternativeNameType(input)
	return &out, nil
}
